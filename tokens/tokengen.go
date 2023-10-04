package tokens

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/adeoluwa/gocommerce/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct{
	Email string 
	First_Name string
	Last_Name string
	Delivery_Address string
	Uid string
	jwt.StandardClaims
}

var UserData *mongo.Collection = database.UserData(database.Client, "Users")

var SECRET_KEY = os.Getenv("SECRET_KEY")

func TokenGenerator(email string, first_name string, last_name string,delivery_address string, uid string)(signedtoken string, signedrefreshtoken string, err error){
	claims := &SignedDetails{
		Email: email,
		First_Name: first_name,
		Last_Name: last_name,
		Delivery_Address: delivery_address,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),  
		},
	}

	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "","", err
	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshtoken, err
}

func ValidateToken(signedtoken string)(claims *SignedDetails, msg string){
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token)(interface{}, error){
		return []byte(SECRET_KEY), nil
	})

	if err != nil{
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return 
	}

	if claims.ExpiresAt < time.Now().Local().Unix(){
		msg = "token is already expired"
		return
	}

	return claims, msg

}

func UpdateAllTokens(signedtoken string, signedrefreshtoken string, userid string){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateobj primitive.D

	updateobj = append (updateobj, bson.E{Key:"token", Value: signedtoken})
	updateobj = append (updateobj, bson.E{Key:"refresh_token", Value: signedrefreshtoken})
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateobj = append(updateobj, bson.E{Key:"updatedat", Value:updated_at})

	upsert := true

	filter := bson.M{"user_id":userid}
	opt := options.UpdateOptions{
		Upsert: &upsert,

	}

	_, err := UserData.UpdateOne(ctx, filter, bson.D{
		{Key:"$set", Value: updateobj},
	},
	&opt)

	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}
}