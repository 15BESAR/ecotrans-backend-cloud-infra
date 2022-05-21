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
	var purchase models.Purchase
	// bind body
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if user exist
	var user models.User
	if err := models.Db.Where("user_id = ?", purchase.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	// Check if voucher exist
	var voucher models.Voucher
	if err := models.Db.Where("voucher_id = ?", purchase.VoucherID).First(&voucher).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Voucher not found!"})
		return
	}
	// Check if stock >= buy qty
	if voucher.Stock < purchase.BuyQuantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Voucher stock is not enough!"})
		return
	}
	// Check if user points > buyqty*price
	if user.Points < purchase.BuyQuantity*voucher.Price {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User points is not enough!"})
		return
	}
	// Perform Purchase Transaction
	voucher.Stock -= purchase.BuyQuantity
	user.Points -= (purchase.BuyQuantity * voucher.Price)
	// Update Voucher stock to DB
	models.Db.Save(voucher)
	// Update User Point to DB
	models.Db.Save(user)
	// Add Purchase history to DB
	models.Db.Create(&purchase)
	// Create purchase receipt
	receipt := models.PurchaseReceipt{Purchase: purchase, UserPointsRemaining: user.Points, VoucherStockRemaining: voucher.Stock}
	c.JSON(http.StatusOK, receipt)
}
