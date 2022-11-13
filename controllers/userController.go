package controllers

import (
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

func GetAllUsers(c *gin.Context) {
	users := []models.User{}

	if err := config.DB.Model(&users).Preload("Issue").Find(&users).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetAllUsersDepartments(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
	userBody := models.CreateUserBody{}

	if err := c.ShouldBindJSON(&userBody); err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	newUser := models.User{FirstName: userBody.FirstName, LastName: userBody.LastName, PersonalNumber: userBody.PersonalNumber, ArmyEmail: userBody.ArmyEmail, Department: userBody.Department}

	if err := config.DB.Create(&newUser).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	if err := ifUserExistsThrowErr(newUser.ArmyEmail, c); err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	tokenStr, err := generateToken(newUser)

	if err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.SetCookie("x-auth-token", tokenStr, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusCreated, tokenStr)
}

func DeleteUser(c *gin.Context) {
	id := helpers.GetParamData(c, "id")
	deletedUser := findUserById(id, c)

	if err := config.DB.Delete(&deletedUser).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, deletedUser)
}

// Helpers

func generateToken(u models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.UserID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func findUserById(id string, c *gin.Context) models.User {
	user := models.User{}

	if err := config.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
	}

	return user
}

func findUserByEmail(email string, c *gin.Context) models.User {
	user := models.User{}

	if err := config.DB.Where("army_email = ?", email).First(&user).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
	}

	return user
}

func ifUserExistsThrowErr(email string, c *gin.Context) error {
	user := models.User{}

	result := config.DB.Where("army_email = ?", email).First(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return nil
	}

	return nil
}
