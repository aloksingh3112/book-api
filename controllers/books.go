package controllers

import (
	"net/http"

	"github.com/aloksingh3112/books-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var books []models.Book
	db.Find((&books))
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input CreateBookInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Author: input.Author}
	db.Create(&book)
	c.JSON(http.StatusOK, gin.H{"book": book})
}

func GetBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	err := db.Where("id=?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	err := db.Where("id=?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var UpdateBookInput UpdateBookInput
	error := c.ShouldBindJSON(&UpdateBookInput)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": error.Error()})
		return
	}

	db.Model(&book).Update(&UpdateBookInput)
	c.JSON(http.StatusOK, gin.H{"data": UpdateBookInput})
}

func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book

	err := db.Where("id=?", c.Param("id")).First(&book).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Data not found"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
