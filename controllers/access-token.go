package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AccessToken(c *gin.Context) {

	var body struct {
		AppName string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "Gagal memuat Request Body",
		})
		return
	}

	appName := os.Getenv("APP_NAME")

	if body.AppName != appName {
		c.JSON(http.StatusBadRequest, models.ResponseOnlyMessage{
			Code:    400,
			Message: "appName tidak valid",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": appName,
		"exp": time.Now().Add(time.Hour * 6).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseOnlyMessage{
			Code:    500,
			Message: "Terjadi kesalahan saat membuat Token",
		})
		return
	}

	var data struct {
		AppName     string `json:"appName"`
		AccessToken string `json:"accessToken"`
	}

	data.AppName = appName
	data.AccessToken = tokenString

	c.JSON(http.StatusOK, models.ResponseWithData{
		Code:    200,
		Message: "Access Token",
		Data:    data,
	})

}
