package main

import (
	"log"
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/controllers"
	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load Environment
	env = LoadEnvironment()

	r := gin.Default()
	r.SetTrustedProxies(nil)
	// Connect to database

	var err error

	models.Db, err = models.ConnectDatabase(env.dbUser, env.dbPass, env.dbName, env.dbTCPHost, env.dbPort)
	if err != nil {
		log.Fatalf("Database not opened")
	}

	// Middleware
	r.Use(TokenAuthMiddleware())

	// Routes
	r.GET("/", root)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/login/refresh", controllers.RefreshToken)
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
