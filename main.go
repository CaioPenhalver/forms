package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Hey fuck")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@mongo:27017"))
	if err != nil {
		log.Fatal(err)

	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)

	}
	defer client.Disconnect(ctx)
	collection := client.Database("forms").Collection("form_schema")
	res, err := collection.InsertOne(context.Background(), bson.M{"name": "My Form"})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	fmt.Println(id)

	result := struct {
		Name string
	}{}
	// var result bson.M

	filter := bson.D{{"name", "My Form"}}

	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(result)
	fmt.Printf("Name: %s \n", result.Name)
}
