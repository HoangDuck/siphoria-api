package repo_impl

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
	"hotel-booking-api/utils"
	"strconv"
	"time"
)

type RoomRepoImpl struct {
	sql *db.Sql
}

func (roomReceiver *RoomRepoImpl) UpdateLockRoom() {
	logger.Info("Update lock room")
	err := roomReceiver.sql.Db.Where("lock_to < current_timestamp AND expired = false").Updates(model.LockRoom{Expired: true})
	if err.Error != nil {
		logger.Error("Error update lock room", zap.Error(err.Error))

	}
}

func (roomReceiver *RoomRepoImpl) GetRatePlanByRoomTypeFilter(queryModel *query.DataQueryModel) ([]model.RatePlan, error) {
	var listRatePlan []model.RatePlan
	err := GenerateQueryGetData(roomReceiver.sql, queryModel, &model.RatePlan{}, queryModel.ListIgnoreColumns)
	err = err.Where("room_type_id = ?", queryModel.DataId)
	var countTotalRows int64
	err.Model(model.RatePlan{}).Count(&countTotalRows)
	queryModel.TotalRows = int(countTotalRows)
	err = err.Find(&listRatePlan)
	if err.Error != nil {
		logger.Error("Error get list rate plan url ", zap.Error(err.Error))
		return listRatePlan, err.Error
	}
	return listRatePlan, nil
}

func (roomReceiver *RoomRepoImpl) GetRoomTypeFacility(roomTypeId string) (model.RoomTypeFacility, error) {
	var roomTypeFacility model.RoomTypeFacility
	err := roomReceiver.sql.Db.Where("room_type_id = ?", roomTypeId).Find(&roomTypeFacility)
	if err.RowsAffected == 0 {
		return roomTypeFacility, err.Error
	}
	return roomTypeFacility, nil
}

func (roomReceiver *RoomRepoImpl) GetRoomTypeViews(roomTypeId string) (model.RoomTypeViews, error) {
	var roomTypeViews model.RoomTypeViews
	err := roomReceiver.sql.Db.Where("room_type_id = ?", roomTypeId).Find(&roomTypeViews)
	if err.RowsAffected == 0 {
		return roomTypeViews, err.Error
	}
	return roomTypeViews, nil
}

func (roomReceiver *RoomRepoImpl) GetRoomNightsByRoomType(c echo.Context, roomTypeId string) ([]model.RoomNights, error) {
	var roomNightList []model.RoomNights
	err := roomReceiver.sql.Db.Where("room_type_id = ?", roomTypeId)
	if c.QueryParam("month") != "" {
		monthValue, _ := strconv.ParseInt(c.QueryParam("month"), 10, 5)
		err = err.Where("DATE_PART('month', availability_at)  = ?", fmt.Sprintf("%v", monthValue+1))
	}
	if c.QueryParam("year") != "" {
		err = err.Where("DATE_PART('year', availability_at)  = ?", c.QueryParam("year"))
	}
	from := c.QueryParam("from")
	to := c.QueryParam("to")
	if from != "" {
		err = err.Where("availability_at >= ?", from)
	}
	if to != "" {
		err = err.Where("availability_at < ?", to)
	}
	err = err.Find(&roomNightList)
	if err.RowsAffected == 0 {
		return roomNightList, err.Error
	}
	return roomNightList, nil
}

func (roomReceiver *RoomRepoImpl) GetListRatePlans(c echo.Context, roomTypeId string) ([]model.RatePlan, error) {
	var ratePlanList []model.RatePlan
	err := roomReceiver.sql.Db.Where("room_type_id = ?", roomTypeId)
	err = err.Find(&ratePlanList)
	if err.RowsAffected == 0 {
		return ratePlanList, err.Error
	}
	return ratePlanList, nil
}

