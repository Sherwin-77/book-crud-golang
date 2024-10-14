package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/book-crud-golang/pkg/controllers"
)

func RegisterRoutes(group *echo.Group) {
	group.GET("/book", controllers.GetBooks)
	group.GET("/book/:id", controllers.GetBookByID)
	group.POST("/book", controllers.CreateBook)
	group.PUT("/book/:id", controllers.UpdateBook)
	group.DELETE("/book/:id", controllers.DeleteBook)
}
