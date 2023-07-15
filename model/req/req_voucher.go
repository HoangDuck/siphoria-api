package req

type RequestAddVoucher struct {
	HotelID    string   `json:"hotel_id"`
	Name       string   `json:"name"`
	Discount   float32  `json:"discount"`
	Activated  bool     `json:"activated"`
	Code       string   `json:"code"`
	BeginAt    string   `json:"begin_at"`
	EndAt      string   `json:"end_at"`
	ExceptRoom []string `json:"except_room"`
}

type RequestUpdateVoucher struct {
	HotelId    string   `json:"hotel_id"`
	Name       string   `json:"name"`
	Discount   float32  `json:"discount"`
	Activated  bool     `json:"activated"`
	Code       string   `json:"code"`
	BeginAt    string   `json:"begin_at"`
	EndAt      string   `json:"end_at"`
	ExceptRoom []string `json:"except_room"`
}

type RequestApplyVoucher struct {
	SessionId string `json:"session_id"`
	Code      string `json:"code"`
}
