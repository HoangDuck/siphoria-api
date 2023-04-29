package repo_impl

import (
	"gorm.io/gorm"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/model"
	"hotel-booking-api/repository"
)

type AdminRepoImpl struct {
	sql *db.Sql
}

func (u *AdminRepoImpl) AcceptHotel(hotel model.Hotel) (model.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (u *AdminRepoImpl) GetHotelFilter() ([]model.Hotel, error) {
	var listHotel []model.Hotel
	err := u.sql.Db.Find(&listHotel)
	if err != nil {
		return listHotel, err.Error
	}
	return listHotel, err.Error
}

func (u *AdminRepoImpl) GetAccountFilter() ([]model.User, error) {
	var listUser []model.User
	err := u.sql.Db.Find(&listUser)
	if err != nil {
		return listUser, err.Error
	}
	return listUser, err.Error
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
	err := u.sql.Db.Model(&staffAccountResult).Where("id=?", staffAccount.ID).Updates(staffAccount)
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
