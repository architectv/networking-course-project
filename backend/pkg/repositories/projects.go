package repositories

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectMongo struct {
	db *mongo.Collection
}

func NewProjectMongo(db *mongo.Database) *ProjectMongo {
	return &ProjectMongo{
		db: db.Collection(viper.GetString("mongo.projectCollection")),
	}
}
