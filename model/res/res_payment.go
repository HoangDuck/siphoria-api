package res

import (
	"hotel-booking-api/model"
	"time"
)

type DataPaymentRes struct {
	Amount       int    `json:"amount"`
	Message      string `json:"message"`
	OrderID      string `json:"orderId"`
	PartnerCode  string `json:"partnerCode"`
	PayURL       string `json:"payUrl"`
	RequestID    string `json:"requestId"`
	ResponseTime int64  `json:"responseTime"`
	ResultCode   int    `json:"resultCode"`
}

type PaymentResponse struct {
	ID             string                  `json:"id"`
	PaymentMethod  string                  `json:"payment_method"`
	RankPrice      float32                 `json:"rank_price"`
	ConvertedPrice float32                 `json:"converted_price"`
	VoucherPrice   float32                 `json:"voucher_price"`
	TotalPrice     float32                 `json:"total_price"`
	StartAt        time.Time               `json:"start_at"`
	EndAt          time.Time               `json:"end_at"`
	TotalDay       int                     `json:"total_day"`
	UpdatedAt      time.Time               `json:"updated_at"`
	User           *model.User             `json:"user,omitempty"`
	RoomType       *model.RoomType         `json:"room_type,omitempty"`
	Hotel          *model.Hotel            `json:"hotel,omitempty"`
	RatePlan       *model.RatePlan         `json:"rate_plans,omitempty"`
	Details        []PaymentDetailResponse `json:"details" gorm:"-"`
}

type PaymentDetailResponse struct {
	ID          string    `json:"id"`
	DayOff      time.Time `json:"day_off"`
	Price       float32   `json:"price"`
	AdultNum    int       `json:"adult_num"`
	ChildrenNum int       `json:"children_num"`
}
