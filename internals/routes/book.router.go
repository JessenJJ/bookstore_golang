package routes

import (
	"golang2bookst/internals/handler"
	"golang2bookst/internals/middleware"
	"golang2bookst/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBookRouter(router *gin.Engine, db *sqlx.DB) {
	bookRouter := router.Group("/book")
	bookRepo := repositories.InitBookRepo(db)
	bookHandler := handler.InitBookHandler(bookRepo)

	// localhost:8000/book
	bookRouter.GET("", middleware.CheckToken, bookHandler.GetBooks)
	// localhost:8000/book/newbook
	bookRouter.POST("/newbook", middleware.CheckToken, bookHandler.CreateBooks)
	// Localhost:8000/book/id method delete
	bookRouter.DELETE("/:id", middleware.CheckToken, bookHandler.DeleteTheBook)
	// Localhost:8000/book/id method patch
	bookRouter.PATCH("/:id", middleware.CheckToken, bookHandler.UpdateTheBook)
	// Localhost:8000/book/id method get
	bookRouter.GET("/:id", middleware.CheckToken, bookHandler.GetBookById)

}
