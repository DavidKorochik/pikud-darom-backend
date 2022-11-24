package controllers

import (
	"net/http"
	"time"

	"github.com/DavidKorochik/pikud-darom-backend/config"
	"github.com/DavidKorochik/pikud-darom-backend/helpers"
	"github.com/DavidKorochik/pikud-darom-backend/initializers"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables("config.env")
}

func GetAllUsers(c *gin.Context) {
	users := []models.User{}

	if err := config.DB.Preload("Issues").Find(&users).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	for _, user := range users {
		if err := config.DB.Model(&models.Issue{}).Where("user_id = ?", user.UserID).Find(&user.Issues).Error; err != nil {
			helpers.DisplayErrorMsg(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, users)
}

func GetAllUsersDepartments(c *gin.Context) {
	departments := []string{}

	if err := config.DB.Model(&models.User{}).Select("department").Find(&departments).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, departments)
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

	tokenStr, err := helpers.GenerateToken(newUser)

	if err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.SetCookie("x-auth-token", tokenStr, int(time.Now().Unix()), "", "", false, true)

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
