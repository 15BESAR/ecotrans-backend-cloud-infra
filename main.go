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
	models.APPLICATION_NAME = env.appName
	models.JWT_SIGNATURE_KEY = env.sigKeyJwt

	r := gin.Default()
	r.SetTrustedProxies(nil)
	// Connect to database

	var err error

	models.Db, err = models.ConnectDatabase(env.dbUser, env.dbPass, env.dbName, env.dbTCPHost, env.dbPort)
	models.Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Database not opened")
	}

	// Middleware
	r.Use(TokenAuthMiddleware())

	// Routes
	r.GET("/", root)
	r.GET("/version", getVersion)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/login/refresh", controllers.RefreshTokenUser)
	r.GET("/users", controllers.FindUsers)
	r.GET("/user/:userId", controllers.FindUserById)
	r.PUT("/user/:userId", controllers.UpdateUserById)
	r.DELETE("/user/:userId", controllers.DeleteUserById)
	r.GET("/autocomplete", controllers.AutocompleteLocation)
	r.GET("/routes", controllers.FindRoutes)
	r.GET("/journeys", controllers.FindAllJourneys)
	r.POST("/journey", controllers.AddJourney)
	r.GET("/vouchers", controllers.FindVouchers)
	r.POST("/voucher", controllers.BuyVoucher)
	r.POST("/partner/register", controllers.RegisterPartner)
	r.POST("/partner/login", controllers.LoginPartner)
	r.GET("/partner/login/refresh", controllers.RefreshTokenPartner)
	r.GET("/partners", controllers.FindPartners)
	r.GET("/partner/:partnerId", controllers.FindPartnerById)

	// Run the server
	r.Run()
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Ecotrans GO Backend API"})
}

func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "0.5", "lastUpdate": "May 22"})
}
