package data

import (
	"context"
	"fmt"
	"github.com/kiwsan/go-jwt-auth/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/
func ClientDb() (*mongo.Client, error) {

	// Database authentication
	env := utils.Config.Database

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/admin", env.User, env.Password, env.Host, env.Port)

	// Set client options and connect
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}
