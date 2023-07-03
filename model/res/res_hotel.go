package res

type HotelDetailModel struct {
	Activated    bool       `json:"activated"`
	CityCode     string     `json:"city_code"`
	CountryCode  string     `json:"country_code"`
	DistrictCode string     `json:"district_code"`
	RawAddress   string     `json:"raw_address"`
	Facilities   Facilities `json:"facilities"`
	ID           string     `json:"id"`
	CreatedAt    int64      `json:"created_at"`
	Name         string     `json:"name"`
	Overview     string     `json:"overview"`
	HotelPhotos  []string   `json:"hotel_photos"`
	Rating       int        `json:"rating"`
	RoomTypes    []RoomType `json:"room_types"`
	UpdatedAt    int64      `json:"updated_at"`
}

type Facilities struct {
	Beach         bool `json:"beach"`
	Pool          bool `json:"pool"`
	Bar           bool `json:"bar"`
	NoSmokingRoom bool `json:"no_smoking_room"`
	Fitness       bool `json:"fitness"`
	Spa           bool `json:"spa"`
	Bath          bool `json:"bath"`
	Wifi          bool `json:"wifi"`
	Breakfast     bool `json:"breakfast"`
	Casio         bool `json:"casio"`
	Parking       bool `json:"parking"`
}

type RoomType struct {
	Description  string          `json:"description"`
	Facilities   map[string]bool `json:"facilities"`
	ID           string          `json:"id"`
	MaxAdult     int             `json:"max_adult"`
	MaxChildren  int             `json:"max_children"`
	BedNums      int             `json:"bed_nums"`
	BathroomNums int             `json:"bathroom_nums"`
	Name         string          `json:"name"`
	Photos       []string        `json:"photos"`
	RoomNights   []RoomNight     `json:"room_nights"`
	RatePlans    []RatePlan      `json:"rate_plans"`
	Views        Views           `json:"views"`
}

type RatePlan struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Type          int           `json:"type"`
	Status        int           `json:"status"`
	FreeBreakfast bool          `json:"free_breakfast"`
	FreeLunch     bool          `json:"free_lunch"`
	FreeDinner    bool          `json:"free_dinner"`
	RatePackages  []RatePackage `json:"rate_packages"`
}

type RatePackage struct {
	AvailableAt string `json:"available_at"`
	Currency    string `json:"currency"`
	Price       int    `json:"price"`
	UpdatedAt   int64  `json:"updated_at"`
}

type RoomNight struct {
	ID             string `json:"id"`
	Inventory      int    `json:"inventory"`
	Remain         int    `json:"remain"`
	AvailabilityAt string `json:"availability_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

type Views struct {
	Beach          bool `json:"beach"`
	City           bool `json:"city"`
	Lake           bool `json:"lake"`
	Mountain       bool `json:"mountain"`
	PrivateBalcony bool `json:"private_balcony"`
	Garden         bool `json:"garden,omitempty"`
	River          bool `json:"river,omitempty"`
	Bay            bool `json:"bay" gorm:"bay"`
	Sea            bool `json:"sea" gorm:"sea"`
	Ocean          bool `json:"ocean" gorm:"ocean"`
}
