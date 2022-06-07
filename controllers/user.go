package controllers

import (
	"Golang_latihan/big_friendship_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateUserInput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /users
// Create new users
func CreateUsers(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users models.User

	db := c.MustGet("db").(*gorm.DB)

	//validate existing username, mail and phone
	if err := db.Where("username = ? OR email = ? OR phone = ?", input.Username, input.Email, input.Phone).First(&users).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already existing username / email / phone!!"})
		return
	}

	// Create users
	users2 := models.User{Username: input.Username, Fullname: input.Fullname, Gender: input.Gender, Phone: input.Phone, Email: input.Email}

	db.Create(&users2)

	c.JSON(http.StatusOK, gin.H{"data": users2})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) { // Get model if exist
	var users models.User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// PATCH /users/:id
// Update a user
func UpdateUsers(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var users models.User
	if err := db.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.User
	updatedInput.Username = input.Username
	updatedInput.Fullname = input.Fullname
	updatedInput.Gender = input.Gender
	updatedInput.Phone = input.Phone
	updatedInput.Email = input.Email

	//validate existing username, mail and phone
	if users.Username != updatedInput.Username {
		var users2 models.User
		if err := db.Where("username = ?", updatedInput.Username).First(&users2).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Already existing username!!"})
			return
		}
	}
	if users.Email != updatedInput.Email {
		var users2 models.User
		if err := db.Where("email = ?", updatedInput.Email).First(&users2).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Already existing email!!"})
			return
		}
	}
	if users.Phone != updatedInput.Phone {
		var users2 models.User
		if err := db.Where("phone = ?", updatedInput.Phone).First(&users2).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Already existing phone!!"})
			return
		}
	}

	db.Model(&users).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DELETE /users/:id
// Delete a user
func DeleteUsers(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var users models.User
	if err := db.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&users)

	c.JSON(http.StatusOK, gin.H{"data": "Delete success!"})
}
