package controllers

import (
	"Golang_latihan/big_friendship_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateRequestInput struct {
	ID               int `json:"id"`
	Id_user_req_from int `json:"id_user_req_from"`
	Id_user_req_to   int `json:"id_user_req_to"`
	Status           int `json:"status"` //1=request, 0=cancel_request
}

type UpdateRequestInput struct {
	Status int `json:"status"` //1=request, 0=cancel_request
}

// GET /requests
// Get all requests
func FindRequests(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var requests []models.Request_friendship
	db.Find(&requests)

	c.JSON(http.StatusOK, gin.H{"data": requests})
}

// POST /requests
// Create new requests
func CreateRequest(c *gin.Context) {
	// Validate input
	var input CreateRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var requests models.Request_friendship

	db := c.MustGet("db").(*gorm.DB)

	// Get Id_user_req_from and Id_user_req_to if exist
	var users models.User
	if err := db.Where("id = ?", input.Id_user_req_from).Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Request From not found!!"})
		return
	}
	var users2 models.User
	if err := db.Where("id = ?", input.Id_user_req_to).Find(&users2).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Request To not found!!"})
		return
	}

	//validate existing requested friendship
	if err := db.Where("id_user_req_from = ? AND id_user_req_to = ? AND status = 1", input.Id_user_req_from, input.Id_user_req_to).Or("id_user_req_from = ? AND id_user_req_to = ? AND status = 1", input.Id_user_req_to, input.Id_user_req_from).First(&requests).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already existing requested friendship from them!!"})
		return
	}

	// Create request
	requests2 := models.Request_friendship{Id_user_req_from: input.Id_user_req_from, Id_user_req_to: input.Id_user_req_to, Status: 1}

	db.Create(&requests2)

	c.JSON(http.StatusOK, gin.H{"data": requests2})
}

// GET /requests/:id
// Find a requests
func FindRequest(c *gin.Context) { // Get model if exist
	var requests models.Request_friendship

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&requests).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requests})
}

// GET /request_to_me
// Find a requests
func FindRequestToMe(c *gin.Context) { // Get model if exist
	var requests []models.Request_friendship

	// Validate input
	var input CreateRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Joins("LEFT JOIN users ON request_friendships.Id_user_req_to=users.id").Where("request_friendships.id_user_req_to = ? AND request_friendships.status=1", input.Id_user_req_to).Find(&requests).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": requests})
}

// PATCH /requests/:id
// Update a requests
func UpdateRequest(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var requests models.Request_friendship
	if err := db.Where("id = ?", c.Param("id")).First(&requests).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Status < 1 || input.Status > 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status 1 untuk request, status 2 untuk cancel request"})
		return
	}

	var updatedInput models.Request_friendship
	updatedInput.Status = input.Status

	db.Model(&requests).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": requests})
}

// DELETE /requests/:id
// Delete a requests
func DeleteRequests(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var requests models.Request_friendship
	if err := db.Where("id = ?", c.Param("id")).First(&requests).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&requests)

	c.JSON(http.StatusOK, gin.H{"data": "Delete success!"})
}
