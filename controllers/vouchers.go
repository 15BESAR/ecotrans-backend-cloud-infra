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
	var vouchers []models.Voucher
	models.Db.Find(&vouchers)
	c.JSON(http.StatusOK, gin.H{"vouchers": vouchers})
}

// POST /voucher
// Add voucher by partner
func AddVoucher(c *gin.Context) {
	var voucher models.Voucher
	// bind body
	if err := c.ShouldBindJSON(&voucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if partner exist
	var partner models.Partner
	if err := models.Db.Where("partner_id = ?", voucher.PartnerID).First(&partner).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partner not found!"})
		return
	}
	// Check if voucher name is same and partner name is the same, then don't add
	if err := models.Db.Where("partner_name = ? AND voucher_name = ?", voucher.PartnerName, voucher.VoucherName).First(&voucher).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Voucher Already Exist!"})
		return
	}
	// in the future maybe add checker to prevent double data
	// add journey
	result := models.Db.Create(&voucher)
	pretty.Println(voucher)
	pretty.Println(result.Error)
	pretty.Println(result.RowsAffected)
	c.JSON(http.StatusOK, voucher)
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
