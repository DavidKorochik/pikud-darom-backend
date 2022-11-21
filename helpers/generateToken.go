package helpers

import (
	"os"
	"time"

	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(u models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.UserID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
