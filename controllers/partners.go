package controllers

import (
	"fmt"
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
)

// GET /partners
// find all partners
func FindPartners(c *gin.Context) {
	var partners []models.Partner
	models.Db.Find(&partners)
	c.JSON(http.StatusOK, gin.H{"partners": partners})
}

// GET /Partner/:partnerid
// Get Partner By ID
func FindPartnerById(c *gin.Context) {
	fmt.Println("GET /user/:userId")
	var user models.User
	if err := models.Db.Where("user_id = ?", c.Param("userId")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DELETE /Partner/:partnerid
// Delete Partner By ID
func DeletePartnerById(c *gin.Context) {
	fmt.Println("DELETE /user/:userId")
	var user models.User
	if err := models.Db.Where("user_id = ?", c.Param("userId")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	models.Db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"msg": "user deleted"})
}
