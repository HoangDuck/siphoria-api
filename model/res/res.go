package res

import "hotel-booking-api/model"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type PaginationRes struct {
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Data  interface{} `json:"data,omitempty"`
}

type ResponseListRoomRatePlanByRoomTypeCode struct {
	TypeRoomCode string                `json:"room_type_code"`
	ListRoom     []ListRoomByFloorItem `json:"list_room"`
	ListRatePlan []model.RatePlan      `json:"list_rate_plan"`
}
