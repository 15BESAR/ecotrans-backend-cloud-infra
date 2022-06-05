package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thoas/go-funk"
)

type ReturnPurchase struct {
	ID          uint
	BuyTime     time.Time
	Voucher     models.Voucher
	IsPurchased bool
}

// GET user/purchase/:userId
// Get Purchase based on user id
func GetPurchase(c *gin.Context) {
	var puchasedVoucher []models.VoucherPurchased
	if err := models.Db.Preload("Voucher").Where("user_id = ?", c.Param("userId")).Find(&puchasedVoucher).Error; err != nil {
		fmt.Println(puchasedVoucher)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error,
		})
		return
	}

	arrReturnPurchaseVoucer := funk.Map(puchasedVoucher, func(lesr models.VoucherPurchased) ReturnPurchase {
		var temp ReturnPurchase
		temp.ID = lesr.ID
		temp.BuyTime = lesr.CreatedAt
		temp.Voucher = lesr.Voucher
		temp.IsPurchased = lesr.IsPurchased
		return temp
	})
	c.JSON(http.StatusOK, gin.H{
		"error":            false,
		"userId":           c.Param("userId"),
		"voucherPurchased": arrReturnPurchaseVoucer,
	})

}
