package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid;primary_key;not null;default:uuid_generate_v4()"`
	FirstName      string    `json:"first_name" gorm:"not null" binding:"required"`
	LastName       string    `json:"last_name" gorm:"not null" binding:"required"`
	PersonalNumber string    `json:"personal_number" gorm:"not null;unique" binding:"required,len=7"`
	Department     string    `json:"department" gorm:"not null" binding:"required"`
	Issues         []Issue   `json:"issues" gorm:"foreign_key:IssueId"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("issue_id", uuid.New())
	return nil
}
