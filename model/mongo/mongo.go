package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func (config *Config) ConnectDB() (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.gufx7.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		config.Username,
		config.Password)

	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database(config.Name)
	return db, nil
}
