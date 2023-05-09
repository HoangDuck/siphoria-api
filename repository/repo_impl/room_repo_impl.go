package repo_impl

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
	"hotel-booking-api/utils"
	"time"
)

type RoomRepoImpl struct {
	sql *db.Sql
}

func (roomReceiver *RoomRepoImpl) GetRoomTypeDetail(room model.RoomType) (model.RoomType, error) {
	err := roomReceiver.sql.Db.Where("id = ?", room.ID).Find(&room)
	if err.RowsAffected == 0 {
		return room, err.Error
	}
	return room, nil
}

func (roomReceiver *RoomRepoImpl) UpdateRoomPhotos(room model.RoomType) (model.RoomType, error) {
	err := roomReceiver.sql.Db.Model(&room).Updates(room)
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return room, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return room, err.Error
		}
		return room, err.Error
	}
	return room, nil
}

func (roomReceiver *RoomRepoImpl) UpdateRoomType(requestUpdateRoomType req.RequestUpdateRoomType, idRoomType string) (model.RoomType, error) {
	roomType := model.RoomType{
		ID:           idRoomType,
		Activated:    requestUpdateRoomType.Activated,
		Name:         requestUpdateRoomType.Name,
		Description:  requestUpdateRoomType.Description,
		MaxAdult:     requestUpdateRoomType.MaxAdult,
		MaxChildren:  requestUpdateRoomType.MaxChildren,
		BedNums:      requestUpdateRoomType.BedNums,
		BathroomNums: requestUpdateRoomType.BathroomNums,
		IsDeleted:    requestUpdateRoomType.IsDelete,
	}
	roomTypeViews := model.RoomTypeViews{
		RoomTypeID:     idRoomType,
		Bay:            requestUpdateRoomType.Bay,
		Sea:            requestUpdateRoomType.Sea,
		City:           requestUpdateRoomType.City,
		Garden:         requestUpdateRoomType.Garden,
		Lake:           requestUpdateRoomType.Lake,
		Mountain:       requestUpdateRoomType.Mountain,
		River:          requestUpdateRoomType.River,
		PrivateBalcony: requestUpdateRoomType.PrivateBalcony,
		IsDeleted:      requestUpdateRoomType.IsDelete,
	}
	roomTypeFacility := model.RoomTypeFacility{
		RoomTypeID:     idRoomType,
		AirConditioner: requestUpdateRoomType.AirConditional,
		TV:             requestUpdateRoomType.Tivi,
		Kitchen:        requestUpdateRoomType.Kitchen,
		PrivatePool:    requestUpdateRoomType.PrivatePool,
		Heater:         requestUpdateRoomType.Heater,
		Iron:           requestUpdateRoomType.Iron,
		Sofa:           requestUpdateRoomType.Sofa,
		Desk:           requestUpdateRoomType.Desk,
		SoundProof:     requestUpdateRoomType.Soundproof,
		Towels:         requestUpdateRoomType.Towels,
		Toiletries:     requestUpdateRoomType.Toiletries,
		Shower:         requestUpdateRoomType.Shower,
		Slipper:        requestUpdateRoomType.Slippers,
		HairDry:        requestUpdateRoomType.Hairdry,
		Fruit:          requestUpdateRoomType.Fuirt,
		Bbq:            requestUpdateRoomType.Bbq,
		Wine:           requestUpdateRoomType.Wine,
		Fryer:          requestUpdateRoomType.Fryer,
		KitchenTools:   requestUpdateRoomType.KitchenTool,
		IsDeleted:      requestUpdateRoomType.IsDelete,
	}
	err := roomReceiver.sql.Db.Model(&roomType).Updates(roomType)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return roomType, err.Error
		}

		return roomType, custom_error.RoomNotUpdated
	}
	err = roomReceiver.sql.Db.Model(&roomTypeViews).Updates(roomTypeViews)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return roomType, err.Error
		}

		return roomType, custom_error.RoomNotUpdated
	}
	err = roomReceiver.sql.Db.Model(&roomTypeFacility).Updates(roomTypeFacility)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return roomType, err.Error
		}

		return roomType, custom_error.RoomNotUpdated
	}
	return roomType, nil
}

