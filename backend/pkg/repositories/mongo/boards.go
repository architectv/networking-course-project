package mongo

import (
	"context"
	"yak/backend/pkg/models"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *BoardMongo) GetAll(userId, projectId string) ([]models.Board, error) {
	// TODO: check permissions
	boards := make([]models.Board, 0)
	ctx := context.TODO()
	cur, err := r.db.Find(ctx, bson.M{
		"projectId": projectId,
	})
	defer cur.Close(ctx)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var board models.Board
		err := cur.Decode(&board)
		if err != nil {
			return nil, err
		}

		boards = append(boards, board)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return boards, nil
}

type BoardUser struct {
	userId      string             `json:"userId" bson:"userId"`
	boardId     string             `json:"boardId" bson:"boardId"`
	permissions *models.Permission `json:"permissions" bson:"permissions"`
}

func (r *BoardMongo) Create(userId, projectId string, board models.Board) (string, error) {
	ctx := context.TODO()
	bsonBoard, err := bson.Marshal(board)
	if err != nil {
		return "", err
	}

	res, err := r.db.InsertOne(ctx, bsonBoard)
	if err != nil {
		return "", err
	}

	boardId := res.InsertedID.(primitive.ObjectID).Hex()

	// relation := BoardUser{
	// 	userId:  userId,
	// 	boardId: boardId,
	// 	permissions: &models.Permission{
	// 		Read:  true,
	// 		Write: true,
	// 		Grant: true,
	// 	},
	// }

	return boardId, nil
}

func (r *BoardMongo) Delete(userId, boardId string) error {
	return nil
}
