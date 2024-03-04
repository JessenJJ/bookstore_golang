package handler

import (
	"golang2bookst/internals/models"
	"golang2bookst/internals/repositories"
	"golang2bookst/pkg"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	*repositories.AuthRepo
}

func InitAuthHandler(a *repositories.AuthRepo) *AuthHandler {

	return &AuthHandler{a}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	// ambil body
	body := &models.AuthModel{}
	if err := ctx.ShouldBind(body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		// menggunakan return agar memberhentikan handlernya saja, bukan servernya
		return
	}
	result, err := a.FindByEmail(*body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//pengecekan duplikat email
	if len(result) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is already registered!",
		})
		return
	}
	hash, err := argon2id.CreateHash(body.Password, argon2id.DefaultParams)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	if err := a.SaveUser(models.AuthModel{
		// Id:       body.Id,
		Email:    body.Email,
		Password: hash,
	}); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success Registered",
	})
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	body := models.AuthModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := a.FindByEmail(body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email not registered",
		})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(body.Password, result[0].Password)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !match {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "bad credentials",
		})
		return
	}

	payload := pkg.NewPayLoad(body.Email)
	token, err := payload.CreateToken()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successfulll",
		"token":   token,
	})
}
