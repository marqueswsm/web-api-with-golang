package controllers

import "github.com/gin-gonic/gin"

func ShowUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"value": "ok",
	})
}
