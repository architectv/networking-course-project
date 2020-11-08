package repositories

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB() (*mongo.Database, error) {
	opts := options.Client().ApplyURI(viper.GetString("mongo.uri"))
	client, err := mongo.NewClient(opts)
	if err != nil {
		logrus.Fatalf("Error occured while establishing connection to mongoDB")
		return nil, err
	}

	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	return client.Database(viper.GetString("mongo.name")), nil
}
