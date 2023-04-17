package req

type RequestAddBooking struct {
	CheckInTime      string  `json:"check_in_time"`
	CheckOutTime     string  `json:"check_out_time"`
	FirstName        string  `json:"first_name"`
	LastName         string  `json:"last_name"`
	Phone            string  `json:"phone"`
	Email            string  `json:"email"`
	Gender           bool    `json:"gender"`
	IdentifierNumber string  `json:"identifier_number"`
	RoomTypeCode     string  `json:"room_type_code"`
	CostRoomType     float32 `json:"cost_room_type"`
	CostRatePlan     float32 `json:"cost_rate_plan"`
	NumberRoom       int     `json:"number_room"`
	RatePlanID       string  `json:"rate_plan_id"`
	TypeBooking      string  `json:"type_booking"`
}

type RequestUpdateBooking struct {
}

type RequestCancelBooking struct {
	BookingID string `json:"booking_id"`
}

type RequestGetListBooking struct {
	CustomerName  string  `json:"customer_name"`
	RoomTypeCode  string  `json:"room_type_code"`
	StatusBooking string  `json:"status_booking"`
	TotalCost     float32 `json:"total_cost"`
	StatusPayment int     `json:"status_payment"`
	CheckInTime   string  `json:"check_in_time"`
	CheckOutTime  string  `json:"check_out_time"`
}

type RequestGetHistoryBooking struct {
	CustomerID string `json:"customer_id"`
}

type RequestGetBookingInfo struct {
	BookingID string `json:"booking_id"`
}

type RequestCheckInBooking struct {
	BookingID    string `json:"booking_id"`
	CheckInTime  string `json:"check_in_time"`
	CheckOutTime string `json:"check_out_time"`
	RoomID       string `json:"room_id"`
	RoomCode     string `json:"room_code"`
}
type RequestMultiCheckInBooking struct {
	ListCheckIn []RequestCheckInBooking `json:"list_check_in"`
}
