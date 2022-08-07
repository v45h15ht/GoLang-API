package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/v45h15h7/REWARD_SYSTEM/pkg/controllers"
)

func main() {
	router := gin.Default()
	router.POST("/book/", controllers.CreateBook)
	router.GET("/book/", controllers.GetBook)
	router.GET("/book/:bookId", controllers.GetBookById)
	router.PUT("/book/:bookId", controllers.UpdateBook)
	router.DELETE("/book/:bookId", controllers.DeleteBook)
	router.Run("localhost:9090")
}
