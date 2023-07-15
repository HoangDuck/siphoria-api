package model

import "time"

type Voucher struct {
	ID        string          `json:"id" gorm:"primary_key"`
	Name      string          `json:"name" gorm:"name"`
	Activated bool            `json:"activated" gorm:"activated"`
	BeginAt   time.Time       `json:"begin_at" gorm:"begin_at"`
	EndAt     time.Time       `json:"end_at" gorm:"end_at"`
	Code      string          `json:"code" gorm:"code"`
	Discount  float32         `json:"discount" gorm:"discount"`
	CreatedAt time.Time       `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"updated_at"`
	HotelId   string          `json:"hotel_id"`
	Hotel     *Hotel          `json:"hotel,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDeleted bool            `json:"-" gorm:"is_deleted"`
	Excepts   []VoucherExcept `json:"excepts" gorm:"-"`
}

type VoucherExcept struct {
	CreatedAt  time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
	VoucherId  string    `json:"voucher_id" gorm:"primary_key"`
	Voucher    *Voucher  `json:"voucher,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoomTypeId string    `json:"room_type_id" gorm:"primary_key"`
	RoomType   *RoomType `json:"room_type,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsDeleted  bool      `json:"-" gorm:"is_deleted"`
}
