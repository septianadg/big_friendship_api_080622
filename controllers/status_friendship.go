package controllers

import (
	"Golang_latihan/big_friendship_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateStatusInput struct {
	ID                    int `json:"id"`
	Id_request_friendship int `json:"id_request_friendship"`
	Status                int `json:"status"` //1=accept, 2=reject
}

type UpdateStatusInput struct {
	Status int `json:"status"` //1=accept, 2=reject
}

// GET /statuss
// Get all statuss
func FindStatuss(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var statuss []models.Status_friendship
	db.Find(&statuss)

	c.JSON(http.StatusOK, gin.H{"data": statuss})
}

// POST /statuss
// Create new statuss
func CreateStatus(c *gin.Context) {
	// Validate input
	var input CreateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var statuss models.Status_friendship

	db := c.MustGet("db").(*gorm.DB)

	// Get Id_request_friendship
	var requests models.Request_friendship
	if err := db.Where("id = ? AND status=1", input.Id_request_friendship).Find(&requests).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Request with status=1 not found!!"})
		return
	}

	//validate existing requested friendship
	if err := db.Where("id_request_friendship = ? AND status = 1", input.Id_request_friendship).First(&statuss).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already existing status accept/reject friendship from that ID's!!"})
		return
	}

	if input.Status < 1 || input.Status > 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status 1 untuk accept, status 2 untuk reject"})
		return
	}

	// Create status
	statuss2 := models.Status_friendship{Id_request_friendship: input.Id_request_friendship, Status: input.Status}

	db.Create(&statuss2)

	c.JSON(http.StatusOK, gin.H{"data": statuss2})
}

// GET /statuss/:id
// Find a statuss
func FindStatus(c *gin.Context) { // Get model if exist
	var statuss models.Status_friendship

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&statuss).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": statuss})
}

// PATCH /statuss/:id
// Update a statuss
func UpdateStatus(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var statuss models.Status_friendship
	if err := db.Where("id = ?", c.Param("id")).First(&statuss).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Status < 1 || input.Status > 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status 1 untuk accept, status 2 untuk reject"})
		return
	}

	var updatedInput models.Status_friendship
	updatedInput.Status = input.Status

	db.Model(&statuss).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": statuss})
}

// DELETE /statuss/:id
// Delete a statuss
func DeleteStatus(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var statuss models.Status_friendship
	if err := db.Where("id = ?", c.Param("id")).First(&statuss).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&statuss)

	c.JSON(http.StatusOK, gin.H{"data": "Delete success!"})
}
