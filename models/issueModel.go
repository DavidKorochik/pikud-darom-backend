package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Issue struct {
	IssueID               uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Date                  string    `json:"date" gorm:"not null"`
	Hour                  string    `json:"hour" gorm:"not null"`
	Unit                  string    `json:"unit" gorm:"not null"`
	Topic                 string    `json:"topic" gorm:"not null"`
	SpecificTopic         string    `json:"specific_topic"`
	MonitoringType        string    `json:"monitoring_type" gorm:"not null"`
	MonitoringSystem      string    `json:"monitoring_system"`
	IssueCause            string    `json:"issue_cause" gorm:"not null"`
	ResponsibleDepartment string    `json:"responsible_department" gorm:"not null"`
	Status                string    `json:"status" gorm:"default:'handling'"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	DeletedAt             time.Time `json:"deleted_at"`
}

type UpdatedIssueBody struct {
	Date                  string `json:"date"`
	Hour                  string `json:"hour"`
	Unit                  string `json:"unit"`
	Topic                 string `json:"topic"`
	SpecificTopic         string `json:"specific_topic"`
	MonitoringType        string `json:"monitoring_type"`
	MonitoringSystem      string `json:"monitoring_system"`
	IssueCause            string `json:"issue_cause"`
	ResponsibleDepartment string `json:"responsible_department"`
	Status                string `json:"status"`
}

func (i *Issue) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("issue_id", uuid.New())
	return nil
}
