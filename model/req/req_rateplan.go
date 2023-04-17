package req

type RequestGetRatePlan struct {
	RatePlanID string `json:"rate_plan_id"`
}

type RequestUpdateRatePlan struct {
	RatePlanID  string  `json:"rate_plan_id"`
	RoomTypeID  string  `json:"room_type_id"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
type RequestDeleteRatePlan struct {
	RatePlanID string `json:"rate_plan_id"`
}

type RequestAddRatePlan struct {
	RoomTypeCode string  `json:"room_type_code"`
	Description  string  `json:"description"`
	Price        float32 `json:"price"`
}
