package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/zarszz/warehouse-rest-api/models"
	"github.com/zarszz/warehouse-rest-api/utils"
	"gorm.io/gorm"
)

type CategoryInput struct {
	Category string `json:"category"`
}

func CreateCategory(c *gin.Context) {
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data should not empty"})
		return
	}
	err := utils.DB.Create(&models.Category{Category: input.Category}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error", "Error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "category successfully created"})
}

func FindCategory(c *gin.Context) {
	categoryId := c.Param("id")
	if categoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Data should not empty"})
		return
	}
	category := models.Category{}
	err := utils.DB.Find(&category, categoryId).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func FindCategories(c *gin.Context) {
	categories := []models.Category{}
	err := utils.DB.Find(&categories).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	if len(categories) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []models.Category{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func UpdateCategory(c *gin.Context) {
	categoryId := c.Param("id")
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil || categoryId == "" {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Data should not empty"})
		return
	}

	var category = models.Category{}

	err := utils.DB.Find(&category, categoryId).Error
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})

	}

	err2 := utils.DB.Model(&category).Updates(models.Category{Category: input.Category}).Where("id = ", categoryId).Error
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "update successfully"})

}

func DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")
	if categoryId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Data should not empty"})
		return
	}
	err := utils.DB.Delete(&models.Category{}, categoryId).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
