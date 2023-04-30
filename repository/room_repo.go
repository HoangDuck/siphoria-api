package repository

import (
	"hotel-booking-api/model"
	"hotel-booking-api/model/req"
)

type RoomRepo interface {
	SaveRoomType(requestAddRoomType req.RequestCreateRoomType) (model.RoomType, error)
}
