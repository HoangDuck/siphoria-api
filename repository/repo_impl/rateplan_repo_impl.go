package repo_impl

import (
	"gorm.io/gorm"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/model"
	"hotel-booking-api/repository"
)

type RatePlanRepoImpl struct {
	sql *db.Sql
}

func NewRatePlanRepo(sql *db.Sql) repository.RatePlanRepo {
	return &RatePlanRepoImpl{
		sql: sql,
	}
}

func (ratePlanReceiver *RatePlanRepoImpl) GetListRatePlanByRoomTypeCode(roomTypeID string) ([]model.RatePlan, error) {
	var listRatePlan []model.RatePlan
	err := ratePlanReceiver.sql.Db.Where("room_type_id=?", roomTypeID).Preload("RoomType").Find(&listRatePlan)
	if err != nil {
		return listRatePlan, err.Error
	}
	return listRatePlan, nil
}

func (ratePlanReceiver *RatePlanRepoImpl) GetListRatePlan() ([]model.RatePlan, error) {
	var listRatePlan []model.RatePlan
	err := ratePlanReceiver.sql.Db.Preload("RoomType").Find(&listRatePlan)
	if err != nil {
		return listRatePlan, err.Error
	}
	return listRatePlan, nil
}

func (ratePlanReceiver *RatePlanRepoImpl) GetRatePlanInfo(ratePlan model.RatePlan) (model.RatePlan, error) {
	var ratePlanInfo = model.RatePlan{}
	err := ratePlanReceiver.sql.Db.Where("id = ?", ratePlan.ID).Preload("RoomType").Find(&ratePlanInfo)
	if err.RowsAffected == 0 {
		return ratePlanInfo, err.Error
	}
	return ratePlanInfo, nil
}

func (ratePlanReceiver *RatePlanRepoImpl) UpdateRatePlanInfo(ratePlan model.RatePlan) (model.RatePlan, error) {
	var ratePlanResult = model.RatePlan{}
	err := ratePlanReceiver.sql.Db.Model(&ratePlanResult).Where("id=?", ratePlan.ID).Preload("RoomType").Updates(ratePlan)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return ratePlanResult, err.Error
		}

		return ratePlanResult, custom_error.RoomNotUpdated
	}
	return ratePlanResult, nil
}

func (ratePlanReceiver *RatePlanRepoImpl) DeleteRatePlanInfo(condition map[string]interface{}) (bool, error) {
	//err := ratePlanReceiver.sql.Db.Where("id = ?", condition["ID"]).Delete(&model.Room{})
	//if err.RowsAffected == 0 {
	//	return false, custom_error.RoomDeleteFailed
	//}
	return true, nil
}

func (ratePlanReceiver *RatePlanRepoImpl) SaveRatePlan(ratePlan model.RatePlan) (model.RatePlan, error) {
	result := ratePlanReceiver.sql.Db.Create(&ratePlan)
	if result.Error != nil {
		return ratePlan, result.Error
	}
	return ratePlan, nil
}

func (ratePlanReceiver *RatePlanRepoImpl) GetRoomTypeInfo(roomType model.RoomType) (model.RoomType, error) {
	var roomTypeResult = model.RoomType{}
	err := ratePlanReceiver.sql.Db.Where("type_room_code =?", roomType.Name).Find(&roomTypeResult)
	if err != nil {
		return roomTypeResult, err.Error
	}
	return roomTypeResult, nil
}
