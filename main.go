package main

import (
	"github.com/gin-gonic/gin"

	"github.com/negiseijin/ddd-golang/infrastructure"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.Run()
	d := infrastructure.NewDB()
	db := d.Connect()
}
