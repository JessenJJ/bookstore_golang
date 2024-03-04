package routes

import (
	"golang2bookst/internals/handler"
	"golang2bookst/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAuthRouter(router *gin.Engine, db *sqlx.DB) {
	// bikin subrouter
	authRouter := router.Group("/auth")

	//construct handler
	authRepo := repositories.InitAuthRepo(db)
	authHandler := handler.InitAuthHandler(authRepo)

	// bikin rute
	// localhost:8000/auth/newuser
	authRouter.POST("/newuser", authHandler.Register)
	// localhost:8000/auth
	authRouter.POST("", authHandler.Login)
}
