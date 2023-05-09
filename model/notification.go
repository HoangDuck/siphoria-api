package model

import "time"

type Notification struct {
	ID        string    `json:"id" gorm:"primary_key"`
	UserId    string    `json:"userId"`
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Token     string    `json:"token" gorm:"token"`
	Title     string    `json:"title" gorm:"title"`
	Body      string    `json:"body" gorm:"body"`
	Data      string    `json:"data" gorm:"data"`
	Seen      bool      `json:"seen" gorm:"seen"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	TimeSent  time.Time `json:"time_sent" gorm:"xtime_sent"`
}
