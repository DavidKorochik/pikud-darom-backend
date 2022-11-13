package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Issue struct {
	IssueID               uuid.UUID `json:"issue_id" gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Date                  string    `json:"date" gorm:"not null" binding:"required" time_format:"2006-01-02"`
	Hour                  string    `json:"hour" gorm:"not null" binding:"required" time_format:"15:04:05"`
	Unit                  string    `json:"unit" gorm:"not null" binding:"required"`
	Topic                 string    `json:"topic" gorm:"not null" binding:"required"`
	SpecificTopic         string    `json:"specific_topic" gorm:"not null" binding:"required"`
	MonitoringType        string    `json:"monitoring_type" gorm:"not null" binding:"required"`
	MonitoringSystem      string    `json:"monitoring_system" binding:"required"`
	IssueCause            string    `json:"issue_cause" gorm:"not null" binding:"required"`
	ResponsibleDepartment string    `json:"responsible_department" gorm:"not null" binding:"required"`
	Status                string    `json:"status" gorm:"default:'handling'"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	DeletedAt             time.Time `json:"deleted_at"`
}

type CreateIssueBody struct {
	Date                  string `json:"date" binding:"required" time_format:"2006-01-02"`
	Hour                  string `json:"hour" binding:"required" time_format:"15:04:05"`
	Unit                  string `json:"unit" binding:"required"`
	Topic                 string `json:"topic" binding:"required"`
	SpecificTopic         string `json:"specific_topic" binding:"required"`
	MonitoringType        string `json:"monitoring_type" binding:"required"`
	MonitoringSystem      string `json:"monitoring_system" binding:"required"`
	IssueCause            string `json:"issue_cause" binding:"required"`
	ResponsibleDepartment string `json:"responsible_department" binding:"required"`
	Status                string `json:"status"`
}

type UpdatedIssueBody struct {
	Date                  string `json:"date" time_format:"2006-01-02"`
	Hour                  string `json:"hour" time_format:"15:04:05"`
	Unit                  string `json:"unit"`
	Topic                 string `json:"topic"`
	SpecificTopic         string `json:"specific_topic"`
	MonitoringType        string `json:"monitoring_type"`
	MonitoringSystem      string `json:"monitoring_system"`
	IssueCause            string `json:"issue_cause"`
	ResponsibleDepartment string `json:"responsible_department"`
	Status                string `json:"status"`
}

func (i *Issue) TableName() string {
	return "issue"
}

func (i *Issue) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("issue_id", uuid.New())
	return nil
}
