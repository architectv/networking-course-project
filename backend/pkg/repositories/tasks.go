package repositories

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskMongo struct {
	db *mongo.Collection
}

func NewTaskMongo(db *mongo.Database) *TaskMongo {
	return &TaskMongo{
		db: db.Collection(viper.GetString("mongo.taskCollection")),
	}
}
