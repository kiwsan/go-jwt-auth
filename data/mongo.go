package data

import (
	"context"
	"fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/
func DatabaseCollection() (*mongo.Collection, error) {
	
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database("identityService").Collection("users")
	
	fmt.Println("Connected to MongoDB!")

	return collection, nil
}