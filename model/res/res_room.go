package res

import "time"

type RoomTypeGroupRatePlan struct {
	TypeRoomCode     string                `json:"type_room_code"`
	TypeRoomName     string                `json:"type_room_name"`
	Remains          int                   `json:"remains"`
	CostType         float32               `json:"cost_type"`
	NumberAdult      string                `json:"number_adult"`
	NumberChildren   string                `json:"number_children"`
	RoomImages       string                `json:"room_images"`
	NumberBed        string                `json:"number_bed"`
	ShortDescription string                `json:"short_description"`
	ListRatePlan     []RatePlanReduceModel `json:"list_rate_plan"`
}
type RatePlanReduceModel struct {
	RatePlanId    string  `json:"rate_plan_id"`
	RatePlanPrice float32 `json:"rate_plan_price"`
	Description   string  `json:"description"`
}

type StatisticRevenueByDay struct {
	Day                 string            `json:"day"`
	ListRoomTypeRevenue []RoomTypeRevenue `json:"list_room_type_revenue"`
}

type RoomTypeRevenue struct {
	RoomTypeCode string  `json:"room_type_code"`
	Sum          float32 `json:"sum"`
}

type ListRoomByFloorItem struct {
	Floor int `json:"floor"`
	//ListRoom []model.Room `json:"list_room"`
}

type RoomNightResponse struct {
	ID             string    `json:"id"`
	AvailabilityAt time.Time `json:"availability_at"`
	Quantity       int       `json:"quantity"`
}

type RatePlanResponse struct {
	RateplanID string  `json:"rateplan_id"`
	Prices     []Price `json:"prices"`
}

type Price struct {
	ID             string    `json:"id"`
	AvailabilityAt time.Time `json:"availability_at"`
	Price          float32   `json:"price"`
}
