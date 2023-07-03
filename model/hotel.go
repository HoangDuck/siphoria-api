package model

import "time"

type Hotel struct {
	ID              string        `json:"id" gorm:"primary_key"`
	Name            string        `json:"name" gorm:"name"`
	Overview        string        `json:"overview" gorm:"overview"`
	Rating          float32       `json:"rating" gorm:"rating"`
	CommissionRate  float32       `json:"commission_rate" gorm:"commission_rate"`
	CreatedAt       time.Time     `json:"created_at" gorm:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" gorm:"updated_at"`
	Status          int           `json:"status" gorm:"status"`
	Activate        bool          `json:"activate" gorm:"activate"`
	ProvinceCode    int           `json:"province_code" gorm:"province_code"`
	DistrictCode    int           `json:"district_code" gorm:"district_code"`
	WardCode        int           `json:"ward_code" gorm:"ward_code"`
	RawAddress      string        `json:"raw_address" gorm:"raw_address"`
	HotelPhotos     string        `json:"hotel_photos" gorm:"hotel_photos"`
	BankAccount     string        `json:"bank_account" gorm:"bank_account"`
	BankBeneficiary string        `json:"bank_beneficiary" gorm:"bank_beneficiary"`
	BankName        string        `json:"bank_name" gorm:"bank_name"`
	BusinessLicence string        `json:"business_licence" gorm:"business_licence"`
	HotelierId      string        `json:"hotelier_id,omitempty"`
	Hotelier        *User         `json:"hotelier,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDeleted       bool          `json:"-" gorm:"is_deleted"`
	PriceHotel      float32       `json:"price_hotel" gorm:"price_hotel"`
	DiscountPrice   float32       `json:"discount_price" gorm:"discount_price"`
	DiscountHotel   float32       `json:"discount_hotel" gorm:"discount_hotel"`
	HotelType       *HotelType    `json:"hotel_type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	HotelFacility   HotelFacility `json:"hotel_facility" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoomTypes       []RoomType    `json:"room_types" gorm:"-"`
}

type HotelWork struct {
	HotelId   string    `json:"hotel_id" gorm:"primary_key"`
	Hotel     Hotel     `json:"hotel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId    string    `json:"user_id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	IsDeleted bool      `json:"-" gorm:"is_deleted"`
}

type HotelType struct {
	HotelId   string    `json:"hotel_id" gorm:"primary_key"`
	Hotel     bool      `json:"hotel" gorm:"hotel"`
	Apartment bool      `json:"apartment" gorm:"apartment"`
	Resort    bool      `json:"resort" gorm:"resort"`
	Villa     bool      `json:"villa" gorm:"villa"`
	Camping   bool      `json:"camping" gorm:"camping"`
	Motel     bool      `json:"motel" gorm:"motel"`
	Homestay  bool      `json:"homestay" gorm:"homestay"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	IsDeleted bool      `json:"-" gorm:"is_deleted"`
}

type HotelFacility struct {
	HotelId       string    `json:"hotel_id" gorm:"primary_key"`
	Beach         bool      `json:"beach" gorm:"beach"`
	Pool          bool      `json:"pool" gorm:"pool"`
	Bar           bool      `json:"bar" gorm:"bar"`
	NoSmokingRoom bool      `json:"no_smoking_room" gorm:"no_smoking_room"`
	Fitness       bool      `json:"fitness" gorm:"fitness"`
	Spa           bool      `json:"spa" gorm:"spa"`
	Bath          bool      `json:"bath" gorm:"bath"`
	Wifi          bool      `json:"wifi" gorm:"wifi"`
	Breakfast     bool      `json:"breakfast" gorm:"breakfast"`
	Casino        bool      `json:"casino" gorm:"casino"`
	Parking       bool      `json:"parking" gorm:"parking"`
	CreatedAt     time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at"`
	IsDeleted     bool      `json:"-" gorm:"is_deleted"`
}

type HotelSearch struct {
	ID           string  `json:"id" gorm:"id"`
	Name         string  `json:"name" gorm:"name"`
	Overview     string  `json:"overview" gorm:"overview"`
	RatingCode   float32 `json:"rating_code" gorm:"rating_code"`
	ProvinceCode int     `json:"province_code" gorm:"province_code"`
	RawAddress   string  `json:"raw_address" gorm:"raw_address"`
	HotelPhotos  string  `json:"hotel_photos" gorm:"hotel_photos"`
	IsAvailable  string  `json:"is_available" gorm:"is_available"`
	AvgPrice     string  `json:"avg_price" gorm:"avg_price"`
}
