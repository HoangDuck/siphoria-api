package repo_impl

//
//import (
//	"fmt"
//	"go.uber.org/zap"
//	"gorm.io/gorm"
//	"hotel-booking-api/custom_error"
//	"hotel-booking-api/db"
//	"hotel-booking-api/logger"
//	"hotel-booking-api/model"
//	"hotel-booking-api/model/query"
//	"hotel-booking-api/model/req"
//	"hotel-booking-api/repository"
//)
//
//type RoomRepoImpl struct {
//	sql *db.Sql
//}
//
//func NewRoomRepo(sql *db.Sql) repository.RoomRepo {
//	return &RoomRepoImpl{
//		sql: sql,
//	}
//}
//
//func (roomReceiver *RoomRepoImpl) GetListRatePlanByRoomTypeCode(roomTypeID string) ([]model.RatePlan, error) {
//	var listRatePlan []model.RatePlan
//	err := roomReceiver.sql.Db.Where("room_type_id=?", roomTypeID).Preload("RoomType").Find(&listRatePlan)
//	if err != nil {
//		return listRatePlan, err.Error
//	}
//	return listRatePlan, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetRoomListByRoomTypeCode(roomTypeCode string) ([]model.Room, error) {
//	var listRoom []model.Room
//	err := roomReceiver.sql.Db.Where("room_type_id=?", roomTypeCode).Preload("RoomType").Find(&listRoom)
//	if err != nil {
//		return listRoom, err.Error
//	}
//	return listRoom, nil
//}
//
//func (roomReceiver *RoomRepoImpl) UpdateRoomBusyStatusDetail(roomBusy model.RoomBusyStatusDetail) (model.RoomBusyStatusDetail, error) {
//	var roomResult = model.RoomBusyStatusDetail{}
//	err := roomReceiver.sql.Db.Model(&roomResult).Where("id=?", roomResult.ID).Preload("RoomBusyStatusCategory").Preload("Room").Preload("Booking").Updates(roomBusy)
//	logger.Debug("Get data", zap.Error(err.Error))
//	if err.Error == gorm.ErrRecordNotFound {
//		return roomResult, err.Error
//	}
//	return roomResult, nil
//}
//
//func (roomReceiver *RoomRepoImpl) SaveRoomBusyStatusDetail(roomBusy model.RoomBusyStatusDetail) (model.RoomBusyStatusDetail, error) {
//	result := roomReceiver.sql.Db.Create(&roomBusy)
//	if result.Error != nil {
//		return roomBusy, result.Error
//	}
//	return roomBusy, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetListRoomAtReception(condition map[string]interface{}) ([]query.RoomAvailableQuery, error) {
//	var listRoomAvailable []query.RoomAvailableQuery
//	queryString := fmt.Sprintf("select "+
//		"tb1.id,tb1.room_code,tb1.room_type_id,tb2.type_room_code,tb2.type_room_name, "+
//		"tb2.description,tb2.number_adult,tb2.number_children,tb2.number_bed, "+
//		"tb2.number_toilet,tb2.cost_type,tb2.room_images "+
//		"from rooms as tb1 LEFT JOIN room_types as tb2 on tb1.room_type_id=tb2.id "+
//		"where lower(tb2.type_room_code) LIKE '%%%s%%' AND tb1.id not in ("+
//		"select id from room_busy_status_details "+
//		"where '%s'::date <= to_time AND '%s'::date >= from_time "+
//		"AND status_id in (select id from room_busy_status_categories where status_code='stay')) ",
//		condition["room_type_code"],
//		condition["start_time"],
//		condition["end_time"])
//	err := roomReceiver.sql.Db.Raw(queryString).Find(&listRoomAvailable)
//	if err.RowsAffected == 0 {
//		return listRoomAvailable, nil
//	}
//	return listRoomAvailable, nil
//}
//
//func (roomReceiver *RoomRepoImpl) SaveRoomBusyStatusCategory(category model.RoomBusyStatusCategory) (bool, error) {
//	result := roomReceiver.sql.Db.Create(&category)
//	if result.Error != nil {
//		return false, result.Error
//	}
//	return true, nil
//}
//
//func (roomReceiver *RoomRepoImpl) SaveRoomType(roomType model.RoomType) (bool, error) {
//	result := roomReceiver.sql.Db.Create(&roomType)
//	if result.Error != nil {
//		return false, result.Error
//	}
//	return true, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetRoomListByCondition(condition map[string]interface{}) ([]model.Room, error) {
//	var listRoom []model.Room
//	if condition["isGetAll"] == "true" {
//		err := roomReceiver.sql.Db.Preload("RoomType").Find(&listRoom)
//		if err != nil {
//			return listRoom, err.Error
//		}
//	} else {
//		queryString := fmt.Sprintf("Select * from rooms "+
//			"where room_type_id in (select ID from room_types where lower(type_room_code) LIKE '%%%s%%') "+ //
//			"AND room_type_id in ("+
//			"select ID from room_types where number_adult>=%d "+
//			"AND number_children>=%d "+
//			"AND number_bed>=%d "+
//			"AND number_toilet>=%d "+
//			")", condition["room_type_code"],
//			condition["number_adult"],
//			condition["number_children"],
//			condition["bed_number"],
//			condition["number_toilet"])
//		if condition["room_type_code"] == "all" {
//			queryString = fmt.Sprintf("Select * from rooms "+
//				"where room_type_id in ("+
//				"select ID from room_types where number_adult>=%d "+
//				"AND number_children>=%d "+
//				"AND number_bed>=%d "+
//				"AND number_toilet>=%d "+
//				")",
//				condition["number_adult"],
//				condition["number_children"],
//				condition["bed_number"],
//				condition["number_toilet"])
//		}
//		err := roomReceiver.sql.Db.Raw(queryString).Preload("RoomType").Find(&listRoom)
//		if err != nil {
//			logger.Error("Get get data", zap.Error(err.Error))
//			return listRoom, err.Error
//		}
//	}
//	return listRoom, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetRoomListAvailable(condition map[string]interface{}) ([]model.Room, error) {
//	var listRoom []model.Room
//	queryString := fmt.Sprintf("Select * from rooms "+
//		"where ID in (%s) AND ID not in ("+
//		"Select room_id from room_busy_status_details "+
//		"where '%s'::date >= from_time AND '%s'::date <= to_time "+
//		"AND status_id in (Select ID from room_busy_status_categories where status_name='Booked')"+
//		")",
//		condition["room_ids"],
//		condition["time_start"],
//		condition["time_end"])
//	err := roomReceiver.sql.Db.Raw(queryString).Preload("RoomType").Find(&listRoom)
//	if err != nil {
//		return listRoom, err.Error
//	}
//	return listRoom, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetRoomListType() ([]model.RoomType, error) {
//	var listRoomType []model.RoomType
//	err := roomReceiver.sql.Db.Find(&listRoomType)
//	if err != nil {
//		return listRoomType, err.Error
//	}
//	return listRoomType, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetRoomBusyStatusType() ([]model.RoomBusyStatusCategory, error) {
//	var listRoomBusyStatusType []model.RoomBusyStatusCategory
//	err := roomReceiver.sql.Db.Find(&listRoomBusyStatusType)
//	if err != nil {
//		return listRoomBusyStatusType, err.Error
//	}
//	return listRoomBusyStatusType, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetRoomInfo(request req.RequestGetRoomInfoByID) (model.Room, error) {
//	var roomInfo = model.Room{}
//	err := roomReceiver.sql.Db.Where("room_code = ?", request.RoomCode).Preload("RoomType").Find(&roomInfo)
//	if err.RowsAffected == 0 {
//		return roomInfo, custom_error.RoomConflict
//	}
//	return roomInfo, nil
//}
//
//func (roomReceiver *RoomRepoImpl) DeleteRoomByID(request req.RequestDeleteRoom) (bool, error) {
//	err := roomReceiver.sql.Db.Where("room_code = ?", request.RoomCode).Delete(&model.Room{})
//	if err.RowsAffected == 0 {
//		return false, custom_error.RoomDeleteFailed
//	}
//	return true, nil
//}
//
//func (roomReceiver *RoomRepoImpl) SaveRoom(room model.Room) (bool, error) {
//	result := roomReceiver.sql.Db.Create(&room)
//	if result.Error != nil {
//		return false, result.Error
//	}
//	return true, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetRoomTypeInfo(roomType model.RoomType) (model.RoomType, error) {
//	var roomTypeResult = model.RoomType{}
//	err := roomReceiver.sql.Db.Where("type_room_code =?", roomType.Name).Find(&roomTypeResult)
//	if err != nil {
//		return roomTypeResult, err.Error
//	}
//	return roomTypeResult, nil
//}
//
//func (roomReceiver *RoomRepoImpl) UpdateRoomByRoomCode(room model.Room) (model.Room, error) {
//	var roomResult = model.Room{}
//	err := roomReceiver.sql.Db.Model(&roomResult).Where("room_code=?", room.RoomCode).Preload("RoomType").Updates(room)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return roomResult, err.Error
//		}
//
//		return roomResult, custom_error.RoomNotUpdated
//	}
//	return roomResult, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetListHistoryBusyStatusRoom(room model.Room) ([]model.RoomBusyStatusDetail, error) {
//	var listHistoryStatus []model.RoomBusyStatusDetail
//	err := roomReceiver.sql.Db.Where("room_id=?", room.ID).Preload("Room").Preload("RoomBusyStatusCategory").Find(&listHistoryStatus)
//	if err != nil {
//		return listHistoryStatus, custom_error.RoomNotFound
//	}
//	return listHistoryStatus, nil
//}
//
//func (roomReceiver *RoomRepoImpl) CheckRoomAvailable(condition map[string]interface{}) ([]model.Room, error) {
//	var listRoomAvailable []model.Room
//	queryString := fmt.Sprintf("Select * from rooms "+
//		"where room_type_id in (Select id from room_types where lower(type_room_code) LIKE '%%%s%%' "+ //
//		"AND ID in (Select room_id from room_busy_status_details where '%s'::date >= from_time::date AND '%s'::date <= to_time)",
//		condition["type_room_code"],
//		condition["start_time"],
//		condition["end_time"])
//	err := roomReceiver.sql.Db.Raw(queryString).Preload("RoomType").Find(&listRoomAvailable)
//	if err.RowsAffected == 0 {
//		return listRoomAvailable, nil
//	}
//	if err != nil {
//		logger.Error("Error query data", zap.Error(err.Error))
//		return listRoomAvailable, err.Error
//	}
//	return listRoomAvailable, nil
//}
//
//func (roomReceiver *RoomRepoImpl) CountAvailableRoomAndGroupByType(condition map[string]interface{}) ([]query.GroupNumberRoomByRoomType, error) {
//	var listRoomAvailable []query.GroupNumberRoomByRoomType
//	queryString := fmt.Sprintf("Select room_types.type_room_code,count(rooms.id) from rooms LEFT JOIN room_types on rooms.room_type_id = room_types.id "+
//		"where rooms.ID in (%s) AND rooms.ID not in ("+
//		"Select room_id from room_busy_status_details "+
//		"where '%s'::date >= from_time AND '%s'::date <= to_time "+
//		"AND status_id in (Select ID from room_busy_status_categories where status_name='Booked')"+
//		")"+
//		"GROUP BY room_types.type_room_code",
//		condition["room_ids"],
//		condition["start_time"],
//		condition["end_time"])
//	err := roomReceiver.sql.Db.Raw(queryString).Find(&listRoomAvailable)
//	if err.RowsAffected == 0 {
//		return listRoomAvailable, nil
//	}
//	return listRoomAvailable, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetRoomListByCapacityAndTimeCheck(condition map[string]interface{}) ([]model.Room, error) {
//	var listRoom []model.Room
//	if condition["isGetAll"] == "true" {
//		err := roomReceiver.sql.Db.Preload("RoomType").Find(&listRoom)
//		if err != nil {
//			return listRoom, err.Error
//		}
//	} else {
//		queryString := fmt.Sprintf("Select * from rooms "+
//			"where room_type_id in (select ID from room_types where lower(type_room_code) LIKE '%%%s%%') "+ //
//			"AND room_type_id in ("+
//			"select ID from room_types where number_adult>=%d "+
//			"AND number_children>=%d "+
//			")", condition["room_type_code"],
//			condition["number_adult"],
//			condition["number_children"])
//		if condition["room_type_code"] == "all" {
//			queryString = fmt.Sprintf("Select * from rooms "+
//				"where room_type_id in ("+
//				"select ID from room_types where number_adult>=%d "+
//				"AND number_children>=%d "+
//				")",
//				condition["number_adult"],
//				condition["number_children"])
//		}
//		err := roomReceiver.sql.Db.Raw(queryString).Preload("RoomType").Find(&listRoom)
//		if err != nil {
//			logger.Error("Get get data", zap.Error(err.Error))
//			return listRoom, err.Error
//		}
//	}
//	return listRoom, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetNumberRoomAvailableByCapacityAndTimeCheck(condition map[string]interface{}) ([]query.RoomRatePlanByRoomType, error) {
//	var listRoomAvailable []query.RoomRatePlanByRoomType
//	queryString := fmt.Sprintf("Select tb3.type_room_code,tb4.description,tb4.price,tb3.count "+
//		"from (Select tb2.id,tb1.type_room_code,tb1.count from "+
//		"(Select room_types.type_room_code,count(rooms.id) "+
//		"from rooms LEFT JOIN room_types on rooms.room_type_id = room_types.id "+
//		"where rooms.ID in (%s) AND rooms.ID not in ("+
//		"Select room_id from room_busy_status_details "+
//		"where ('%s'::date BETWEEN from_time AND to_time) OR ('%s'::date BETWEEN from_time AND to_time) "+
//		"AND status_id in (Select ID from room_busy_status_categories where status_name='Booked') "+
//		") "+
//		"GROUP BY room_types.type_room_code) as tb1 "+
//		"LEFT JOIN room_types as tb2 on tb1.type_room_code=tb2.type_room_code) as tb3 "+
//		"LEFT JOIN rate_plans as tb4 on tb3.id=tb4.room_type_id where tb3.count>=%d",
//		condition["room_ids"],
//		condition["start_time"],
//		condition["end_time"],
//		condition["number_room"])
//	err := roomReceiver.sql.Db.Raw(queryString).Find(&listRoomAvailable)
//	logger.Debug("Get data", zap.Error(err.Error))
//	if err.RowsAffected == 0 {
//		return listRoomAvailable, nil
//	}
//	return listRoomAvailable, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetNumberRoomAvailableByCapacityAndTimeCheckV2(condition map[string]interface{}) ([]query.RoomRatePlanCostByRoomType, error) {
//	var listRoomAvailable []query.RoomRatePlanCostByRoomType
//	queryString := fmt.Sprintf("select "+
//		"tb5.type_room_code,tb5.type_room_name,tb5.room_images,tb5.number_adult,tb5.number_children,tb5.number_bed,tb5.short_description,tb5.remains,tb6.description,tb5.cost_type,tb6.id as rate_plan_id,tb6.price as rate_plan_price "+
//		"from "+
//		"(select tb4.id, tb4.type_room_name,tb4.room_images, tb4.number_adult, tb4.number_children,tb4.number_bed,tb4.short_description, tb3.type_room_code, tb3.remains,tb4.cost_type from "+
//		"(select tb1.type_room_code, (count1-COALESCE(count2,0)) as remains from "+
//		"(Select room_types.type_room_code,count(rooms.id) as count1 "+
//		"from rooms RIGHT JOIN room_types on rooms.room_type_id = room_types.id "+
//		"where room_types.number_adult>=%d AND room_types.number_children>=%d "+
//		"GROUP BY room_types.type_room_code) as tb1 "+
//		"LEFT JOIN "+
//		"(Select room_types.type_room_code,COALESCE(sum(bookings.number_room),0) as count2 "+
//		"from bookings RIGHT JOIN room_types on bookings.room_type_id = room_types.id "+
//		"where (bookings.status_booking_id = 0 OR bookings.status_booking_id = 1) AND '%s'::date <= check_out_time AND '%s'::date >= check_in_time "+
//		"GROUP BY room_types.type_room_code) as tb2 "+
//		"on tb1.type_room_code=tb2.type_room_code) as tb3 "+
//		"LEFT JOIN room_types as tb4 on tb3.type_room_code=tb4.type_room_code) as tb5 "+
//		"LEFT JOIN rate_plans as tb6 "+
//		"on tb5.id=tb6.room_type_id "+
//		"where remains>=%d",
//		condition["number_adult"],
//		condition["number_children"],
//		condition["start_time"],
//		condition["end_time"],
//		condition["number_room"])
//	err := roomReceiver.sql.Db.Raw(queryString).Find(&listRoomAvailable)
//	logger.Debug("Get get data", zap.Error(err.Error))
//	if err.RowsAffected == 0 {
//		return listRoomAvailable, nil
//	}
//	return listRoomAvailable, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetListRoomCheckOut(condition map[string]interface{}) ([]query.RoomStayedQuery, error) {
//	var listRoomAvailable []query.RoomStayedQuery
//	err := roomReceiver.sql.Db.Raw("select "+
//		"tb1.floor,tb1.room_code,tb2.id as status_detail_id,tb2.booking_id "+
//		"from rooms as tb1 INNER JOIN room_busy_status_details as tb2 "+
//		"on tb1.ID=tb2.room_id "+
//		"where tb2.room_busy_status_category_id='2' AND ?::date >= tb2.from_time",
//		condition["time"]).Find(&listRoomAvailable)
//	logger.Debug("Get get data", zap.Error(err.Error))
//	if err.RowsAffected == 0 {
//		return listRoomAvailable, nil
//	}
//	return listRoomAvailable, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetStatusRoomCustomerInfo(condition map[string]interface{}) (query.RoomStatusInfoQuery, error) {
//	var result query.RoomStatusInfoQuery
//	err := roomReceiver.sql.Db.Raw("select "+
//		"(select full_name from customers where id=tb1.customer_id),tb1.check_in_time,"+
//		"tb1.check_out_time,tb2.room_code,tb2.floor,tb1.id as booking_id "+
//		"from bookings as tb1 inner join booking_details as tb2 on tb1.ID= tb2.booking_id "+
//		"where tb1.id in (?)", condition["booking_id"]).Find(&result)
//	if err != nil {
//		logger.Error("Get get data", zap.Error(err.Error))
//		return result, nil
//	}
//	return result, nil
//}
//
//func (roomReceiver *RoomRepoImpl) GetListRoomCheckIn(roomTypeID string) ([]model.Room, error) {
//	var listRoom []model.Room
//	err := roomReceiver.sql.Db.Where("room_type_id=? AND room_busy_status_category_id = ?", roomTypeID, "e5d5bfbb-64dd-11ed-837f-089798c34e0e").Preload("RoomType").Find(&listRoom)
//	if err != nil {
//		return listRoom, err.Error
//	}
//	return listRoom, nil
//}
//
//func (roomReceiver *RoomRepoImpl) UpdateRoomByID(room model.Room) (model.Room, error) {
//	var roomResult = model.Room{}
//	err := roomReceiver.sql.Db.Model(&roomResult).Where("room_code=?", room.RoomCode).Preload("RoomType").Updates(room)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return roomResult, err.Error
//		}
//
//		return roomResult, custom_error.RoomNotUpdated
//	}
//	return roomResult, nil
//}
