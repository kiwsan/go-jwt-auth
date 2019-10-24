package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/
func ClientDb() (*mongo.Client, error) {

	// Database authentication
	host := "localhost"
	user := "root" // The database user is required administrator roles.
	password := "Str0ngPassword!"

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:27017/admin", user, password, host)

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
