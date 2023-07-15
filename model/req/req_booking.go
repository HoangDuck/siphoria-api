package req

type RequestBookNow struct {
	FromDate         string `json:"from_date"`
	NumberOfAdults   int    `json:"number_of_adults" validate:"required,min=1"`
	NumberOfChildren int    `json:"number_of_children" validate:"required,min=0"`
	NumberOfRooms    int    `json:"number_of_rooms" validate:"required,min=1"`
	RatePlanID       string `json:"rate_plan_id"`
	RoomTypeID       string `json:"room_type_id"`
	HotelID          string `json:"hotel_id"`
	ToDate           string `json:"to_date"`
	UserId           string
}

type RequestCancelBooking struct {
	PaymentId string `json:"payment_id"`
}
