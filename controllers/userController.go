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
	"golang.org/x/crypto/bcrypt"
)

func init() {
	initializers.LoadEnvVariables()
}

func GetAllUsers(c *gin.Context) {
	users := []models.User{}

	if err := config.DB.Find(&users).Error; err != nil {
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

	newUser := models.User{FirstName: userBody.FirstName, LastName: userBody.LastName, PersonalNumber: userBody.PersonalNumber, ArmyEmail: userBody.ArmyEmail, Department: userBody.Department, Issues: userBody.Issues}

	if err := config.DB.Create(&newUser).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

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
	}

	tokenStr, tokenErr := generateToken(userFoundWithEmail)

	if tokenErr != nil {
		helpers.DisplayErrorMsg(c, tokenErr)
	}

	c.JSON(http.StatusOK, tokenStr)
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
