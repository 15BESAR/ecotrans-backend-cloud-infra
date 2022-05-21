package controllers

import (
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
)

// GET /purchases
// GET ALL Purchases
func FindAllPurchases(c *gin.Context) {
	var journey models.Journey
	// bind body
	if err := c.ShouldBindJSON(&journey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if user exist
	var user models.User
	if err := models.Db.Where("user_id = ?", journey.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	// in the future maybe add checker to prevent double data
	// add journey
	models.Db.Create(&journey)
	c.JSON(http.StatusOK, journey)
}

// POST /purchase
// Purchase a Voucher
func PurchaseVoucher(c *gin.Context) {
	var journey models.Journey
	// bind body
	if err := c.ShouldBindJSON(&journey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if user exist
	var user models.User
	if err := models.Db.Where("user_id = ?", journey.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	// in the future maybe add checker to prevent double data
	// add journey
	models.Db.Create(&journey)
	c.JSON(http.StatusOK, journey)
}
