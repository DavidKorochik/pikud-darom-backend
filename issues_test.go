package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DavidKorochik/pikud-darom-backend/config"
	"github.com/DavidKorochik/pikud-darom-backend/controllers"
	"github.com/DavidKorochik/pikud-darom-backend/db"
	"github.com/DavidKorochik/pikud-darom-backend/initializers"
	"github.com/DavidKorochik/pikud-darom-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	initializers.LoadEnvVariables("config.env")
	db.DBConnection()
}

func TestGetAllIssues(t *testing.T) {
	actual := []models.Issue{}
	expected := []models.Issue{}

	req, w := sestGetBooksRouter()

	a := assert.New(t)

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)

	if err != nil {
		a.Error(err)
	}

	config.DB.Find(&expected)

	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	a.Equal(expected, actual)
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

	a := assert.New(t)

	reqBody, err := json.Marshal(mockIssue)

	if err != nil {
		a.Error(err)
	}

	req, w, err := setPostBooksRouter(bytes.NewBuffer(reqBody))

	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusCreated, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)

	if err != nil {
		a.Error(err)
	}

	actual := models.Issue{}

	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := mockIssue

	a.Equal(expected, actual)
}

// Helpers

func sestGetBooksRouter() (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()

	r.GET("/api/issues", controllers.GetAllIssues)

	req, err := http.NewRequest(http.MethodGet, "/api/issues", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return req, w
}

func setPostBooksRouter(body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()

	r.POST("/api/issues", controllers.CreateIssue)

	req, err := http.NewRequest(http.MethodPost, "/api/issues", body)

	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return req, w, nil
}
