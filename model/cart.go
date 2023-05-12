package model

import "time"

type AddCart struct {
	ID          string    `json:"id" gorm:"primary_key"`
	StartAt     time.Time `json:"start_at" gorm:"start_at"`
	EndAt       time.Time `json:"end_at" gorm:"end_at"`
	AdultNum    int       `json:"adult_num" gorm:"adult_num"`
	ChildrenNum int       `json:"children_num" gorm:"children_num"`
	RatePlanId  string    `json:"rate_plan_id"`
	UserId      string    `json:"user_id"`
	RoomTypeId  string    `json:"room_type_id"`
	HotelId     string    `json:"hotel_id"`
}

type Cart struct {
	ID             string    `json:"id" gorm:"primary_key"`
	IsSoldOut      bool      `json:"is_sold_out" gorm:"is_sold_out"`
	RankPrice      float32   `json:"rank_price" gorm:"rank_price"`
	ConvertedPrice float32   `json:"converted_price" gorm:"converted_price"`
	TotalPrice     float32   `json:"total_price" gorm:"total_price"`
	StartAt        time.Time `json:"start_at" gorm:"start_at"`
	EndAt          time.Time `json:"end_at" gorm:"end_at"`
	TotalDay       int       `json:"total_day" gorm:"total_day"`
	RatePlanId     string    `json:"rate_plan_id"`
	RatePlan       *RatePlan `json:"rate_plans" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId         string    `json:"user_id"`
	User           *User     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoomTypeId     string    `json:"room_type_id"`
	RoomType       *RoomType `json:"room_type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HotelId        string    `json:"hotel_id"`
	Hotel          *Hotel    `json:"hotel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt      time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"updated_at"`
}

type CartDetail struct {
	ID            string      `json:"id" gorm:"primary_key"`
	Remain        int         `json:"remain" gorm:"remain"`
	AdultNum      int         `json:"adult_num" gorm:"adult_num"`
	ChildrenNum   int         `json:"children_num" gorm:"children_num"`
	DayOff        time.Time   `json:"day_off" gorm:"day_off"`
	Price         float32     `json:"price" gorm:"price"`
	CartId        string      `json:"cart_id"`
	Cart          Cart        `json:"cart" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RatePackageId string      `json:"rate_package_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RatePackage   RatePackage `json:"rate_package"`
	RoomNightsId  string      `json:"room_nights_id"`
	RoomNights    RoomNights  `json:"room_nights" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time   `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at" gorm:"updated_at"`
}
