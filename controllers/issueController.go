package controllers

import (
	"net/http"

	"github.com/DavidKorochik/pikud-darom-backend/config"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllIssues(c *gin.Context) {
	issues := []models.Issue{}

	if err := config.DB.Find(&issues).Error; err != nil {
		displayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, issues)
}

func CreateIssue(c *gin.Context) {
	createIssueBody := models.CreateIssueBody{}

	if err := c.ShouldBindJSON(&createIssueBody); err != nil {
		displayErrorMsg(c, err)
		return
	}

	newIssue := models.Issue{Date: createIssueBody.Date, Hour: createIssueBody.Hour, Unit: createIssueBody.Unit, Topic: createIssueBody.Topic, SpecificTopic: createIssueBody.SpecificTopic, MonitoringType: createIssueBody.MonitoringType, MonitoringSystem: createIssueBody.MonitoringSystem, IssueCause: createIssueBody.IssueCause, ResponsibleDepartment: createIssueBody.ResponsibleDepartment, Status: createIssueBody.Status}

	if err := config.DB.Create(&newIssue).Error; err != nil {
		displayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusCreated, newIssue)
}

// Keep the updating from here!!!

func UpdateIssue(c *gin.Context) {
	id := getParamData(c, "id")
	updateIssueBody := models.UpdatedIssueBody{}
	issue := findIssueById(id, c)

	if err := c.ShouldBindJSON(&updateIssueBody); err != nil {
		displayErrorMsg(c, err)
		return
	}

	if err := config.DB.Model(&issue).Updates(models.Issue{Date: updateIssueBody.Date, Hour: updateIssueBody.Hour, Unit: updateIssueBody.Unit, Topic: updateIssueBody.Topic, SpecificTopic: updateIssueBody.SpecificTopic, MonitoringType: updateIssueBody.MonitoringType, MonitoringSystem: updateIssueBody.MonitoringSystem, IssueCause: updateIssueBody.IssueCause, ResponsibleDepartment: updateIssueBody.ResponsibleDepartment, Status: updateIssueBody.Status}).Error; err != nil {
		displayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, issue)
}

func DeleteIssue(c *gin.Context) {
	id := getParamData(c, "id")
	deletedIssue := findIssueById(id, c)

	if err := config.DB.Delete(&deletedIssue).Error; err != nil {
		displayErrorMsg(c, err)
		return
	}

	c.JSON(http.StatusOK, deletedIssue)
}

// Helpers

func findIssueById(id string, c *gin.Context) models.Issue {
	issue := models.Issue{}

	if err := config.DB.Where("issue_id = ?", id).First(&issue).Error; err != nil {
		displayErrorMsg(c, err)
	}

	return issue
}

func getParamData(c *gin.Context, param string) string {
	return c.Param(param)
}

func displayErrorMsg(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
