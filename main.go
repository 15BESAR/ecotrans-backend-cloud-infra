package main

import (
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/controllers"
	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/", root)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/users", controllers.FindUsers)
	r.GET("/user/:userId", controllers.FindUserById)
	r.PUT("/user/:userId", controllers.UpdateUserById)
	r.GET("/autocomplete", controllers.AutocompleteLocation)
	r.GET("/routes", controllers.FindRoutes)
	r.POST("/finish", controllers.AddJourney)
	r.GET("/vouchers", controllers.FindVouchers)
	r.POST("/voucher", controllers.BuyVoucher)

	// Run the server
	r.Run()
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Ecotrans GO Backend API"})
}
