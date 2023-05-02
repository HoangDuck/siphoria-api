package model

import "time"

type Notification struct {
	ID        string    `json:"id"`
	UserId    string    `json:"userId"`
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Token     string    `json:"token"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Data      string    `json:"data"`
	Seen      bool      `json:"seen"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	TimeSent  time.Time `json:"time_sent" gorm:"xtime_sent"`
}
