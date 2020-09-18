package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zarszz/warehouse-rest-api/controllers"
	"github.com/zarszz/warehouse-rest-api/utils"
)

func main() {
	router := gin.Default()

	// connect to database and running autoCommit
	utils.ConnectDatabase()

	// user context
	router.POST("/user", controllers.CreateUser)
	router.GET("/users", controllers.FindUsers)
	router.GET("/user/:id", controllers.FindUser)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)

	// warehouse context
	router.POST("/warehouse", controllers.CreateWarehouse)
	router.GET("/warehouses", controllers.FindWarehouses)
	router.GET("/warehouse/:id", controllers.FindWarehouse)
	router.PUT("/warehouse/:id", controllers.UpdateWarehouse)
	router.DELETE("/warehouse/:id", controllers.DeleteWarehouse)

	// user address context
	router.PUT("/user/:id/address", controllers.UpdateUserAddress)

	router.Run(utils.GetEnvVariable("LISTEN_PORT"))
}
