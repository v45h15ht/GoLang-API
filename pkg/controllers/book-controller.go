package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/v45h15h7/REWARD_SYSTEM/pkg/models"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	c.IndentedJSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	c.IndentedJSON(http.StatusOK, bookDetails)
}

func CreateBook(c *gin.Context) {
	CreateBook := &models.Book{}

	if err := c.BindJSON(&CreateBook); err != nil {
		return
	}

	b := CreateBook.CreateBook()
	c.IndentedJSON(http.StatusCreated, b)
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	c.IndentedJSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var updateBook = &models.Book{}

	if err := c.BindJSON(&updateBook); err != nil {
		return
	}

	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	booksDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	c.IndentedJSON(http.StatusAccepted, booksDetails)
}
