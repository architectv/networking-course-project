package mongo

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardMongo struct {
	db *mongo.Collection
}

func NewBoardMongo(db *mongo.Database) *BoardMongo {
	return &BoardMongo{
		db: db.Collection(viper.GetString("mongo.boardCollection")),
	}
}
