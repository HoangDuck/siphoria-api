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

type RequestCreateRoomType struct {
	ID             string
	Name           string `json:"name"`
	Description    string `json:"description"`
	BedNums        int    `json:"bed_nums"`
	BathroomNums   int    `json:"bathroom_nums"`
	Activated      bool   `json:"activated"`
	MaxChildren    int    `json:"max_children"`
	MaxAdult       int    `json:"max_adult"`
	HotelID        string `json:"hotel_id"`
	Photos         string `json:"photos"`
	Bay            bool   `json:"bay"`
	Ocean          bool   `json:"ocean"`
	Sea            bool   `json:"sea"`
	City           bool   `json:"city"`
	Garden         bool   `json:"garden"`
	Lake           bool   `json:"lake"`
	Mountain       bool   `json:"mountain"`
	River          bool   `json:"river"`
	PrivateBalcony bool   `json:"private_balcony"`
	AirConditional bool   `json:"air_conditional"`
	Tivi           bool   `json:"tivi"`
	Kitchen        bool   `json:"kitchen"`
	PrivatePool    bool   `json:"private_pool"`
	Iron           bool   `json:"iron"`
	Sofa           bool   `json:"sofa"`
	Desk           bool   `json:"desk"`
	Soundproof     bool   `json:"soundproof"`
	Towels         bool   `json:"towels"`
	Toiletries     bool   `json:"toiletries"`
	Shower         bool   `json:"shower"`
	Slippers       bool   `json:"slippers"`
	Hairdry        bool   `json:"hairdry"`
	Fruit          bool   `json:"fruit"`
	Bbq            bool   `json:"bbq"`
	Wine           bool   `json:"wine"`
	Fryer          bool   `json:"fryer"`
	KitchenTool    bool   `json:"kitchen_tool"`
}
