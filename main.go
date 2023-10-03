package main

import (
	"log"

	"github.com/adeoluwa/gocommerce/controllers"
	"github.com/adeoluwa/gocommerce/database"
	"github.com/adeoluwa/gocommerce/middleware"
	"github.com/adeoluwa/gocommerce/routes"
	"github.com/gin-gonic/gin"

	"os"
)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client,"Users"))

	router := gin.New() 
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAdresses())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}