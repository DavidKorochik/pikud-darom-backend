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
	issueBody := models.CreateIssueBody{}

	if err := c.ShouldBindJSON(&issueBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newIssue := models.Issue{Date: issueBody.Date, Hour: issueBody.Hour, Unit: issueBody.Unit, Topic: issueBody.Topic, SpecificTopic: issueBody.SpecificTopic, MonitoringType: issueBody.MonitoringType, MonitoringSystem: issueBody.MonitoringSystem, IssueCause: issueBody.IssueCause, ResponsibleDepartment: issueBody.ResponsibleDepartment, Status: issueBody.Status}
	config.DB.Create(&newIssue)

	c.JSON(http.StatusCreated, newIssue)
}

// Keep the updating from here!!!

func UpdateIssue(c *gin.Context) {
	id := getParamData(c, "id")
	issue := models.UpdatedIssueBody{}

	if err := c.BindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is an error with the issue body"})
		return
	}

	if err := config.DB.Model(&models.Issue{}).Where("issueId = ?", id).Updates(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't Update or Find the issue"})
		return
	}

	updatedIssue := findIssueById(id, c)

	c.JSON(http.StatusOK, updatedIssue)
}

func DeleteIssue(c *gin.Context) {
	id := getParamData(c, "id")

	if err := config.DB.Delete(&models.Issue{}, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't Delete or Find the issue"})
		return
	}

	deletedIssue := findIssueById(id, c)

	c.JSON(http.StatusOK, deletedIssue)
}

// Helpers

func findIssueById(id string, c *gin.Context) *models.Issue {
	issue := &models.Issue{}

	if err := config.DB.Find(&issue, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't find the issue"})
		return nil
	}

	return issue
}

func getParamData(c *gin.Context, param string) string {
	return c.Param(param)
}
