package main

import (
	"api/config"
	"api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.GetConfig()
	db := db.NewDB(config)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": db,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}