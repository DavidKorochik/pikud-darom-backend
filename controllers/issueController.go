package controllers

import (
	"net/http"

	"github.com/DavidKorochik/pikud-darom-backend/config"
	"github.com/DavidKorochik/pikud-darom-backend/helpers"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllIssues(c *gin.Context) {
	_, err := c.Cookie("x-auth-token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	issues := []models.Issue{}

	if err := config.DB.Find(&issues).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, issues)
}

func FilterIssuesByDepartment(c *gin.Context) {
	queryData := helpers.GetQueryData(c, "department")
	issues := []models.Issue{}

	if err := config.DB.Where("department = ?", queryData).Find(&issues).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, issues)
}

func FilterIssuesByMonitoringSystem(c *gin.Context) {
	queryData := helpers.GetQueryData(c, "monitoring_system")
	issues := []models.Issue{}

	if err := config.DB.Where("monitoring_system = ?", queryData).Find(&issues).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, issues)
}

func FilterIssuesByIssueCause(c *gin.Context) {
	queryData := helpers.GetQueryData(c, "issue_cause")
	issues := []models.Issue{}

	if err := config.DB.Where("issue_cause = ?", queryData).Find(&issues).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, issues)
}

func CreateIssue(c *gin.Context) {
	createIssueBody := models.CreateIssueBody{}

	if err := c.ShouldBindJSON(&createIssueBody); err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	newIssue := models.Issue{Date: createIssueBody.Date, Hour: createIssueBody.Hour, Unit: createIssueBody.Unit, Topic: createIssueBody.Topic, SpecificTopic: createIssueBody.SpecificTopic, MonitoringType: createIssueBody.MonitoringType, MonitoringSystem: createIssueBody.MonitoringSystem, IssueCause: createIssueBody.IssueCause, ResponsibleDepartment: createIssueBody.ResponsibleDepartment, Status: createIssueBody.Status}

	if err := config.DB.Create(&newIssue).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusCreated, newIssue)
}

func UpdateIssue(c *gin.Context) {
	id := helpers.GetParamData(c, "id")
	updateIssueBody := models.UpdatedIssueBody{}
	issue := findIssueById(id, c)

	if err := c.ShouldBindJSON(&updateIssueBody); err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	if err := config.DB.Model(&issue).Updates(models.Issue{Date: updateIssueBody.Date, Hour: updateIssueBody.Hour, Unit: updateIssueBody.Unit, Topic: updateIssueBody.Topic, SpecificTopic: updateIssueBody.SpecificTopic, MonitoringType: updateIssueBody.MonitoringType, MonitoringSystem: updateIssueBody.MonitoringSystem, IssueCause: updateIssueBody.IssueCause, ResponsibleDepartment: updateIssueBody.ResponsibleDepartment, Status: updateIssueBody.Status}).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, issue)
}

func DeleteIssue(c *gin.Context) {
	id := helpers.GetParamData(c, "id")
	deletedIssue := findIssueById(id, c)

	if err := config.DB.Delete(&deletedIssue).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, deletedIssue)
}

// Helpers

func findIssueById(id string, c *gin.Context) models.Issue {
	issue := models.Issue{}

	if err := config.DB.Where("issue_id = ?", id).First(&issue).Error; err != nil {
		helpers.DisplayErrorMsg(c, err)
	}

	return issue
}
