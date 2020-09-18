package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/warehouse-rest-api/models"
	"github.com/zarszz/warehouse-rest-api/utils"
	"gorm.io/gorm"
)

type userInput struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
	State       string `json:"state" binding:"required"`
	Country     string `json:"country" binding:"required"`
	PostalCode  string `json:"postal_code" binding:"required"`
}

type userUpdateInput struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}

type userUpdateAddressInput struct {
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	Country    string `json:"country" binding:"required"`
	PostalCode string `json:"postal_code" binding:"required"`
}

// CreateUser - create a new user
// POST /user
func CreateUser(c *gin.Context) {
	var input userInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// parse birth date to YYYY-MM-dd format
	birthDate, err := time.Parse("2006-Jan-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	user := models.User{IsAdmin: false, Email: input.Email, Password: input.Password, FirstName: input.FirstName, LastName: input.LastName, Gender: input.Gender, DateOfBirth: birthDate}
	userResult := utils.DB.Create(&user)
	if userResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": userResult.Error.Error()})
		return
	}

	// insert user address data to database
	userAddress := models.UserAddress{Address: input.Address, City: input.City, State: input.State, Country: input.Country, PostalCode: input.PostalCode, UserID: user.ID}
	userAddressResult := utils.DB.Create(&userAddress)
	if userAddressResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": userAddressResult.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "sucessfully registered", "user": user})
}

// FindUsers - retrieve all registered users
// GET /users
func FindUsers(c *gin.Context) {
	var users []models.User

	err := utils.DB.Preload("UserAddress").Find(&users).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// FindUser - retrieve user base on received id
// GET /user
func FindUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	err := utils.DB.Preload("UserAddress").Where("users.id = ?", id).Find(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser - Update user base on id
// PUT /user/{id}
func UpdateUser(c *gin.Context) {
	var input userUpdateInput
	var user models.User

	id := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// parse birth date to YYYY-MM-dd format
	birthDate, err := time.Parse("2006-Jan-02", input.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// find target user to update, reference to id
	utils.DB.Find(&user, id)

	// err = utils.DB.Where("users.id = ?", id).Updates(user).Error
	// execute update
	err = utils.DB.Model(&user).Updates(models.User{IsAdmin: false, Email: input.Email, Password: input.Password, FirstName: input.FirstName, LastName: input.LastName, Gender: input.Gender, DateOfBirth: birthDate}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "update successfully"})
}

// UpdateUserAddress - Update user address base on id
// PUT /user/{id}/address
func UpdateUserAddress(c *gin.Context) {
	var input userUpdateAddressInput

	id := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// parse userID to integer
	intParsedUserID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// parse userID to uint
	parsedUserID := uint(intParsedUserID)

	// create address object
	address := models.UserAddress{UserID: parsedUserID, Address: input.Address, City: input.City, State: input.State, Country: input.Country, PostalCode: input.PostalCode}

	err = utils.DB.Where("user_id = ?", id).Updates(address).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "update successfully"})
}

// DeleteUser - Delete a user base on id
// DELETE /user/{id}
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	userErr := utils.DB.Delete(&models.User{}, id).Error
	if userErr != nil {
		if errors.Is(userErr, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": userErr.Error()})
		return
	}

	userAddressErr := utils.DB.Delete(&models.UserAddress{}, id).Error
	if userAddressErr != nil {
		if errors.Is(userAddressErr, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": userAddressErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
