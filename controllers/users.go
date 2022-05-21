package controllers

import (
	"fmt"
	"net/http"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// GET /users
// get all users data
func FindUsers(c *gin.Context) {
	fmt.Println("GET /users")
	var users []models.User
	models.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})

}

// GET /user/:userid
// get user data with userid
func FindUserById(c *gin.Context) {
	fmt.Println("GET /user/:userId")
	var user models.User
	if err := models.Db.Where("user_id = ?", c.Param("userId")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(http.StatusOK, user)

}

// DELETE /user/:userid
// Delete user data with userid
func DeleteUserById(c *gin.Context) {
	fmt.Println("DELETE /user/:userId")
	var user models.User
	if err := models.Db.Where("user_id = ?", c.Param("userId")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	models.Db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"msg": "user deleted"})

}

// PUT /user/:userid
// update user data with userid
func UpdateUserById(c *gin.Context) {
	fmt.Println("GET /user/:userid")
	// Bind body

	// check if data valid
	// Update to DB
	// Return User Information
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("test"))
}
