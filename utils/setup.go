package utils

import (
	"fmt"

	"github.com/zarszz/warehouse-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB - database object
var DB *gorm.DB

// ConnectDatabase - connect and create migration
func ConnectDatabase() {
	dsn := GetEnvVariable("CONN_STRING")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	database.AutoMigrate(&models.Category{})
	database.AutoMigrate(&models.UserAddress{})
	database.AutoMigrate(&models.Item{})
	database.AutoMigrate(&models.WarehouseAddress{})
	database.AutoMigrate(&models.Warehouse{})
	database.AutoMigrate(&models.User{})

	DB = database
}
