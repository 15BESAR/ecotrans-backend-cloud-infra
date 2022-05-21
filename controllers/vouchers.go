package controllers

import (
	"fmt"
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
)

// GET /vouchers
// find all voucher
func FindVouchers(c *gin.Context) {
	fmt.Println("POGET /vouchers")
	var stringData string = `
	{
		"vouchers":[
			{
				"voucherId":"c2njbn4",
				"name":"Free 3 donut",
				"category":"food and beverages",
				"description": "free 3 donut berlaku untuk seluruh cabang Jco",
				"image": "https://storage.googleapis.com/voucher-images-2909/jco.jpg",
				"partner": "Jco",
				"price" : 1000,
				"expire date": "2018-12-10T16:49:51.141Z"
			}
		]
	}
	`
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(stringData))
}

// POST /voucher
// Add voucher by partner
func AddVoucher(c *gin.Context) {
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

// GET /voucher/:voucherId
// GET Partner By ID
func FindVoucherById(c *gin.Context) {
	var partner models.Partner
	if err := models.Db.Where("partner_id = ?", c.Param("partnerId")).First(&partner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partner not found!"})
		return
	}
	models.Db.Delete(&partner)

	c.JSON(http.StatusOK, gin.H{"msg": "user deleted"})
}

// PUT /voucher/:userid
// update voucher data by partner
func UpdateVoucherById(c *gin.Context) {
	fmt.Println("GET /user/:userid")
	var input models.UserUpdate
	var user models.User

	// Find user
	if err := models.Db.Where("user_id = ?", c.Param("userId")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	// Bind body, Validate Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// check if data valid
	if err := validateUpdateUserInput(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pretty.Println(input)
	// Update to DB
	models.Db.Model(&user).Updates(structs.Map(input))
	c.JSON(http.StatusOK, user)

}

// DELETE /voucher/:voucherId
// Delete Partner By ID
func DeleteVoucherById(c *gin.Context) {
	var partner models.Partner
	if err := models.Db.Where("partner_id = ?", c.Param("partnerId")).First(&partner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partner not found!"})
		return
	}
	models.Db.Delete(&partner)

	c.JSON(http.StatusOK, gin.H{"msg": "user deleted"})
}
