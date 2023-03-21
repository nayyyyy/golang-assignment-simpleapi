package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nayyyyy/golang-assignment/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.IndexBook)
	router.GET("/books/:id", controllers.ShowBook)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
