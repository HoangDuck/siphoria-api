package req

import "time"

type RequestGetRoomTypeInfo struct {
	TypeRoomCode string `json:"type_room_code"`
}

type RequestGetListRoomType struct {
}

type RequestGetRoomList struct {
}

type RequestGetRoomListFilterSearch struct {
	RoomTypeCode   string `json:"room_type_code"`
	NumberAdult    int    `json:"number_adult"`
	NumberChildren int    `json:"number_children"`
	NumberBed      int    `json:"bed_number"`
	NumberToilet   int    `json:"number_toilet"`
}

type RequestGetRoomListAvailable struct {
	TimeStart      string `json:"time_start"`
	TimeEnd        string `json:"time_end"`
	RoomTypeCode   string `json:"room_type_code"`
	NumberAdult    int    `json:"number_adult"`
	NumberChildren int    `json:"number_children"`
	NumberBed      int    `json:"bed_number"`
	NumberToilet   int    `json:"number_toilet"`
}

type RequestSearchGetRoomListAvailable struct {
	TimeStart      string `json:"time_start"`
	TimeEnd        string `json:"time_end"`
	NumberAdult    int    `json:"number_adult"`
	NumberChildren int    `json:"number_children"`
	NumberBed      int    `json:"bed_number"`
	NumberToilet   int    `json:"number_toilet"`
}

type RequestCheckRoomAvailable struct {
	TypeRoomCode string `json:"room_code"`
	TimeStart    string `json:"time_start"`
	TimeEnd      string `json:"time_end"`
}

type RequestUpdateRoomInfo struct {
	RoomCode     string `json:"room_code"`
	TypeRoomCode string `json:"type_room_code"`
}

type RequestGetRoomInfoByID struct {
	RoomCode string `json:"room_code"`
}

type RequestDeleteRoom struct {
	RoomCode string `json:"room_code"`
}

type RequestAddRoom struct {
	TypeRoomCode string `json:"type_room_code"`
	Floor        int    `json:"floor"`
}

type RequestAddRoomType struct {
	TypeRoomCode     string  `json:"type_room_code"`
	TypeRoomName     string  `json:"type_room_name"`
	Description      string  `json:"description"`
	ShortDescription string  `json:"short_description"`
	NumberAdult      int     `json:"number_adult"`
	NumberChildren   int     `json:"number_children"`
	NumberBed        int     `json:"bed_number"`
	NumberToilet     int     `json:"number_toilet"`
	CostType         float32 `json:"cost_type"`
	RoomImages       string  `json:"room_images"`
	Rating           int     `json:"rating"`
}

type RequestAddRoomBusyStatusCategory struct {
	StatusCode  string `json:"status_code"`
	StatusName  string `json:"status_name"`
	Description string `json:"description"`
}

type RequestSearchRoomListByCapacityAndTimeCheck struct {
	TimeStart      string `json:"time_start"`
	TimeEnd        string `json:"time_end"`
	NumberAdult    int    `json:"number_adult"`
	NumberChildren int    `json:"number_children"`
}

type RequestSearchRoomAvailable struct {
	TimeStart      string `json:"time_start"`
	TimeEnd        string `json:"time_end"`
	NumberAdult    int    `json:"number_adult"`
	NumberChildren int    `json:"number_children"`
	NumberRoom     int    `json:"number_room"`
}

type RequestSearchRoomAvailableAtReception struct {
	TimeStart    string `json:"time_start"`
	TimeEnd      string `json:"time_end"`
	RoomTypeCode string `json:"room_type_code"`
}

type RequestAddRoomBusyStatusDetail struct {
	RoomID    string `json:"room_id"`
	StatusID  string `json:"status_id"`
	BookingID string `json:"booking_id"`
	FromTime  string `json:"from_time"`
	ToTime    string `json:"to_time"`
}

type RequestUpdateRoomBusyStatusDetail struct {
	ID       string    `json:"-"`
	RoomID   string    `json:"room_id"`
	StatusID string    `json:"status_id"`
	FromTime time.Time `json:"from_time"`
	ToTime   time.Time `json:"to_time"`
}

type RequestCheckOutBooking struct {
	//RoomStatusDetailID string `json:"room_status_detail_id"`
	RoomID    string `json:"room_id"`
	BookingID string `json:"booking_id"`
	FromTime  string `json:"from_time"`
	ToTime    string `json:"to_time"`
}

type RequestGetRoomStatusInfo struct {
	BookingID string `json:"booking_id"`
}

type RequestGetRoomRatePlan struct {
	RoomTypeCode string `json:"room_type_code"`
}

type RequestGetRoomAvailable struct {
	RoomTypeID string `json:"room_type_id"`
}
