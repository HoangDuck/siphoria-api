package model

import (
	"time"
)

type RoomType struct {
	ID               string           `json:"id" gorm:"primary_key"`
	Activated        bool             `json:"activated" gorm:"activated"`
	Name             string           `json:"name" gorm:"name"`
	Description      string           `json:"description" gorm:"description"`
	MaxAdult         int              `json:"max_adult" gorm:"max_adult"`
	MaxChildren      int              `json:"max_children" gorm:"max_children"`
	BedNums          int              `json:"bed_nums" gorm:"bed_nums"`
	BathroomNums     int              `json:"bathroom_nums" gorm:"bathroom_nums"`
	Photos           string           `json:"photos" gorm:"photos"`
	CreatedAt        time.Time        `json:"created_at" gorm:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at" gorm:"updated_at"`
	HotelId          string           `json:"hotel_id"`
	Hotel            *Hotel           `json:"hotel" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDeleted        bool             `json:"-" gorm:"is_deleted"`
	RoomTypeFacility RoomTypeFacility `json:"room_type_facility" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoomNights       []RoomNights     `json:"room_nights" gorm:"-"`
	RatePlans        []RatePlan       `json:"rate_plans" gorm:"-"`
	RoomTypeViews    RoomTypeViews    `json:"room_type_views" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RoomTypeViews struct {
	RoomTypeID     string    `json:"room_type_id" gorm:"primary_key"`
	Bay            bool      `json:"bay" gorm:"bay"`
	Sea            bool      `json:"sea" gorm:"sea"`
	Beach          bool      `json:"beach" gorm:"beach"`
	Ocean          bool      `json:"ocean" gorm:"ocean"`
	City           bool      `json:"city" gorm:"city"`
	Garden         bool      `json:"garden" gorm:"garden"`
	Lake           bool      `json:"lake" gorm:"lake"`
	Mountain       bool      `json:"mountain" gorm:"mountain"`
	River          bool      `json:"river" gorm:"river"`
	PrivateBalcony bool      `json:"private_balcony" gorm:"private_balcony"`
	CreatedAt      time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"updated_at"`
	IsDeleted      bool      `json:"-" gorm:"is_deleted"`
}

type RoomTypeFacility struct {
	RoomTypeID     string    `json:"room_type_id" gorm:"primary_key"`
	AirConditional bool      `json:"air_conditional" gorm:"air_conditional"`
	Tivi           bool      `json:"tivi" gorm:"tivi"`
	Kitchen        bool      `json:"kitchen" gorm:"kitchen"`
	PrivatePool    bool      `json:"private_pool" gorm:"private_pool"`
	Heater         bool      `json:"heater" gorm:"heater"`
	Iron           bool      `json:"iron" gorm:"iron"`
	Sofa           bool      `json:"sofa" gorm:"sofa"`
	Desk           bool      `json:"desk" gorm:"desk"`
	Soundproof     bool      `json:"soundproof" gorm:"soundproof"`
	Towels         bool      `json:"towels" gorm:"towels"`
	Toiletries     bool      `json:"toiletries" gorm:"toiletries"`
	Shower         bool      `json:"shower" gorm:"shower"`
	Slippers       bool      `json:"slippers" gorm:"slippers"`
	Hairdry        bool      `json:"hairdry" gorm:"hairdry"`
	Fruit          bool      `json:"fruit" gorm:"fruit"`
	Bbq            bool      `json:"bbq" gorm:"bbq"`
	Wine           bool      `json:"wine" gorm:"wine"`
	Fryer          bool      `json:"fryer" gorm:"fryer"`
	KitchenTool    bool      `json:"kitchen_tool" gorm:"kitchen_tool"`
	CreatedAt      time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"updated_at"`
	IsDeleted      bool      `json:"-" gorm:"is_deleted"`
}

type RoomNights struct {
	ID             string    `json:"id" gorm:"primary_key"`
	AvailabilityAt time.Time `json:"availability_at" gorm:"availability_at"`
	Inventory      int       `json:"inventory" gorm:"inventory"`
	Remain         int       `json:"remain" gorm:"remain"`
	CreatedAt      time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt      time.Time `json:"-" gorm:"updated_at"`
	RoomTypeId     string    `json:"room_type_id"`
	RoomType       *RoomType `json:"room_type,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type LockRoom struct {
	ID           string      `json:"id" gorm:"primary_key"`
	RoomNightsId string      `json:"room_nights_id"`
	RoomNights   *RoomNights `json:"room_nights,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoomTypeID   string      `json:"room_type_id"`
	RoomType     *RoomType   `json:"room_type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId       string      `json:"user_id"`
	User         *User       `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LockFrom     time.Time   `json:"lock_from" gorm:"lock_from"`
	LockTo       time.Time   `json:"lock_to" gorm:"lock_to"`
	Expired      bool        `json:"expired" gorm:"expired"`
}
