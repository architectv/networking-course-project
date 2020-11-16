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
}

func NewProjectMongo(db *mongo.Database) *ProjectMongo {
	return &ProjectMongo{
		db: db.Collection(viper.GetString("mongo.projectCollection")),
	}
}


func (r *ProjectMongo) GetAll(userId string) ([]models.Project, error) {
	projects := make([]models.Project, 0)
	ctx := context.TODO()
	fmt.Println(userId)
	cur, err := r.db.Find(ctx, bson.M{"ownerId": userId})
	defer cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var project models.Project
		err := cur.Decode(&project)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return projects, nil
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