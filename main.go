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
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.FindUsers)
	router.GET("/users/:id", controllers.FindUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// warehouse context
	router.POST("/warehouses", controllers.CreateWarehouse)
	router.GET("/warehouses", controllers.FindWarehouses)
	router.GET("/warehouses/:id", controllers.FindWarehouse)
	router.PUT("/warehouses/:id", controllers.UpdateWarehouse)
	router.DELETE("/warehouses/:id", controllers.DeleteWarehouse)

	// category context
	router.POST("/categories", controllers.CreateCategory)
	router.GET("/categories", controllers.FindCategories)
	router.GET("/categories/:id", controllers.FindCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)

	// user address context
	router.PUT("/user/:id/address", controllers.UpdateUserAddress)

	router.Run(utils.GetEnvVariable("LISTEN_PORT"))
}
