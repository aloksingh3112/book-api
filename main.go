package main

import (
	"github.com/aloksingh3112/books-api/controllers"
	"github.com/aloksingh3112/books-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	// })
	db := models.DbConnection()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBooks)
	r.GET("/books/:id", controllers.GetBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()

}
