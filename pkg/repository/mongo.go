package repository

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host       string
	Port       string
	Username   string
	Password   string
	DataBase   string
	Collection string
}

func NewMongoDB(cfg Config) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port))

	fmt.Print(clientOptions)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	logrus.Info("Connected to MongoDB")

	return client.Database(cfg.DataBase).Collection(cfg.Collection), nil
}
