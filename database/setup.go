package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/tools/go/analysis/passes/defers"
)

func DBSet() *mongo.Client{
	Sk := godotenv.Load()
	if Sk != nil {
		log.Fatal("Error loading .env file")
	}
	
	connection_string := os.Getenv("CONNECTION_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))

	if err!=nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err!= nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err!= nil {
		log.Println("failed to connect :(")
		return nil
	}

	fmt.Println("Successfully connected to mongo")

	return client
}

var Client *mongo.Client = DBSet()



func UserData(client *mongo.Client, collectionName string) *mongo.Collection{
	if client == nil {
		log.Fatal("Cannot get user data: MongoDB client is not connected")
	}
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection{
	if client == nil {
		log.Fatal("cannot get product data: MongoDB client is not connected")
	}
	var productCollection *mongo.Collection =  client.Database("Ecommerce").Collection(collectionName)

	return productCollection
}

