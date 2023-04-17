package model

//
//import "time"
//
//type Booking struct {
//	ID                   string        `json:"ID" gorm:"primary_key"`
//	CustomerID           string        `json:"customer_id"`
//	Customer             Customer      `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//	CheckInTime          time.Time     `json:"check_in_time" gorm:"check_in_time"`
//	CheckOutTime         time.Time     `json:"check_out_time" gorm:"check_out_time"`
//	RoomTypeID           string        `json:"room_type_id"`
//	RoomType             RoomType      `json:"room_type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//	RatePlanID           string        `json:"rate_plan_id"`
//	RatePlan             RatePlan      `json:"rate_plan" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//	NumberRoom           int           `json:"number_room"`
//	CreatedAt            time.Time     `json:"created_at" gorm:"created_at"`
//	UpdatedAt            time.Time     `json:"updated_at" gorm:"updated_at"`
//	CanceledAt           time.Time     `json:"canceled_at" gorm:"canceled_at"`
//	TotalCost            float32       `json:"total_cost" gorm:"total_cost"`
//	Tax                  float32       `json:"tax" gorm:"tax"`
//	StatusBookingID      int           `json:"status_booking_id"`
//	StatusBooking        StatusBooking `json:"status_booking" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//	PaymentStatusID      int           `json:"payment_status" gorm:"payment_status"`
//	StaffIDHandleBooking string        `json:"staff_id_handle_booking" gorm:"staff_id_handle_booking"`
//}
//
//type StatusBooking struct {
//	ID          int    `json:"id" gorm:"primary_key;autoIncrement"`
//	StatusCode  string `json:"status_code" gorm:"status_code"`
//	StatusName  string `json:"status_name" gorm:"status_name"`
//	Description string `json:"description" gorm:"description"`
//}
//
//type BookingDetail struct {
//	ID        string    `json:"-" gorm:"primary_key"`
//	BookingID string    `json:"booking_id"`
//	Booking   Booking   `json:"booking" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//	RoomCode  string    `json:"room_code" gorm:"room_code"`
//	Floor     int       `json:"floor" gorm:"floor"`
//	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
//	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
//	Note      string    `json:"note" gorm:"note"`
//}
