package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marqueswsm/web-api-with-golang/database"
	"github.com/marqueswsm/web-api-with-golang/models"
)

func ShowUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID must to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var user models.User
	err = db.First(&user, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "User not found: " + err.Error(),
		})

		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to parse user to json",
		})

		return
	}

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create an user" + err.Error(),
		})

		return
	}

	c.JSON(200, user)
}

func ShowAllUsers(c *gin.Context) {
	db := database.GetDatabase()
	var users []models.User
	err := db.Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func UpdateUser(c *gin.Context) {
	db := database.GetDatabase()
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to parse user to json",
		})

		return
	}

	err = db.Save(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create an user" + err.Error(),
		})

		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID must to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var user models.User
	err = db.Delete(&user, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "User not deleted",
		})

		return
	}

	c.JSON(204, user)
}
