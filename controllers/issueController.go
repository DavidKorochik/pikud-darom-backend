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
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, issues)
}

func CreateIssue(c *gin.Context) {
	createIssueBody := models.CreateIssueBody{}

	if err := c.ShouldBindJSON(&createIssueBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newIssue := models.Issue{Date: createIssueBody.Date, Hour: createIssueBody.Hour, Unit: createIssueBody.Unit, Topic: createIssueBody.Topic, SpecificTopic: createIssueBody.SpecificTopic, MonitoringType: createIssueBody.MonitoringType, MonitoringSystem: createIssueBody.MonitoringSystem, IssueCause: createIssueBody.IssueCause, ResponsibleDepartment: createIssueBody.ResponsibleDepartment, Status: createIssueBody.Status}
	config.DB.Create(&newIssue)

	c.JSON(http.StatusCreated, newIssue)
}

// Keep the updating from here!!!

func UpdateIssue(c *gin.Context) {
	id := getParamData(c, "id")
	updateIssueBody := models.UpdatedIssueBody{}
	issue := findIssueById(id, c)

	if err := c.ShouldBindJSON(&updateIssueBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&issue).Updates(models.Issue{Date: updateIssueBody.Date, Hour: updateIssueBody.Hour, Unit: updateIssueBody.Unit, Topic: updateIssueBody.Topic, SpecificTopic: updateIssueBody.SpecificTopic, MonitoringType: updateIssueBody.MonitoringType, MonitoringSystem: updateIssueBody.MonitoringSystem, IssueCause: updateIssueBody.IssueCause, ResponsibleDepartment: updateIssueBody.ResponsibleDepartment, Status: updateIssueBody.Status}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, issue)
}

func DeleteIssue(c *gin.Context) {
	id := getParamData(c, "id")
	deletedIssue := findIssueById(id, c)

	if err := config.DB.Delete(&deletedIssue).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedIssue)
}

// Helpers

func findIssueById(id string, c *gin.Context) models.Issue {
	issue := models.Issue{}

	if err := config.DB.Where("issue_id = ?", id).First(&issue).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't find the issue"})
	}

	return issue
}

func getParamData(c *gin.Context, param string) string {
	return c.Param(param)
}
