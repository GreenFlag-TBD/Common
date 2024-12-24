package db_driver

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDriver struct {
	*mongo.Client
}

func NewMongoDriver(dsn string) (*MongoDriver, error) {
	mongoOptions := options.Client().ApplyURI(dsn)
	mongoOptions.SetMaxConnIdleTime(10)
	mongoOptions.SetMaxPoolSize(100)
	mongoOptions.SetMinPoolSize(10)
	client, err := mongo.Connect(context.Background(), mongoOptions)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	return &MongoDriver{client}, nil
}
