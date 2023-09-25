package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adeoluwa/gocommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/internal"
	"golang.org/x/net/context"
)



func AddAdress() gin.HandlerFunc{
	return func(c *gin.Context){
		user_id := c.Query("id")
		if user_id == ""{
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error":"Invalid code"})
			c.Abort()
			return
		}
		address, err := primitive.ObjectIDFromHex(user_id) // not sure of the primitive
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var addresses models.Address

		addresses.Address_ID = primitive.NewObjectID() 

		if err = c.BindJSON(&address); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second) 

		match_filter := bson.D{{Key:"$match", Value: bson.D{primitive.E{Key:"_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value:bson.D{primitive.E{Key: "path", Value:"$address"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value:"$address_id"}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}

		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter, unwind, group})

		if err!= nil {
			c.IndentedJSON(500, "Internal server error")
		}

		var addressinfo []bson.M

		if err = pointcursor.All(ctx, &addressinfo); err != nil {
			panic(err)
		}

		var size int32
		for _, address_no := range addressinfo {
			count := address_no["count"]
			size = count.(int32)
		}

		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{Key:"$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
			_, err := UserCollection.UpdateOne(ctx, filter, update)
			if err != nil{
				fmt.Println(err)
			}
		}else{
			c.IndentedJSON(400, "Not Allowed")
		}
		defer cancel()
		ctx.Done() 
	}


}

func EditHomeAddress() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}

func EditWorkAddress() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}

func DeleteAddress() gin.HandlerFunc{
	return func(c *gin.Context){
		user_id := c.Query("id")

		if user_id ==""{
			c.Header("Context-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error":"Invalid Search Index"})
			c.Abort()
			return
		}

		addresses :=  make([]models.Address, 0)
		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "internal Server Error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key:"_id", Value: usert_id}}
		update := bson.D{{Key:"$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil{
			c.IndentedJSON(404, "Not Found, wrong command")
		}

		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully deleted the address")
	
	}
}