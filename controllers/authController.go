package controllers

import (
	"net/http"

	"github.com/DavidKorochik/pikud-darom-backend/helpers"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LogInUser(c *gin.Context) {
	user := models.LogInUserBody{}

	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	userFoundWithEmail := findUserByEmail(user.ArmyEmail, c)

	err := bcrypt.CompareHashAndPassword([]byte(userFoundWithEmail.PersonalNumber), []byte(user.PersonalNumber))

	if err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	tokenStr, tokenErr := helpers.GenerateToken(userFoundWithEmail)

	if tokenErr != nil {
		helpers.DisplayErrorMsg(c, tokenErr)
		return
	}

	c.Header("x-auth-token", tokenStr)

	c.JSON(http.StatusOK, tokenStr)
}
