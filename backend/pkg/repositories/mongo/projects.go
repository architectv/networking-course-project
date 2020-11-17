package mongo

import (
	"fmt"
	"context"
	"yak/backend/pkg/models"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectMongo struct {
	db *mongo.Collection
	projectUserDb *mongo.Collection
}

func NewProjectMongo(db *mongo.Database) *ProjectMongo {
	return &ProjectMongo{
		db: db.Collection(viper.GetString("mongo.projectCollection")),
		projectUserDb: db.Collection(viper.GetString("mongo.projectUser")),
	}
}

func (r *ProjectMongo) Create(project models.Project) (string, error) {
	ctx := context.TODO()
	bsonProject, _ := bson.Marshal(project)

	res, err := r.db.InsertOne(ctx, bsonProject)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *ProjectMongo) ProjectUserCreate(projectUser models.ProjectUser) error {
	ctx := context.TODO()
	bsonProjectUser, _ := bson.Marshal(projectUser)

	_, err := r.projectUserDb.InsertOne(ctx, bsonProjectUser)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProjectMongo) GetById(projectId string) (models.Project, error) {
	var project models.Project
	ctx := context.TODO()
	objID, _ := primitive.ObjectIDFromHex(projectId)
	
	err := r.db.FindOne(ctx, bson.M{"_id": objID}).Decode(&project)
    fmt.Println(projectId)
	return project, err
}

func (r *ProjectMongo) GetPermission(userId, projectId string) (*models.Permission, error) {
	var projectUser models.ProjectUser
	ctx := context.TODO()
	err := r.projectUserDb.FindOne(ctx, bson.M{
		"userId": userId,
		"projectId": projectId,
	}).Decode(&projectUser)
	
	return projectUser.Permissions, err
}

func (r *ProjectMongo) GetAll(projectsId []string) ([]models.Project, error) {
	projects := make([]models.Project, 0)
	for _, val := range projectsId {
		curProject, err := r.GetById(val)
		if err != nil {
			return nil, err
		}
		projects = append(projects, curProject)
	}
	return projects, nil
}


func (r *ProjectMongo) ProjectIdGetByPermissions(userId string) ([]string, error) {
	projectsId := make([]string, 0)
	ctx := context.TODO()
	cur, err := r.projectUserDb.Find(ctx, bson.M{"userId": userId})
	defer cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var projectUser models.ProjectUser
		err := cur.Decode(&projectUser)
		if err != nil {
			return nil, err
		}
		if projectUser.Permissions.Read == true {        // TODO: условие должно быть в Find 
			projectsId = append(projectsId, projectUser.ProjectId)
		}
		
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return projectsId, nil
}

func (r *ProjectMongo) Update(projectId string, project models.Project) error {
	ctx := context.TODO()
	objID, _ := primitive.ObjectIDFromHex(projectId)
	filter := bson.M{"_id": objID}
	update := bson.D{
		{"$set", bson.M{
			"defaultPermissions": project.DefaultPermissions,
			"datetimes": project.Datetimes,
			"title": project.Title,
			"description": project.Description,
		}},
	}

	_, err := r.db.UpdateOne(ctx, filter, update)
	return err
}

