package query

import "time"

type DataQueryModel struct {
	Limit             int
	Page              int
	Filter            map[string]interface{}
	Search            string
	Sort              string
	Order             string
	Start             string
	End               string
	IsShowDeleted     bool
	Role              string
	UserId            string
	DataId            string
	ListIgnoreColumns []string
	TotalRows         int
	TotalPages        int
	PageViewIndex     int
}

type ResultTotalPrice struct {
	Sum float64 `gorm:"total_price"`
}

type GroupNumberRoomByRoomType struct {
	TypeRoomCode string `gorm:"type_room_code"`
	Count        int    `gorm:"count"`
}

type RoomRatePlanByRoomType struct {
	TypeRoomCode string  `gorm:"type_room_code"`
	Description  string  `gorm:"description"`
	Price        float32 `gorm:"price"`
	Count        int     `gorm:"count"`
}

type RoomRatePlanCostByRoomType struct {
	TypeRoomCode     string  `gorm:"type_room_code"`
	TypeRoomName     string  `gorm:"type_room_name"`
	Remains          int     `gorm:"remains"`
	Description      string  `gorm:"description"`
	ShortDescription string  `gorm:"short_description"`
	CostType         float32 `gorm:"cost_type"`
	RatePlanPrice    float32 `gorm:"rate_plan_price"`
	RoomImages       string  `gorm:"room_images"`
	NumberAdult      string  `gorm:"number_adult"`
	NumberChildren   string  `gorm:"number_children"`
	NumberBed        string  `gorm:"number_bed"`
	RatePlanId       string  `gorm:"rate_plan_id"`
}

type RoomAvailableQuery struct {
	ID             string  `gorm:"id"`
	RoomCode       string  `gorm:"room_code"`
	RoomTypeID     string  `gorm:"room_type_id"`
	TypeRoomCode   string  `gorm:"type_room_code"`
	TypeRoomName   string  `gorm:"type_room_name"`
	Description    string  `gorm:"description"`
	NumberAdult    int     `gorm:"number_adult"`
	NumberChildren int     `gorm:"number_children"`
	NumberBed      int     `gorm:"number_bed"`
	NumberToilet   int     `gorm:"number_toilet"`
	CostType       float32 `gorm:"cost_type"`
	RoomImages     string  `gorm:"room_images"`
}

type RoomStayedQuery struct {
	RoomCode       string `gorm:"room_code"`
	Floor          int    `gorm:"floor"`
	StatusDetailID string `gorm:"status_detail_id"`
	BookingID      string `gorm:"booking_id"`
}

type RoomStatusInfoQuery struct {
	FullName     string    `gorm:"full_name"`
	CheckInTime  time.Time `gorm:"check_in_time"`
	CheckOutTime time.Time `gorm:"check_out_time"`
	RoomCode     string    `gorm:"room_code"`
	Floor        int       `gorm:"floor"`
	BookingID    string    `gorm:"booking_id"`
}

type StatisticRevenueByTimeQuery struct {
	PaymentTime  time.Time `gorm:"payment_time"`
	RoomTypeCode string    `gorm:"room_type_code"`
	Sum          float32   `gorm:"sum"`
}

type StatisticRevenueByTypeRoomCode struct {
	RoomTypeCode string  `gorm:"room_type_code"`
	Sum          float32 `gorm:"sum"`
}
