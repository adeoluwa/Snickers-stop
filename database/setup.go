package database

import (
	"context"
	"fmt"
	"log"
	"time"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/tools/go/analysis/passes/defers"
)

func DBSet() *mongo.Client{
	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// opts := options.Client().ApplyURI("mongodb+srv://kiishijoseph:ByPf4UvZCr0OqKPe@cluster0.akomxxd.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	

	// client, err := mongo.Connect(context.TODO(), opts)
	// if err != nil {
	// 	panic(err)
	// }

	// defer func(){
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// // send a ping to confirm a successful connection
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Pinged your deployment. You successfully connect to MongoDB!")

	/////////

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kiishijoseph:ByPf4UvZCr0OqKPe@cluster0.akomxxd.mongodb.net/?retryWrites=true&w=majority"))

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

