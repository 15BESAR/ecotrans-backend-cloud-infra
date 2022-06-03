package controllers

import (
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
)

// GET /partners
// find all partners
func FindPartners(c *gin.Context) {
	var partners []models.Partner
	models.Db.Find(&partners)
	c.JSON(http.StatusOK, gin.H{
		"error":    false,
		"msg":      "partnerts found",
		"partners": partners})
}

// GET /Partner/:partnerid
// Get Partner By ID
func FindPartnerById(c *gin.Context) {
	var partner models.Partner
	if err := models.Db.Where("partner_id = ?", c.Param("partnerId")).First(&partner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Partner not found!"})
		return
	}

	c.JSON(http.StatusOK, partner)
}

// DELETE /Partner/:partnerid
// Delete Partner By ID
func DeletePartnerById(c *gin.Context) {
	var partner models.Partner
	if err := models.Db.Where("partner_id = ?", c.Param("partnerId")).First(&partner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Partner not found!"})
		return
	}
	models.Db.Delete(&partner)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "user deleted"})
}
