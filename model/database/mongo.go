package database

import (
	"context"
	userscore "harfishRestAPI/business/userscore"
	"harfishRestAPI/helpers/messages"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserScore struct {
	Conn *mongo.Database
}

func NewMongoUserScore(conn *mongo.Database) userscore.Repository {
	return &MongoUserScore{
		Conn: conn,
	}
}

func (repository *MongoUserScore) GetAllScore(ctx context.Context) ([]userscore.UserScoreDomain, error) {
	var result []userscore.UserScoreDomain
	cursor, err := repository.Conn.Collection("UserScore").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return []userscore.UserScoreDomain{}, err
	}
	// fmt.Println(result)
	return result, nil
}

func (repositrory *MongoUserScore) CreateScoreAndFeedback(ctx context.Context, data *userscore.UserScoreDomain) error {
	scoreInsert := UserScoreFromDomain(*data)
	_, err := repositrory.Conn.Collection("UserScore").InsertOne(ctx, scoreInsert)
	if err != nil {
		return nil
	}
	return messages.ErrInsertSuccess
}
