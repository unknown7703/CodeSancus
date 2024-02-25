package controllers

import "github.com/gin-gonic/gin"

func Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"msg":    "Admin Created",
	})
}
