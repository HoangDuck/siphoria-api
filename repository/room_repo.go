package repository

import (
	"github.com/labstack/echo/v4"
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
	"hotel-booking-api/model/req"
)

type RoomRepo interface {
	SaveRoomType(requestAddRoomType req.RequestCreateRoomType) (model.RoomType, error)
	UpdateRoomNight(requestAddRoomType req.RequestUpdateRoomNight) ([]model.RoomNights, error)
	UpdateRatePackages(requestAddRatePackages req.RequestUpdateRatePackage) ([]model.RatePackage, error)
	UpdateRoomType(requestUpdateRoomType req.RequestUpdateRoomType, idRoomType string) (model.RoomType, error)
	UpdateRoomPhotos(room model.RoomType) (model.RoomType, error)
	GetRoomTypeDetail(room model.RoomType) (model.RoomType, error)
	GetListRoomTypeDetail(room model.RoomType) ([]model.RoomType, error)
	GetRoomTypeFacility(roomTypeId string) (model.RoomTypeFacility, error)
	GetRoomTypeViews(roomTypeId string) (model.RoomTypeViews, error)
	GetRoomNightsByRoomType(c echo.Context, roomTypeId string) ([]model.RoomNights, error)
	GetListRatePlans(c echo.Context, roomTypeId string) ([]model.RatePlan, error)
	GetListRatePackages(c echo.Context, ratePlanId string) ([]model.RatePackage, error)
	GetRatePlanByRoomTypeFilter(queryModel *query.DataQueryModel) ([]model.RatePlan, error)
	UpdateLockRoom()
	CountLockRoomByAvailabilityDay(roomTypeId string, roomNightsId string) (int, error)
}
