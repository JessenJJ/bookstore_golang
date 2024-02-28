package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// buat rute nya
	router.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World")
	})

	return router

}