func (roomReceiver *RoomRepoImpl) GetListRatePackages(c echo.Context, ratePlanId string) ([]model.RatePackage, error) {
	var ratePackageList []model.RatePackage
	err := roomReceiver.sql.Db.Where("rate_plan_id = ?", ratePlanId)

	if c.QueryParam("month") != "" {
		monthValue, _ := strconv.ParseInt(c.QueryParam("month"), 10, 5)
		err = err.Where("DATE_PART('month', availability_at)  = ?", fmt.Sprintf("%v", monthValue+1))
	}
	if c.QueryParam("year") != "" {
		err = err.Where("DATE_PART('year', availability_at)  = ?", c.QueryParam("year"))
	}

	from := c.QueryParam("from")
	to := c.QueryParam("to")
	if from != "" {
		err = err.Where("availability_at >= ?", from)
	}
	if to != "" {
		err = err.Where("availability_at < ?", to)
	}
	err = err.Find(&ratePackageList)
	if err.RowsAffected == 0 {
		return ratePackageList, err.Error
	}
	return ratePackageList, nil
}

func (roomReceiver *RoomRepoImpl) GetListRoomTypeDetail(room model.RoomType) ([]model.RoomType, error) {
	var roomTypeList []model.RoomType
	err := roomReceiver.sql.Db.Where("hotel_id = ?", room.HotelId).Find(&roomTypeList)
	if err.RowsAffected == 0 {
		return roomTypeList, err.Error
	}
	return roomTypeList, nil
}

func (roomReceiver *RoomRepoImpl) GetRoomTypeDetail(room model.RoomType) (model.RoomType, error) {
	err := roomReceiver.sql.Db.Preload("RoomTypeFacility").Preload("RoomTypeViews").
		Where("id = ?", room.ID).Find(&room)
	if err.RowsAffected == 0 {
		return room, custom_error.RoomNotFound
	}
	if err.Error != nil {
		return room, err.Error
	}
	return room, nil
}

func (roomReceiver *RoomRepoImpl) UpdateRoomPhotos(room model.RoomType) (model.RoomType, error) {
	err := roomReceiver.sql.Db.Model(&room).Select("photos").Updates(map[string]interface{}{"photos": room.Photos})
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
		Ocean:          requestUpdateRoomType.Ocean,
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
		AirConditional: requestUpdateRoomType.AirConditional,
		Tivi:           requestUpdateRoomType.Tivi,
		Kitchen:        requestUpdateRoomType.Kitchen,
		PrivatePool:    requestUpdateRoomType.PrivatePool,
		Heater:         requestUpdateRoomType.Heater,
		Iron:           requestUpdateRoomType.Iron,
		Sofa:           requestUpdateRoomType.Sofa,
		Desk:           requestUpdateRoomType.Desk,
		Soundproof:     requestUpdateRoomType.Soundproof,
		Towels:         requestUpdateRoomType.Towels,
		Toiletries:     requestUpdateRoomType.Toiletries,
		Shower:         requestUpdateRoomType.Shower,
		Slippers:       requestUpdateRoomType.Slippers,
		Hairdry:        requestUpdateRoomType.Hairdry,
		Fruit:          requestUpdateRoomType.Fuirt,
		Bbq:            requestUpdateRoomType.Bbq,
		Wine:           requestUpdateRoomType.Wine,
		Fryer:          requestUpdateRoomType.Fryer,
		KitchenTool:    requestUpdateRoomType.KitchenTool,
		IsDeleted:      requestUpdateRoomType.IsDelete,
	}
	err := roomReceiver.sql.Db.Model(&roomType).Updates(roomType)
	err = roomReceiver.sql.Db.Select("activated").Model(&roomType).Updates(roomType)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return roomType, err.Error
		}

		return roomType, custom_error.RoomNotUpdated
	}
	err = roomReceiver.sql.Db.Model(&roomTypeViews).Updates(utils.ConvertStructToMap(&roomTypeViews, []string{
		"room_type_id", "created_at", "updated_at", "-",
	}))
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return roomType, err.Error
		}

		return roomType, custom_error.RoomNotUpdated
	}
	err = roomReceiver.sql.Db.Model(&roomTypeFacility).Updates(utils.ConvertStructToMap(&roomTypeFacility, []string{
		"room_type_id", "created_at", "updated_at", "-",
	}))
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
	listDayAvailable := utils.DecodeJSONArray(requestUpdateRoomNight.SelectedDate, false)
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
