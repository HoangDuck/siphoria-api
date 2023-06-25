package model

import "time"

type Review struct {
	ID        string    `json:"id" gorm:"primary_key"`
	HotelId   string    `json:"hotel_id"`
	Hotel     *Hotel    `json:"hotel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId    string    `json:"user_id"`
	User      *User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rating    int       `json:"rating" gorm:"rating"`
	Content   string    `json:"content" gorm:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	IsDeleted bool      `json:"-" gorm:"is_deleted"`
}
