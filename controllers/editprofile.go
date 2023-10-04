package controllers

import (
    "net/http"
	"time"
	"context"

    // "github.com/adeoluwa/gocommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	DB *mongo.Database
}

type UpdateUserRequest struct {
	FirstName			*string		`json:"first_name"`
	Last_Name			*string		`json:"last_name"`
	Phone				*string		`json:"phone"`
	Delivery_Address	*string		`json:"delivery_address"`
	Email				*string		`json:"email"`
}

func (u *UserController) UpdateUserProfile(c *gin.Context){
	var request UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id":id}
	update := bson.M{
		"$set":bson.M{
            "first_name":request.FirstName,
            "last_name":request.Last_Name,
            "phone":request.Phone,
            "delivery_address":request.Delivery_Address,
            "email":request.Email,
			"updated_at": time.Now(),
        },
	}

	_, err := u.DB.Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Could not update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"User profile updated successfully"})
}