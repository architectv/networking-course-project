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

func (r *ProjectMongo) GetById(id string) (models.Project, error) {
	var project models.Project
	ctx := context.TODO()
	objID, _ := primitive.ObjectIDFromHex(id)
	
	err := r.db.FindOne(ctx, bson.M{"_id": objID}).Decode(&project)
    fmt.Println(id)
	return project, err
}

func (r *ProjectMongo) GetAll(projectsId []string) ([]models.Project, error) {
	projects := make([]models.Project, 0)
	for _, val := range projectsId {
		fmt.Println(val)
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
	// p := models.Permission {true, false, false}
	cur, err := r.db.Find(ctx, bson.M{"permissions":bson.M{"read": true}})
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

		projectsId = append(projectsId, projectUser.ProjectId)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return projectsId, nil
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
