package controllers

import (
	"github.com/gin-gonic/gin"
)

type ItemInput struct {
	ItemName    string `json:"item_name"`
	CategoryID  uint   `json:"category_id"`
	WarehouseID string `json:"warehouse_id"`
	OwnerID     string `json:"owner_Id"`
}

func CreateItem(c *gin.Context) {

}

func FindItem(c *gin.Context) {

}

func FindItems(c *gin.Context) {

}

func UpdateItem(c *gin.Context) {

}

func DeleteItem(c *gin.Context) {

}

func TakeItemFromWarehouse(c *gin.Context) {

}
