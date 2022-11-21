package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/DavidKorochik/pikud-darom-backend/controllers"
	"github.com/DavidKorochik/pikud-darom-backend/db"
	"github.com/DavidKorochik/pikud-darom-backend/helpers"
	"github.com/DavidKorochik/pikud-darom-backend/initializers"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/google/uuid"
)

func init() {
	initializers.LoadEnvVariables("config.env")
	db.DBConnection()
}

func TestGetAllIssues(t *testing.T) {
	var issues []models.Issue

	a, w, r := helpers.CreateTestSuite(t)

	r.GET("/issues", controllers.GetAllIssues)

	req, err := http.NewRequest(http.MethodGet, "/issues", nil)

	if err != nil {
		a.Error(err)
		return
	}

	r.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &issues); err != nil {
		a.Error(err)
		return
	}

	a.Equal(http.StatusOK, w.Code)
	a.NotEmpty(issues)
}

func TestCreateIssue(t *testing.T) {
	mockIssue := models.Issue{
		IssueID:               uuid.New(),
		Date:                  "2003-12-05",
		Hour:                  "16:56",
		Unit:                  "89",
		Topic:                 "Just a topic",
		SpecificTopic:         "Spec topic",
		MonitoringType:        "MonType",
		MonitoringSystem:      "MonSystem",
		ResponsibleDepartment: "Not me that's for sure",
		Status:                "Completed",
		CreatedAt:             time.Now(),
		DeletedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}

	expected := models.Issue{}

	a, w, r := helpers.CreateTestSuite(t)

	r.POST("/issues", controllers.CreateIssue)

	jsonIssue, err := json.Marshal(mockIssue)

	if err != nil {
		a.Error(err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, "/issues", bytes.NewBuffer(jsonIssue))

	if err != nil {
		a.Error(err)
		return
	}

	if err := json.Unmarshal(w.Body.Bytes(), &expected); err != nil {
		a.Error(err)
		return
	}

	r.ServeHTTP(w, req)

	a.Equal(http.StatusCreated, w.Code)
	a.Equal(expected, mockIssue)
}

func TestUpdateIssue(t *testing.T) {
	jsonIssueStr := []byte(`{"unit": "Baah 35"}`)
	expected := models.Issue{}

	a, w, r := helpers.CreateTestSuite(t)

	r.PUT("/issues/1fceea4f-9e18-4a5f-b211-534c7ec817f9", controllers.UpdateIssue)

	req, err := http.NewRequest(http.MethodPut, "/issues/1fceea4f-9e18-4a5f-b211-534c7ec817f9", bytes.NewBuffer(jsonIssueStr))

	if err != nil {
		a.Error(err)
		return
	}

	r.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &expected); err != nil {
		a.Error(err)
		return
	}

	a.Equal(http.StatusOK, w.Code)
	a.Equal("Baah 35", expected.Unit)
	a.Equal(time.Now(), expected.UpdatedAt)
}

func TestDeleteIssue(t *testing.T) {
	deletedIssue := models.Issue{}

	a, w, r := helpers.CreateTestSuite(t)

	r.DELETE("/issues/1fceea4f-9e18-4a5f-b211-534c7ec817f9", controllers.DeleteIssue)

	req, err := http.NewRequest(http.MethodDelete, "/issues/1fceea4f-9e18-4a5f-b211-534c7ec817f9", nil)

	if err != nil {
		a.Error(err)
		return
	}

	if err := json.Unmarshal(w.Body.Bytes(), &deletedIssue); err != nil {
		a.Error(err)
		return
	}

	r.ServeHTTP(w, req)

	a.Equal(http.StatusOK, w.Code)
	a.Equal(time.Now(), deletedIssue.DeletedAt)
}
