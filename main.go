package main

import (
	"github.com/aloksingh3112/books-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })
	db := models.DbConnection()
	r.Use
	r.Run()

}
