package models

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"gorm.io/datatypes"
)

type Issue struct {
	ID                    uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;not null"`
	Date                  datatypes.Date `json:"date" gorm:"not null"`
	Hour                  datatypes.Time `json:"hour" gorm:"not null"`
	Unit                  string         `json:"unit" gorm:"not null"`
	Topic                 string         `json:"topic" gorm:"not null"`
	SpecificTopic         string         `json:"specificTopic"`
	MonitoringType        string         `json:"monitoringType" gorm:"not null"`
	MonitoringSystem      string         `json:"monitoringSystem"`
	IssueCause            string         `json:"issueCase" gorm:"not null"`
	ResponsibleDepartment string         `json:"responsibleDepartment" gorm:"not null"`
	Status                string         `json:"status" gorm:"default:'handling'"`
	CreatedAt             time.Time      `json:"createdAt"`
	UpdatedAt             time.Time      `json:"updatedAt"`
	DeletedAt             time.Time      `json:"deletedAt"`
}
