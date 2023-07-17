package repo_impl

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
)

type AdminRepoImpl struct {
	sql *db.Sql
}

func (u *AdminRepoImpl) GetPayoutRequest(queryModel *query.DataQueryModel) ([]model.PayoutRequest, error) {
	var listPayoutRequest []model.PayoutRequest
	err := GenerateQueryGetData(u.sql, queryModel, &model.PayoutRequest{}, queryModel.ListIgnoreColumns)
	err = err.Preload("Hotel").Preload("Pettioner").Preload("Payer")
	var countTotalRows int64
	err.Model(&model.PayoutRequest{}).Count(&countTotalRows)
	queryModel.TotalRows = int(countTotalRows)
	err = err.Find(&listPayoutRequest)
	if err.Error != nil {
		logger.Error("Error get list hotel work url ", zap.Error(err.Error))
		return listPayoutRequest, err.Error
	}
	return listPayoutRequest, nil
}

func (u *AdminRepoImpl) SaveHotelWorkByEmployee(hotelWork model.HotelWork) (model.HotelWork, error) {
	result := u.sql.Db.Create(&hotelWork)
	if result.Error != nil {
		return hotelWork, result.Error
	}
	return hotelWork, nil
}

func (u *AdminRepoImpl) DeleteHotelWorkByEmployee(requestDeleteHotelWorkByEmployee req.RequestDeleteHotelWorkByEmployee) (bool, error) {
	hotelWork := model.HotelWork{
		IsDeleted: true,
	}
	err := u.sql.Db.Model(&hotelWork).
		Where("hotel_id = ? AND user_id = ?", requestDeleteHotelWorkByEmployee.HotelId,
			requestDeleteHotelWorkByEmployee.UserId).Updates(hotelWork)
	if err.Error != nil {
		logger.Error("Error update hotel failed ", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return false, err.Error
		}
		return false, custom_error.UserNotUpdated
	}
	return true, nil
}

func (u *AdminRepoImpl) GetHotelWorkByEmployee(queryModel *query.DataQueryModel) ([]model.Hotel, error) {
	var listHotelWork []model.Hotel
	err := GenerateQueryGetData(u.sql, queryModel, &model.Hotel{}, queryModel.ListIgnoreColumns)
	err = err.Preload("HotelType").Preload("HotelFacility").Where("id in (select hotel_id from hotel_works where user_id = ? AND is_deleted = ?)", queryModel.UserId, queryModel.IsShowDeleted)
	err = err.Find(&listHotelWork)
	if err.Error != nil {
		logger.Error("Error get list hotel work url ", zap.Error(err.Error))
		return listHotelWork, err.Error
	}
	logger.Info("get list hotel work url ", zap.Error(err.Error))
	return listHotelWork, nil
}

func (u *AdminRepoImpl) UpdateRatingHotel(hotel model.Hotel) (model.Hotel, error) {
	err := u.sql.Db.Model(&hotel).Updates(hotel)
	if err.Error != nil {
		logger.Error("Error update hotel failed ", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}
		return hotel, custom_error.HotelNotUpdated
	}
	return hotel, nil
}

func (u *AdminRepoImpl) UpdateCommissionRatingHotel(hotel model.Hotel) (model.Hotel, error) {
	err := u.sql.Db.Model(&hotel).Updates(hotel)
	if err.Error != nil {
		logger.Error("Error update hotel failed ", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}
		return hotel, custom_error.HotelNotUpdated
	}
	return hotel, nil
}

func (u *AdminRepoImpl) ApprovePayoutRequestHotel(hotelPayoutRequest model.PayoutRequest) (model.PayoutRequest, error) {
	err := u.sql.Db.Model(&hotelPayoutRequest).Where("id = ?", hotelPayoutRequest.ID).Updates(hotelPayoutRequest)
	if err.Error != nil {
		logger.Error("Error update payment failed ", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return hotelPayoutRequest, err.Error
		}
		return hotelPayoutRequest, custom_error.PaymentNotUpdated
	}
	return hotelPayoutRequest, nil
}

func (u *AdminRepoImpl) AcceptHotel(hotel model.Hotel) (model.Hotel, error) {
	err := u.sql.Db.Model(&hotel).Updates(hotel)
	if err.Error != nil {
		logger.Error("Error update user failed ", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}
		return hotel, custom_error.HotelNotUpdated
	}
	return hotel, nil
}

func (u *AdminRepoImpl) GetHotelFilter(queryModel *query.DataQueryModel) ([]model.Hotel, error) {
	var listHotel []model.Hotel
	err := GenerateQueryGetData(u.sql, queryModel, &model.Hotel{}, queryModel.ListIgnoreColumns)
	err = err.Preload("HotelType").Preload("HotelFacility").Find(&listHotel)
	if err.Error != nil {
		logger.Error("Error get list hotel url ", zap.Error(err.Error))
		return listHotel, err.Error
	}
	return listHotel, nil
}

func (u *AdminRepoImpl) GetAccountFilter(queryModel *query.DataQueryModel) ([]model.User, error) {
	var listUser []model.User
	err := GenerateQueryGetData(u.sql, queryModel, &model.User{}, queryModel.ListIgnoreColumns)
	err = err.Find(&listUser)
	if err.Error != nil {
		logger.Error("Error get list user url ", zap.Error(err.Error))
		return listUser, err.Error
	}
	return listUser, nil
}

func (u *AdminRepoImpl) CheckEmail(email string) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.Where("email = ?", email).Preload("Staff").Preload("StatusAccount").Find(&user)
	user.Email = email
	if err.RowsAffected != 0 {
		user.Email = ""
		return user, custom_error.EmailAlreadyExists
	}

	return user, nil
}

func (u *AdminRepoImpl) UpdateAccount(staffAccount model.User) (model.User, error) {
	err := u.sql.Db.Model(&staffAccount).Updates(staffAccount)
	if err.Error != nil {
		logger.Error("Error update user failed ", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return staffAccount, err.Error
		}
		return staffAccount, custom_error.UserNotUpdated
	}
	return staffAccount, nil
}

func (u *AdminRepoImpl) SaveAccount(account model.User) (model.User, error) {
	result := u.sql.Db.Create(&account)
	if result.Error != nil {
		return account, result.Error
	}
	return account, nil
}

func NewAdminRepo(sql *db.Sql) repository.AdminRepo {
	return &AdminRepoImpl{
		sql: sql,
	}
}
