package repo_impl

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
	"hotel-booking-api/repository"
)

type AdminRepoImpl struct {
	sql *db.Sql
}

func (u *AdminRepoImpl) ApprovePayoutRequestHotel(hotelPayoutRequest model.PayoutRequest) (model.PayoutRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (u *AdminRepoImpl) AcceptHotel(hotel model.Hotel) (model.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (u *AdminRepoImpl) GetHotelFilter(queryModel query.DataQueryModel) ([]model.Hotel, error) {
	var listHotel []model.Hotel
	err := GenerateQueryGetData(u.sql, queryModel, &model.Hotel{}, queryModel.ListIgnoreColumns)
	err = err.Find(&listHotel)
	if err.Error != nil {
		logger.Error("Error get list hotel url ", zap.Error(err.Error))
		return listHotel, err.Error
	}
	return listHotel, nil
}

func (u *AdminRepoImpl) GetAccountFilter(queryModel query.DataQueryModel) ([]model.User, error) {
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
	var staffAccountResult = model.User{}
	err := u.sql.Db.Model(&staffAccountResult).Updates(staffAccount)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return staffAccountResult, err.Error
		}

		return staffAccountResult, custom_error.UserNotUpdated
	}
	return staffAccountResult, nil
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
