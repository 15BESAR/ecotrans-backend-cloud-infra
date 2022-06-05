package controllers

import (
	"net/http"
	"strings"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
)

// GET /purchases
// GET ALL Purchases
func FindAllPurchases(c *gin.Context) {
	var purchases []models.Purchase
	models.Db.Find(&purchases)
	// Check if there's query
	queryUser, isQueryUser := c.GetQuery("user")
	if isQueryUser {
		queryUser = strings.ToLower(queryUser)
		temp := make([]models.Purchase, 0)
		for _, purchase := range purchases {
			if strings.ToLower(purchase.UserID) == queryUser {
				temp = append(temp, purchase)
			}
		}
		purchases = temp
	}
	c.JSON(http.StatusOK, gin.H{
		"error":     false,
		"msg":       "get purchases",
		"purchases": purchases,
	})
}

// POST /purchase
// Purchase a Voucher
func PurchaseVoucher(c *gin.Context) {
	var purchase models.Purchase
	var voucherPurchased models.VoucherPurchased
	// bind body
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	// Check if user exist
	var user models.User
	if err := models.Db.Where("user_id = ?", purchase.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "User not found!",
		})
		return
	}
	// Check if voucher exist
	var voucher models.Voucher
	if err := models.Db.Where("voucher_id = ?", purchase.VoucherID).First(&voucher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Voucher not found!",
		})
		return
	}
	// Check if stock >= buy qty
	if voucher.Stock < purchase.BuyQuantity {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Voucher stock is not enough!",
		})
		return
	}
	// Check if user points > buyqty*price
	if user.Points < purchase.BuyQuantity*voucher.Price {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "User points is not enough!",
		})
		return
	}

	// Perform Purchase Transaction
	voucher.Stock -= purchase.BuyQuantity
	user.Points -= (purchase.BuyQuantity * voucher.Price)
	// Update Voucher stock to DB
	models.Db.Save(&voucher)
	// Update User Point to DB
	models.Db.Save(&user)
	// Add Purchase history to DB
	models.Db.Create(&purchase)

	// create voucherPurchased
	voucherPurchased.UserID = user.UserID
	voucherPurchased.VoucherID = purchase.VoucherID
	models.Db.Create(&voucherPurchased)

	// Create purchase receipt
	receipt := models.PurchaseReceipt{Purchase: purchase, UserPointsRemaining: user.Points, VoucherStockRemaining: voucher.Stock}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"receipt": receipt,
	})
}

// GET /purchase/:purchaseId
// GET purchase By purchaseId
func FindPurchaseById(c *gin.Context) {
	var purchase models.Purchase
	if err := models.Db.Where("purchase_id = ?", c.Param("purchaseId")).First(&purchase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Purchase not found!",
		})
		return
	}
	c.JSON(http.StatusOK, purchase)
}

// DELETE /voucher/:voucherId
// Delete Partner By ID
func DeletePurchaseById(c *gin.Context) {
	var purchase models.Purchase
	if err := models.Db.Where("purchase_id = ?", c.Param("purchaseId")).First(&purchase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Purchase not found!",
		})
		return
	}
	models.Db.Delete(&purchase)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Purchase deleted",
	})
}