func (roomReceiver *RoomRepoImpl) UpdateRatePackages(requestAddRatePackages req.RequestUpdateRatePackage) ([]model.RatePackage, error) {
	var listRatePackages []model.RatePackage
	listRatePackages = []model.RatePackage{}
	for _, dataRatePackageItem := range requestAddRatePackages.Data {
		ratePackageId, err := utils.GetNewId()
		date, err := time.Parse("2006-01-02", dataRatePackageItem.Date)
		if err != nil {
			continue
		}
		listRatePackages = append(listRatePackages, model.RatePackage{
			ID:             ratePackageId,
			AvailabilityAt: date,
			Price:          dataRatePackageItem.Price,
			RatePlanId:     dataRatePackageItem.RatePlan,
			Currency:       "VND",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})
	}
	if len(listRatePackages) == 0 {
		return listRatePackages, nil
	}
	err := roomReceiver.sql.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "rate_plan_id"}, {Name: "availability_at"}},
		DoUpdates: clause.AssignmentColumns([]string{"price", "updated_at"}),
	}).Create(&listRatePackages).Error
	if err != nil {
		return listRatePackages, err
	}
	return listRatePackages, nil
}

func (roomReceiver *RoomRepoImpl) UpdateRoomNight(requestUpdateRoomNight req.RequestUpdateRoomNight) ([]model.RoomNights, error) {
	listDayAvailable := utils.DecodeJSONArray(requestUpdateRoomNight.SelectedDate)
	var listRoomNights []model.RoomNights
	listRoomNights = []model.RoomNights{}
	for _, roomNightDay := range listDayAvailable {
		roomNightId, err := utils.GetNewId()
		date, err := time.Parse("2006-01-02", roomNightDay)
		if err != nil {
			continue
		}
		listRoomNights = append(listRoomNights, model.RoomNights{
			ID:             roomNightId,
			AvailabilityAt: date,
			RoomTypeId:     requestUpdateRoomNight.RoomTypeID,
			Inventory:      requestUpdateRoomNight.Quantity,
			Remain:         requestUpdateRoomNight.Quantity,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})
	}
	if len(listRoomNights) == 0 {
		return listRoomNights, nil
	}
	err := roomReceiver.sql.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "room_type_id"}, {Name: "availability_at"}},
		DoUpdates: clause.AssignmentColumns([]string{"inventory", "remain", "updated_at"}),
	}).Create(&listRoomNights).Error
	if err != nil {
		return listRoomNights, err
	}
	return listRoomNights, nil
}

func (roomReceiver *RoomRepoImpl) SaveRoomType(requestAddRoomType req.RequestCreateRoomType) (model.RoomType, error) {
	err := roomReceiver.sql.Db.Exec("call sp_addroomtype(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);",
		requestAddRoomType.ID,
		requestAddRoomType.Name,
		requestAddRoomType.Description,
		requestAddRoomType.BedNums,
		requestAddRoomType.BathroomNums,
		requestAddRoomType.Activated,
		requestAddRoomType.MaxChildren,
		requestAddRoomType.MaxAdult,
		requestAddRoomType.HotelID,
		requestAddRoomType.Photos,
		requestAddRoomType.Bay,
		requestAddRoomType.Ocean,
		requestAddRoomType.City,
		requestAddRoomType.Garden,
		requestAddRoomType.Lake,
		requestAddRoomType.Mountain,
		requestAddRoomType.River,
		requestAddRoomType.PrivateBalcony,
		requestAddRoomType.AirConditional,
		requestAddRoomType.Tivi,
		requestAddRoomType.Kitchen,
		requestAddRoomType.PrivatePool,
		requestAddRoomType.Heater,
		requestAddRoomType.Iron,
		requestAddRoomType.Sofa,
		requestAddRoomType.Desk,
		requestAddRoomType.Soundproof,
		requestAddRoomType.Towels,
		requestAddRoomType.Toiletries,
		requestAddRoomType.Shower,
		requestAddRoomType.Slippers,
		requestAddRoomType.Hairdry,
		requestAddRoomType.Fruit,
		requestAddRoomType.Bbq,
		requestAddRoomType.Wine,
		requestAddRoomType.Fryer,
		requestAddRoomType.KitchenTool)
	var roomType model.RoomType
	if err.Error != nil {
		return roomType, err.Error
	} else {
		roomType = model.RoomType{
			ID:           requestAddRoomType.ID,
			Name:         requestAddRoomType.Name,
			Description:  requestAddRoomType.Description,
			MaxAdult:     requestAddRoomType.MaxAdult,
			MaxChildren:  requestAddRoomType.MaxChildren,
			BedNums:      requestAddRoomType.BedNums,
			BathroomNums: requestAddRoomType.BathroomNums,
			Photos:       requestAddRoomType.Photos,
			HotelId:      requestAddRoomType.HotelID,
		}
	}
	return roomType, nil
}

func NewRoomRepo(sql *db.Sql) repository.RoomRepo {
	return &RoomRepoImpl{
		sql: sql,
	}
}
