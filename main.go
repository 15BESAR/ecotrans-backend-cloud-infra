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
	models.Db.AutoMigrate(&models.Voucher{})
	models.Db.AutoMigrate(&models.VoucherPurchased{})
	models.Db.AutoMigrate(&models.Partner{})
	models.Db.AutoMigrate(&models.Purchase{})

	if err != nil {
		log.Fatalf("Database not opened")
	}

	// Middleware
	r.Use(TokenAuthMiddleware())

	// API ROUTE
	// Root and Version
	r.GET("/", root)
	r.GET("/version", getVersion)
	// Auth User
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/refresh", controllers.RefreshTokenUser)
	// CRUD User
	r.GET("/users", controllers.FindUsers)
	r.GET("/user/:userId", controllers.FindUserById)
	r.PUT("/user/:userId", controllers.UpdateUserById)
	r.DELETE("/user/:userId", controllers.DeleteUserById)

	r.GET("/user/purchase/:userId", controllers.GetPurchase)
	// GMAPS API
	r.POST("/autocomplete", controllers.AutocompleteLocation)
	r.POST("/routes", controllers.FindRoutes)
	// CRUD Journeys
	r.GET("/journeys", controllers.FindAllJourneys)
	r.POST("/journey", controllers.AddJourney)
	r.GET("/journey/:journeyId", controllers.FindJourneyById)
	// Forecasting API
	r.GET("/forecast", controllers.FindForecast)
	// CRUD Vouchers
	r.GET("/vouchers", controllers.FindVouchers)
	r.POST("/voucher", controllers.AddVoucher)
	r.GET("/voucher/:voucherId", controllers.FindVoucherByVoucherId)
	r.PUT("/voucher/:voucherId", controllers.UpdateVoucherById)
	r.DELETE("/voucher/:voucherId", controllers.DeleteVoucherById)
	// Purchase Voucher API
	r.GET("/purchases", controllers.FindAllPurchases)
	r.POST("/purchase", controllers.PurchaseVoucher)
	r.GET("/purchase/:purchaseId", controllers.FindPurchaseById)
	r.DELETE("/purchase/:purchaseId", controllers.DeletePurchaseById)
	// Partner Auth API
	r.POST("/company/register", controllers.RegisterPartner)
	r.POST("/company/login", controllers.LoginPartner)
	r.GET("/company/refresh", controllers.RefreshTokenPartner)
	// CRUD Partner
	r.GET("/partners", controllers.FindPartners)
	r.GET("/partner/:partnerId", controllers.FindPartnerById)
	r.DELETE("/partner/:partnerId", controllers.DeletePartnerById)
	// Handler for wrong Path
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Path not defined, try reading the docs in ReadMe or contact CC people"})
	})

	// Run the server
	r.Run()
}

func root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Ecotrans GO Backend API"})
}

func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "0.51", "lastUpdate": "May 21"})
}
