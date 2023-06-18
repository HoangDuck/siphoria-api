package model

import "time"

type Rank struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"name"`
	Price     float32   `json:"price" gorm:"price"`
	Discount  float32   `json:"discount" gorm:"discount"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

type UserRank struct {
	ID            string    `json:"id" gorm:"primary_key"`
	UserId        string    `json:"user_id"`
	User          *User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RankId        string    `json:"rank_id"`
	Rank          Rank      `json:"rank" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BeginAt       time.Time `json:"begin_at" gorm:"begin_at"`
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
	ExpiredAt     time.Time `json:"expired_at" gorm:"expired_at"`
	DurationYear  int       `json:"duration_year" gorm:"duration_year"`
	DurationMonth int       `json:"duration_month" gorm:"duration_month"`
	DurationDay   int       `json:"duration_day" gorm:"duration_day"`
}
