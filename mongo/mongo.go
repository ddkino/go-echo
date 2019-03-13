package main

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	// Rest of the code will go here
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	context.WithTimeout(context.Background(), 10*time.Second)

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("kb").Collection("programmesseloger")

	defer client.Disconnect(context.TODO())
	// err = client.Disconnect(context.TODO())

	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Connection to MongoDB closed.")

	// create a value into which the result can be decoded
	var result interface{}
	filter := bson.M{}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
}
