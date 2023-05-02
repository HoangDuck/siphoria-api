package repository

import (
	"hotel-booking-api/model"
	"hotel-booking-api/model/req"
)

type RoomRepo interface {
	SaveRoomType(requestAddRoomType req.RequestCreateRoomType) (model.RoomType, error)
	UpdateRoomNight(requestAddRoomType req.RequestUpdateRoomNight) ([]model.RoomNights, error)
	UpdateRatePackages(requestAddRatePackages req.RequestUpdateRatePackage) ([]model.RatePackage, error)
	UpdateRoomType(requestUpdateRoomType req.RequestUpdateRoomType, idRoomType string) (model.RoomType, error)
}
