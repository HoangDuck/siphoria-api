package req

import "encoding/json"

type RequestGetRatePlan struct {
	RatePlanID string `json:"rate_plan_id"`
}

type RequestUpdateRatePlan struct {
	Name          string `json:"name,omitempty"`
	Type          string `json:"type,omitempty"`
	Status        int    `json:"status,omitempty"`
	Activated     bool   `json:"activated,omitempty"`
	FreeBreakfast bool   `json:"free_breakfast,omitempty"`
	FreeLunch     bool   `json:"free_lunch,omitempty"`
	FreeDinner    bool   `json:"free_dinner,omitempty"`
	IsDelete      bool   `json:"is_delete,omitempty"`
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

//Request update list rate package

func UnmarshalRequestUpdateRatePackage(data []byte) (RequestUpdateRatePackage, error) {
	var requestModel RequestUpdateRatePackage
	err := json.Unmarshal(data, &requestModel)
	return requestModel, err
}

type RequestUpdateRatePackage struct {
	Data []RatePackageItem `json:"data"`
}

type RatePackageItem struct {
	Date     string  `json:"date"`
	RatePlan string  `json:"ratePlan"`
	Price    float32 `json:"price"`
}
