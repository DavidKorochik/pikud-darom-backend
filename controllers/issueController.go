package controllers

import (
	"net/http"

	"github.com/DavidKorochik/pikud-darom-backend/db"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAllIssues(c *gin.Context) {
	issues := []models.Issue{}
	db.DB.Find(&issues)

	c.JSON(http.StatusOK, gin.H{"issues": issues})
}

func CreateIssue(c *gin.Context) {
	issue := models.Issue{}

	if err := c.BindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	db.DB.Create(&issue)
	c.JSON(http.StatusCreated, gin.H{"issues": issue})
}

// Keep the updating from here!!!

func UpdateIssue(c *gin.Context) {
	id := getParamData(c, "id")
	issue := &models.UpdatedIssueBody{}

	if err := c.BindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There is an error with the issue body"})
	}

	if err := db.DB.Model(&models.Issue{}).Where("issueId = ?", id).Updates(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't Update or Find the issue"})
	}

	updatedIssue := findIssueById(id, c)

	c.JSON(http.StatusOK, gin.H{"issues": updatedIssue})
}

func DeleteIssue(c *gin.Context) {
	id := getParamData(c, "id")

	if err := db.DB.Delete(&models.Issue{}, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't Delete or Find the issue"})
	}

	deletedIssue := findIssueById(id, c)

	c.JSON(http.StatusOK, gin.H{"issues": deletedIssue})
}

func findIssueById(id string, c *gin.Context) *models.Issue {
	issue := &models.Issue{}

	if err := db.DB.Find(&issue, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't find the issue"})
	}

	return issue
}

func getParamData(c *gin.Context, param string) string {
	return c.Param(param)
}
