package controllers

import (
	"net/http"
	"strings"

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
	// Check if there's query
	queryCompany, isQueryCompany := c.GetQuery("company")
	if isQueryCompany {
		queryCompany = strings.ToLower(queryCompany)
		temp := make([]models.Voucher, 0)
		for _, voucher := range vouchers {
			if strings.ToLower(voucher.PartnerName) == queryCompany {
				temp = append(temp, voucher)
			}
		}
		vouchers = temp
	}
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
// GET Partner By Voucher ID
func FindVoucherByVoucherId(c *gin.Context) {
	var voucher models.Voucher
	if err := models.Db.Where("voucher_id = ?", c.Param("voucherId")).First(&voucher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Partner not found!"})
		return
	}
	c.JSON(http.StatusOK, voucher)
}

// PUT /voucher/:userid
// update voucher data by partner
func UpdateVoucherById(c *gin.Context) {
	var input models.Voucher
	var voucher models.Voucher

	// Find voucher
	if err := models.Db.Where("voucher_id = ?", c.Param("voucherId")).First(&voucher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Voucher not found!"})
		return
	}
	// Bind body, Validate Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// check if data valid
	if err := validateUpdateVoucherInput(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.VoucherID = c.Param("voucherId")
	// Update to DB
	models.Db.Model(&voucher).Updates(structs.Map(input))
	c.JSON(http.StatusOK, voucher)

}

// DELETE /voucher/:voucherId
// Delete Partner By ID
func DeleteVoucherById(c *gin.Context) {
	var voucher models.Voucher
	if err := models.Db.Where("voucher_id = ?", c.Param("voucherId")).First(&voucher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Voucher not found!"})
		return
	}
	models.Db.Delete(&voucher)

	c.JSON(http.StatusOK, gin.H{"msg": "Voucher deleted"})
}

func validateUpdateVoucherInput(input *models.Voucher) error {
	return nil
}
