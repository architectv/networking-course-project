package mongo

import (
	"github.com/sirupsen/logrus"
	"context"
	"yak/backend/pkg/models"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "encoding/json"
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

func (r *UserMongo) GetById(id string) (models.User, error) {
	var user models.User
	ctx := context.TODO()
	logrus.Println(id)
	objID, _ := primitive.ObjectIDFromHex(id)
	logrus.Println(objID)
	filter := bson.M{"_id": objID}
	err := r.db.FindOne(ctx, filter).Decode(&user)
	// err := r.db.FindOne(ctx, bson.M{
	// 	"_id": objID,
	// }).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserMongo) Create(user models.User) (string, error) {
	ctx := context.TODO()
	bsonUser, _ := bson.Marshal(user)

	res, err := r.db.InsertOne(ctx, bsonUser)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}