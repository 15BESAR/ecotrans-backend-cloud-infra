package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /partners
// find all partners
func FindPartners(c *gin.Context) {
	fmt.Println("GET /partners")

	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(`{"msg":"GET /partners"}`))
}

// GET /Partner/:partnerid - Get Partner By ID
// Register Partner
func FindPartnerById(c *gin.Context) {
	fmt.Println("GET /partner/:partnerid")
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(`{"msg":"GET /partners"}`))

}
