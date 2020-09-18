package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/warehouse-rest-api/models"
	"github.com/zarszz/warehouse-rest-api/utils"
	"gorm.io/gorm/clause"
)

// Warehouse - warehouse model struct
type Warehouse struct {
	WarehouseID   string `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	City          string `json:"city" binding:"required"`
	Province      string `json:"province" binding:"required"`
	Country       string `json:"country" binding:"required"`
	PostalCode    string `json:"postal_code" binding:"required"`
}

// CreateWarehouse - create a new house
func CreateWarehouse(c *gin.Context) {
	var input Warehouse

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// insert warehouse address
	warehouse := models.Warehouse{
		ID:            input.WarehouseID,
		WarehouseName: input.WarehouseName,
	}
	warehouseInsertError := utils.DB.Create(&warehouse).Error
	if warehouseInsertError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": warehouseInsertError.Error()})
		return
	}

	// insert warehouse address first
	warehouseAddress := models.WarehouseAddress{
		WarehouseID: warehouse.ID,
		Address:     input.Address,
		City:        input.City,
		Province:    input.Province,
		Country:     input.Country,
		PostalCode:  input.PostalCode,
	}
	warehouseAddressInsertError := utils.DB.Create(&warehouseAddress).Error
	if warehouseAddressInsertError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": warehouseInsertError.Error()})
		return
	}

	warehouse = models.Warehouse{
		ID:            input.WarehouseID,
		WarehouseName: input.WarehouseName,
		Address:       warehouseAddress,
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"Warehouse": warehouse,
		},
		"status": "success",
	})

}

// FindWarehouse - find a warehouse base on id
func FindWarehouse(c *gin.Context) {
	var warehouse models.Warehouse
	id := c.Param("id")

	err := utils.DB.Where("id = ?", id).Preload(clause.Associations).Find(&warehouse).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, warehouse)
}

// FindWarehouses - retrieve all warehouses data
func FindWarehouses(c *gin.Context) {
	var warehouses []models.Warehouse
	err := utils.DB.Preload(clause.Associations).Find(&warehouses).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, warehouses)
}

// UpdateWarehouse - update warehouse data based on id
func UpdateWarehouse(c *gin.Context) {
	var input Warehouse
	var warehouse models.Warehouse
	var warehouseAddress models.WarehouseAddress

	id := c.Param("id")

	// ensure id not blank
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "id cannot be empty"})
		return
	}

	// bind received json to object
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// find warehouse by id
	err := utils.DB.Where("id = ?", id).Find(&warehouse).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// ensure if warehouse data is available
	if warehouse.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": fmt.Sprintf("warehouse with %s id not found", id)})
		return
	}

	// find warehouseAddress by id
	err = utils.DB.Where("warehouse_id = ?", id).Find(&warehouseAddress).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// create object
	updatedWarehouse := models.Warehouse{
		ID:            warehouse.ID,
		WarehouseName: input.WarehouseName,
	}
	updatedWarehouseAddress := models.WarehouseAddress{
		WarehouseID: warehouse.ID,
		Address:     input.Address,
		City:        input.City,
		Province:    input.Province,
		Country:     input.Country,
		PostalCode:  input.PostalCode,
	}

	// apply new data to database
	err = utils.DB.Model(&warehouseAddress).Updates(updatedWarehouseAddress).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// apply new data to database
	err = utils.DB.Model(&warehouse).Updates(updatedWarehouse).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "update data success"})
}

// DeleteWarehouse - delete warehouse data based on id
func DeleteWarehouse(c *gin.Context) {
	var warehouse models.Warehouse

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "id cannot be empty"})
		return
	}

	err := utils.DB.Where("id = ?", id).Find(&warehouse).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	if warehouse.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": fmt.Sprintf("warehouse with %s id not found", id)})
		return
	}
	err = utils.DB.Where("id = ?", id).Delete(&warehouse).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "delete data successfully"})
}
