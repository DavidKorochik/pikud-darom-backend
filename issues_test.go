package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DavidKorochik/pikud-darom-backend/controllers"
	"github.com/DavidKorochik/pikud-darom-backend/db"
	"github.com/DavidKorochik/pikud-darom-backend/initializers"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	initializers.LoadEnvVariables("config.env")
	db.DBConnection()
}

func TestGetAllIssues(t *testing.T) {
	var issues []models.Issue

	a := assert.New(t)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

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

// func TestCreateIssue(t *testing.T) {
// 	mockIssue := models.Issue{
// 		IssueID:               uuid.New(),
// 		Date:                  "2003-12-05",
// 		Hour:                  "16:56",
// 		Unit:                  "89",
// 		Topic:                 "Just a topic",
// 		SpecificTopic:         "Spec topic",
// 		MonitoringType:        "MonType",
// 		MonitoringSystem:      "MonSystem",
// 		ResponsibleDepartment: "Not me that's for sure",
// 		Status:                "Completed",
// 		CreatedAt:             time.Now(),
// 		DeletedAt:             time.Now(),
// 		UpdatedAt:             time.Now(),
// 	}

// }
