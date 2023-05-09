package req

type RequestAddVoucher struct {
	HotelID   string  `json:"hotel_id"`
	Name      string  `json:"name"`
	Discount  float32 `json:"discount"`
	Activated bool    `json:"activated"`
	Code      string  `json:"code"`
	BeginAt   string  `json:"begin_at"`
	EndAt     string  `json:"end_at"`
}

type RequestUpdateVoucher struct {
	Name      string  `json:"name"`
	Discount  float32 `json:"discount"`
	Activated bool    `json:"activated"`
	Code      string  `json:"code"`
	BeginAt   string  `json:"begin_at"`
	EndAt     string  `json:"end_at"`
}
