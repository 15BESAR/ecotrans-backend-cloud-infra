package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/15BESAR/ecotrans-backend-cloud-infra/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kr/pretty"
)

// GET /users
// get all users data
func FindUsers(c *gin.Context) {
	fmt.Println("GET /users")
	var users []models.User
	models.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "success",
		"users": users,
	})

}

// GET /user/:userid
// get user data with userid
func FindUserById(c *gin.Context) {
	fmt.Println("GET /user/:userId")
	var user models.User
	if err := models.Db.Where("user_id = ?", c.Param("userId")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "User not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"user":  user})

}

// DELETE /user/:userid
// Delete user data with userid
func DeleteUserById(c *gin.Context) {
	fmt.Println("DELETE /user/:userId")
	var user models.User
	if err := models.Db.Where("user_id = ?", c.Param("userId")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "User not found!",
		})
		return
	}
	models.Db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "user deleted",
	})

}

// PUT /user/:userid
// update user data with userid
func UpdateUserById(c *gin.Context) {
	fmt.Println("GET /user/:userid")
	var input models.UserUpdate
	var user models.User

	// Find user
	if err := models.Db.Where("user_id = ?", c.Param("userId")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "User not found!",
		})
		return
	}
	// Bind body, Validate Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	// check if data valid
	if err := validateUpdateUserInput(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	pretty.Println(input)
	// Update to DB
	models.Db.Model(&user).Updates(structs.Map(input))
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "User get",
		"user":  user})

}

func validateUpdateUserInput(input *models.UserUpdate) error {
	seds := reflect.ValueOf(input).Elem()
	for i := 0; i < seds.NumField(); i++ {
		fieldValue := seds.Field(i).String()
		if fieldValue == "" {
			return fmt.Errorf("%s is  empty", seds.Type().Field(i).Name)
		}
	}
	// gender check
	if input.Gender != "m" && input.Gender != "f" {
		return errors.New("gender not in input")
	}

	// check category
	category := [5]string{"Electronic", "Fashion", "Food", "Transportation", "Ecommerce"}
	voucherCatSplitter := strings.Split(input.VoucherInterest, ",")
	fmt.Println(voucherCatSplitter)
	for _, voucName := range voucherCatSplitter {
		isCategoryFound := false
		for _, catName := range category {
			if voucName == catName {
				isCategoryFound = true
			}
		}
		if !isCategoryFound {
			return fmt.Errorf("%s is not in category", voucName)
		}
	}

	// check user preference
	preference := [4]string{"walking", "bicycling", "driving", "transit"}
	isVehicleFound := false
	for _, prefName := range preference {
		if prefName == input.Vehicle {
			isVehicleFound = true
		}
	}
	if !isVehicleFound {
		return fmt.Errorf("%s is not in preferences", input.Vehicle)
	}

	return nil
}
