package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DavidKorochik/pikud-darom-backend/config"
	"github.com/DavidKorochik/pikud-darom-backend/helpers"
	"github.com/DavidKorochik/pikud-darom-backend/initializers"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func init() {
	initializers.LoadEnvVariables()
}

func AuthToken(c *gin.Context) {
	user := models.User{}
	tokenSecret := os.Getenv("JWT_SECRET")
	tokenStr, err := c.Cookie("x-auth-token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authorized"})
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header)
		}

		return []byte(tokenSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if time.Now().Unix() > claims["exp"].(int64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if err := config.DB.Where("user_id = ?", claims["sub"]).First(&user).Error; err != nil {
			helpers.DisplayErrorMsg(c, err)
			return
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
