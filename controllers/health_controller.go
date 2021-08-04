package controllers

import "github.com/gin-gonic/gin"

func SendStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"API": "Up and Running",
	})
}
