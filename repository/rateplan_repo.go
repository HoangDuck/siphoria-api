package repository

import "hotel-booking-api/model"

type RatePlanRepo interface {
	GetListRatePlan() ([]model.RatePlan, error)
	GetRatePlanInfo(ratePlan model.RatePlan) (model.RatePlan, error)
	UpdateRatePlanInfo(ratePlan model.RatePlan) (model.RatePlan, error)
	DeleteRatePlanInfo(condition map[string]interface{}) (bool, error)
	SaveRatePlan(ratePlan model.RatePlan) (bool, error)
	GetRoomTypeInfo(roomType model.RoomType) (model.RoomType, error)
	GetListRatePlanByRoomTypeCode(roomTypeID string) ([]model.RatePlan, error)
}
