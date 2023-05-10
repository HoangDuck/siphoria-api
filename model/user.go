package model

import (
	"time"
)

type User struct {
	ID              string    `json:"id" gorm:"primary_key"`
	Avatar          string    `json:"avatar" gorm:"avatar"`
	Email           string    `json:"email" gorm:"email"`
	FirstName       string    `json:"first_name" gorm:"first_name"`
	LastName        string    `json:"last_name" gorm:"last_name"`
	FullName        string    `json:"full_name" gorm:"full_name"`
	Phone           string    `json:"phone" gorm:"phone"`
	Gender          bool      `json:"gender" gorm:"gender"`
	Role            string    `json:"role" gorm:"role"`
	Status          int       `json:"status" gorm:"status"`
	Password        string    `json:"-" gorm:"password"`
	Token           *Token    `json:"token,omitempty" gorm:"-"`
	UserKeyFirebase string    `json:"user_key_firebase" gorm:"user_key_firebase"`
	CreatedAt       time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt       time.Time `json:"-" gorm:"deleted_at"`
	IsDeleted       bool      `json:"-" gorm:"is_deleted"`
	UserRank        *UserRank `json:"user_rank"`
}

type StatusUser struct {
	ID          int    `json:"id" gorm:"primary_key;autoIncrement"`
	StatusCode  string `json:"status_code" gorm:"status_code"`
	StatusName  string `json:"status_name" gorm:"status_name"`
	Description string `json:"description" gorm:"description"`
}

type Token struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiredTime  time.Duration `json:"expired_time"`
}
