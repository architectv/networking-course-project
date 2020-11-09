package mongo

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskListMongo struct {
	db *mongo.Collection
}

func NewTaskListMongo(db *mongo.Database) *TaskListMongo {
	return &TaskListMongo{
		db: db.Collection(viper.GetString("mongo.taskListCollection")),
	}
}
