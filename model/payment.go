package model

import "time"

type PaymentMethod struct {
	ID          int    `json:"id" gorm:"primary_key;autoIncrement"`
	MethodName  string `json:"method_name" gorm:"method_name"`
	Description string `json:"description" gorm:"description"`
	Provider    string `json:"provider" gorm:"provider"`
}

type PaymentStatus struct {
	ID          int    `json:"id" gorm:"primary_key;autoIncrement"`
	StatusCode  string `json:"status_code" gorm:"status_code"`
	StatusName  string `json:"status_name" gorm:"status_name"`
	Description string `json:"description" gorm:"description"`
}

type Payment struct {
	ID              string          `json:"id" gorm:"primary_key"`
	Price           float32         `json:"price" gorm:"price"`
	DayOff          time.Time       `json:"day_off" gorm:"day_off"`
	PaymentMethod   string          `json:"payment_method" gorm:"payment_method"`
	AdultNum        int             `json:"adult_num" gorm:"adult_num"`
	ChildrenNum     int             `json:"children_num" gorm:"children_num"`
	ConvertedPrice  float32         `json:"converted_price" gorm:"converted_price"`
	RankPrice       float32         `json:"rank_price" gorm:"rank_price"`
	VoucherPrice    float32         `json:"voucher_price" gorm:"voucher_price"`
	TotalPrice      float32         `json:"total_price" gorm:"total_price"`
	StartAt         time.Time       `json:"start_at" gorm:"start_at"`
	EndAt           time.Time       `json:"end_at" gorm:"end_at"`
	TotalDay        int             `json:"total_day" gorm:"total_day"`
	Status          string          `json:"status" gorm:"status"`
	UserId          string          `json:"user_id"`
	User            User            `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoomTypeId      string          `json:"room_type_id"`
	RoomType        *RoomType       `json:"room_type,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VoucherId       string          `json:"voucher_id"`
	Voucher         *Voucher        `json:"voucher,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PayoutRequestId string          `json:"payout_request_id"`
	PayoutRequest   *PayoutRequest  `json:"payout_request,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HotelId         string          `json:"hotel_id"`
	Hotel           *Hotel          `json:"hotel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SessionId       string          `json:"session_id" gorm:"session_id"`
	RatePlanId      string          `json:"rate_plan_id"`
	RatePlan        *RatePlan       `json:"rate_plan" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt       time.Time       `json:"created_at" gorm:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"updated_at"`
	CartId          string          `json:"cart_id" gorm:"cart_id"`
	PaymentDetail   []PaymentDetail `json:"payment_details" gorm:"-"`
}

type PaymentDetail struct {
	ID            string       `json:"id" gorm:"primary_key"`
	Price         float32      `json:"price" gorm:"price"`
	AdultNum      int          `json:"adult_num" gorm:"adult_num"`
	ChildrenNum   int          `json:"children_num" gorm:"children_num"`
	DayOff        time.Time    `json:"day_off" gorm:"day_off"`
	PaymentId     string       `json:"payment_id"`
	Payment       *Payment     `json:"payment,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RatePlanId    string       `json:"rate_plan_id"`
	RatePlan      *RatePlan    `json:"rate_plan,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId        string       `json:"user_id"`
	User          *User        `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RatePackageId string       `json:"rate_package_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RatePackage   *RatePackage `json:"rate_package,omitempty"`
}

type PayoutRequest struct {
	ID           string    `json:"id" gorm:"primary_key"`
	OpenAt       time.Time `json:"open_at" gorm:"open_at"`
	CloseAt      time.Time `json:"close_at" gorm:"close_at"`
	Resolve      bool      `json:"resolve" gorm:"resolve"`
	TotalRequest int       `json:"total_request" gorm:"total_request"`
	TotalPrice   float32   `json:"total_price" gorm:"total_price"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"updated_at"`
	HotelId      string    `json:"hotel_id"`
	Hotel        Hotel     `json:"hotel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PettionerId  string    `json:"pettioner_id"`
	Pettioner    User      `json:"pettioner" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PayerId      string    `json:"payer_id"`
	Payer        User      `json:"payer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentList  string    `json:"payment_list" gorm:"payment_list"`
}
