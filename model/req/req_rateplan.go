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
	RoomTypeID    string `json:"room_type_id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Status        int    `json:"status"`
	Activated     bool   `json:"activated"`
	FreeBreakfast bool   `json:"free_breakfast"`
	FreeLunch     bool   `json:"free_lunch"`
	FreeDinner    bool   `json:"free_dinner"`
}
