package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	FirstName      string    `json:"first_name" gorm:"not null" binding:"required"`
	LastName       string    `json:"last_name" gorm:"not null" binding:"required"`
	ArmyEmail      string    `json:"army_email" gorm:"not null;unique" binding:"required,email"`
	PersonalNumber string    `json:"personal_number" gorm:"not null;unique" binding:"required,len=7"`
	Department     string    `json:"department" gorm:"not null" binding:"required"`
	Issues         []Issue   `json:"issues"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

type CreateUserBody struct {
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	ArmyEmail      string `json:"army_email" binding:"required,email"`
	PersonalNumber string `json:"personal_number" binding:"required,len=7"`
	Department     string `json:"department" binding:"required"`
}

type LogInUserBody struct {
	ArmyEmail      string `json:"army_email" binding:"required,email"`
	PersonalNumber string `json:"personal_number" binding:"required,len=7"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("issue_id", uuid.New())

	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	personalNumberBs := []byte(u.PersonalNumber)
	hashedPersonalNumber, err := bcrypt.GenerateFromPassword(personalNumberBs, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.PersonalNumber = string(hashedPersonalNumber)

	return nil
}
