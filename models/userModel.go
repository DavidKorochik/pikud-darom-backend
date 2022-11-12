package models

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type User struct {
	UserID         uuid.UUID `json:"userId" gorm:"type:uuid;primary_key;not null"`
	FirstName      string    `json:"firstName" gorm:"not null"`
	LastName       string    `json:"lastName" gorm:"not null"`
	PersonalNumber string    `json:"personalNumber" gorm:"not null;unique;size:7"`
	Department     string    `json:"department" gorm:"not null"`
	Issues         []Issue   `json:"issues" gorm:"foreign_key:IssueId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt"`
}

func (u *User) TableName() string {
	return "user"
}
