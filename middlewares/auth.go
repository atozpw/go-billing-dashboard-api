package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/atozpw/go-billing-dashboard-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *gin.Context) {

	var responseCode int = 401

	tokenString := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)

	if tokenString == "" {

		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseOnlyMessage{
			Code:    responseCode,
			Message: "Request tidak valid",
		})

	} else {

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			appName := os.Getenv("APP_NAME")

			if claims["sub"] != appName {
				c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseOnlyMessage{
					Code:    responseCode,
					Message: "Token tidak terdaftar pada Pengguna",
				})
			}

			c.Next()

		} else if errors.Is(err, jwt.ErrTokenExpired) {

			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseOnlyMessage{
				Code:    responseCode,
				Message: "Token kedaluwarsa",
			})

		} else {

			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseOnlyMessage{
				Code:    responseCode,
				Message: "Token tidak valid",
			})

		}

	}

}
