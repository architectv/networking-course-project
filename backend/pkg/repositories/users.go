package repositories

import (
	"context"
	"yak/backend/pkg/models"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongo struct {
	db *mongo.Collection
}

func NewUserMongo(db *mongo.Database) *UserMongo {
	return &UserMongo{
		db: db.Collection(viper.GetString("mongo.userCollection")),
	}
}

func (r *UserMongo) GetAll() ([]models.User, error) {
	users := make([]models.User, 0)
	ctx := context.TODO()
	cur, err := r.db.Find(ctx, bson.M{})
	defer cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
