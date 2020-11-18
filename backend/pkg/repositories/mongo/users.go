package mongo

import (
	"context"
	"yak/backend/pkg/models"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *UserMongo) GetUser(username, password string) (models.User, error) {
	var user models.User
	ctx := context.TODO()
	filter := bson.M{"username": username, "password": password}
	err := r.db.FindOne(ctx, filter).Decode(&user)

	return user, err
}

func (r *UserMongo) Create(ctx context.Context, user *models.User) (string, error) {
	bsonUser, err := bson.Marshal(user)
	if err != nil {
		return "", err
	}

	res, err := r.db.InsertOne(ctx, bsonUser)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *UserMongo) GetByNickname(ctx context.Context, nickname string) (*models.User, error) {
	user := &models.User{}
	filter := bson.M{"nickname": nickname}
	err := r.db.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, err
}
