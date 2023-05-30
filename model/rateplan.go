package model

import "time"

type RatePlan struct {
	ID            string        `json:"id" gorm:"primary_key"`
	Name          string        `json:"name" gorm:"name"`
	Type          string        `json:"type" gorm:"type"`
	Status        int           `json:"status" gorm:"status"`
	Activate      bool          `json:"activate" gorm:"activate"`
	CreatedAt     time.Time     `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" gorm:"updated_at"`
	FreeBreakfast bool          `json:"free_breakfast" gorm:"free_breakfast"`
	FreeLunch     bool          `json:"free_lunch" gorm:"free_lunch"`
	FreeDinner    bool          `json:"free_dinner" gorm:"free_dinner"`
	RoomTypeId    string        `json:"room_type_id"`
	RoomType      *RoomType     `json:"room_type,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDeleted     bool          `json:"-" gorm:"is_deleted"`
	RatePackages  []RatePackage `json:"rate_packages,omitempty" gorm:"-"`
}

type RatePackage struct {
	ID             string    `json:"id" gorm:"primary_key"`
	Currency       string    `json:"currency" gorm:"currency"`
	AvailabilityAt time.Time `json:"availability_at" gorm:"availability_at"`
	Price          float32   `json:"price" gorm:"price"`
	CreatedAt      time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"updated_at"`
	RatePlanId     string    `json:"rate_plan_id"`
	RatePlan       *RatePlan `json:"rate_plan,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
