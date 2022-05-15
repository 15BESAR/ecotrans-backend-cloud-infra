package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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
// buy voucher
func BuyVoucher(c *gin.Context) {
	fmt.Println("POST /voucher")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Data:", string(body))
	c.JSON(http.StatusOK, gin.H{"msg": "Purchase Successful !", "pointRemaining": 2000})
}
