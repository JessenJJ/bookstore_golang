package middleware

import (
	"golang2bookst/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckToken(ctx *gin.Context) {
	// ambil header authorization
	bearerToken := ctx.GetHeader("Authorization")
	// Bearer token
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Please login first",
		})
		return
	}
	if !strings.Contains(bearerToken, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Authorization",
		})
		return
	}

	// ambil token
	token := strings.Replace(bearerToken, "Bearer ", "", -1)

	_, err := pkg.VerifyToken(token)
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Please Login Again, session expired",
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// lanjut ke handler selanjutnya
	ctx.Next()

}
