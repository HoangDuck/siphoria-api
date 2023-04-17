package repo_impl

//
//import (
//	"gorm.io/gorm"
//	"hotel-booking-api/custom_error"
//	"hotel-booking-api/db"
//	"hotel-booking-api/model"
//	"hotel-booking-api/repository"
//)
//
//type StaffRepoImpl struct {
//	sql *db.Sql
//}
//
//func NewStaffRepo(sql *db.Sql) repository.StaffRepo {
//	return &StaffRepoImpl{
//		sql: sql,
//	}
//}
//
//func (u *StaffRepoImpl) CheckEmail(email string) (bool, error) {
//	var user = model.StaffAccount{}
//	err := u.sql.Db.Where("email = ?", email).Preload("StatusAccount").Preload("Staff").Find(&user)
//	user.Email = email
//	if err.RowsAffected != 0 {
//		user.Email = ""
//		return false, custom_error.EmailAlreadyExists
//	}
//
//	return true, nil
//}
//
//func (u *StaffRepoImpl) UpdateAvatarInfo(accountId string, avatarUrl string) (model.StaffAccount, error) {
//	var staffAccountResult = model.StaffAccount{}
//	var staffUpdate = model.StaffAccount{
//		Avatar: avatarUrl,
//	}
//	err := u.sql.Db.Model(&staffAccountResult).Where("id=?", accountId).Updates(staffUpdate)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return staffAccountResult, err.Error
//		}
//
//		return staffAccountResult, custom_error.UserNotUpdated
//	}
//	return staffAccountResult, nil
//}
//
//func (u *StaffRepoImpl) UpdatePersonalInfo(staffProfile model.Staff) (model.Staff, error) {
//	var staffResult = model.Staff{}
//	err := u.sql.Db.Model(&staffResult).Where("id=?", staffProfile.ID).Updates(staffProfile)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return staffResult, err.Error
//		}
//
//		return staffResult, custom_error.UserNotUpdated
//	}
//	return staffResult, nil
//}
//
//func (u *StaffRepoImpl) GetAccountById(userId string) (model.StaffAccount, error) {
//	var user = model.StaffAccount{}
//	err := u.sql.Db.Where("id=?", userId).Preload("Staff").Preload("StatusAccount").Find(&user)
//	if err != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return user, custom_error.UserNotFound
//		}
//		return user, err.Error
//	}
//	return user, nil
//}
//
//func (u *StaffRepoImpl) CheckLogin(email string) (model.StaffAccount, error) {
//	var user = model.StaffAccount{}
//	err := u.sql.Db.Where("email=?", email).Preload("Staff").Preload("StatusAccount").Find(&user)
//	if err != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return user, custom_error.UserNotFound
//		}
//		return user, err.Error
//	}
//	return user, nil
//}
//
//func (u *StaffRepoImpl) UpdatePassword(accountId string, hashedPassword string) (bool, error) {
//	var account = model.StaffAccount{}
//	err := u.sql.Db.Model(&account).Where("id=?", accountId).Update("password", hashedPassword)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return false, err.Error
//		}
//		return false, err.Error
//	}
//	return true, nil
//}
