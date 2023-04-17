package repository

//
//import (
//	"hotel-booking-api/model"
//	"hotel-booking-api/model/query"
//	"hotel-booking-api/model/req"
//)
//
//type RoomRepo interface {
//	GetRoomListByCondition(condition map[string]interface{}) ([]model.Room, error)
//	GetRoomListAvailable(condition map[string]interface{}) ([]model.Room, error)
//	GetRoomListType() ([]model.RoomType, error)
//	GetListHistoryBusyStatusRoom(room model.Room) ([]model.RoomBusyStatusDetail, error)
//	GetRoomBusyStatusType() ([]model.RoomBusyStatusCategory, error)
//	GetRoomTypeInfo(roomType model.RoomType) (model.RoomType, error)
//	UpdateRoomByRoomCode(room model.Room) (model.Room, error)
//	GetRoomInfo(request req.RequestGetRoomInfoByID) (model.Room, error)
//	DeleteRoomByID(request req.RequestDeleteRoom) (bool, error)
//	SaveRoom(room model.Room) (bool, error)
//	SaveRoomType(roomType model.RoomType) (bool, error)
//	CheckRoomAvailable(condition map[string]interface{}) ([]model.Room, error)
//	SaveRoomBusyStatusCategory(category model.RoomBusyStatusCategory) (bool, error)
//	CountAvailableRoomAndGroupByType(condition map[string]interface{}) ([]query.GroupNumberRoomByRoomType, error)
//	GetRoomListByCapacityAndTimeCheck(condition map[string]interface{}) ([]model.Room, error)
//	GetNumberRoomAvailableByCapacityAndTimeCheck(condition map[string]interface{}) ([]query.RoomRatePlanByRoomType, error)
//	GetNumberRoomAvailableByCapacityAndTimeCheckV2(condition map[string]interface{}) ([]query.RoomRatePlanCostByRoomType, error)
//	GetListRoomAtReception(condition map[string]interface{}) ([]query.RoomAvailableQuery, error)
//	SaveRoomBusyStatusDetail(roomBusy model.RoomBusyStatusDetail) (model.RoomBusyStatusDetail, error)
//	UpdateRoomBusyStatusDetail(roomBusy model.RoomBusyStatusDetail) (model.RoomBusyStatusDetail, error)
//	GetListRoomCheckOut(condition map[string]interface{}) ([]query.RoomStayedQuery, error)
//	GetStatusRoomCustomerInfo(condition map[string]interface{}) (query.RoomStatusInfoQuery, error)
//	GetRoomListByRoomTypeCode(roomTypeCode string) ([]model.Room, error)
//	GetListRatePlanByRoomTypeCode(roomTypeID string) ([]model.RatePlan, error)
//	GetListRoomCheckIn(roomTypeID string) ([]model.Room, error)
//	UpdateRoomByID(room model.Room) (model.Room, error)
//}
