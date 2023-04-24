package repo_impl

import (
	"hotel-booking-api/db"
	"hotel-booking-api/repository"
)

type AdminRepoImpl struct {
	sql *db.Sql
}

func NewAdminRepo(sql *db.Sql) repository.AdminRepo {
	return &AdminRepoImpl{
		sql: sql,
	}
}

//func (u *AdminRepoImpl) SavePaymentStatus(paymentStatus model.PaymentStatus) (model.PaymentStatus, error) {
//	result := u.sql.Db.Create(&paymentStatus)
//	if result.Error != nil {
//		return paymentStatus, custom_error.PaymentNotSaved
//	}
//	return paymentStatus, nil
//}
//
//func (u *AdminRepoImpl) SaveStatusWork(statusWork model.StatusWork) (model.StatusWork, error) {
//	result := u.sql.Db.Create(&statusWork)
//	if result.Error != nil {
//		logger.Error("Error insert", zap.Error(result.Error))
//		return statusWork, result.Error
//	}
//	return statusWork, nil
//}
//
//func (u *AdminRepoImpl) SaveStatusAccount(statusAccount model.StatusAccount) (model.StatusAccount, error) {
//	result := u.sql.Db.Create(&statusAccount)
//	if result.Error != nil {
//		return statusAccount, result.Error
//	}
//	return statusAccount, nil
//}
//
//func (u *AdminRepoImpl) SaveStatusBooking(statusBooking model.StatusBooking) (model.StatusBooking, error) {
//	result := u.sql.Db.Create(&statusBooking)
//	if result.Error != nil {
//		logger.Error("Error insert", zap.Error(result.Error))
//		return statusBooking, result.Error
//	}
//	return statusBooking, nil
//}
//
//func (u *AdminRepoImpl) CheckEmail(email string) (model.StaffAccount, error) {
//	var user = model.StaffAccount{}
//	err := u.sql.Db.Where("email = ?", email).Preload("Staff").Preload("StatusAccount").Find(&user)
//	user.Email = email
//	if err.RowsAffected != 0 {
//		user.Email = ""
//		return user, custom_error.EmailAlreadyExists
//	}
//
//	return user, nil
//}
//
//func (u *AdminRepoImpl) GetAccountById(userId string) (model.StaffAccount, error) {
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
//func (u *AdminRepoImpl) ChangeRoleAccount(accountId string, roleName string) (bool, error) {
//	var staffResult = model.StaffAccount{}
//	err := u.sql.Db.Model(&staffResult).Where("id=?", accountId).Update("role", roleName)
//	if err.Error == gorm.ErrRecordNotFound {
//		return false, err.Error
//	}
//
//	return true, nil
//}
//
//func (u *AdminRepoImpl) UpdateStaffAccount(staffAccount model.StaffAccount) (model.StaffAccount, error) {
//	var staffAccountResult = model.StaffAccount{}
//	err := u.sql.Db.Model(&staffAccountResult).Where("id=?", staffAccount.ID).Updates(staffAccount)
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
//func (u *AdminRepoImpl) UpdateStaffProfile(staffProfile model.Staff) (model.Staff, error) {
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
//func (u *AdminRepoImpl) SaveStaffProfile(staffProfile model.Staff) (model.Staff, error) {
//	result := u.sql.Db.Create(&staffProfile)
//	if result.Error != nil {
//		return staffProfile, result.Error
//	}
//	return staffProfile, nil
//}
//
//func (u *AdminRepoImpl) SaveStaffAccount(accountStaff model.StaffAccount) (model.StaffAccount, error) {
//	result := u.sql.Db.Create(&accountStaff)
//	if result.Error != nil {
//		return accountStaff, result.Error
//	}
//	return accountStaff, nil
//}
//
//func (u *AdminRepoImpl) GetStaffProfile(staffId string) (model.Staff, error) {
//	var user = model.Staff{}
//	err := u.sql.Db.Where("id=?", staffId).Preload("StatusWork").Find(&user)
//	if err != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return user, custom_error.UserNotFound
//		}
//		return user, err.Error
//	}
//	return user, nil
//}
//
//func (u *AdminRepoImpl) UpdatePassword(accountId string, hashedPassword string) (bool, error) {
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
//
//func (u *AdminRepoImpl) CheckLogin(email string) (model.StaffAccount, error) {
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
//func (u *AdminRepoImpl) ActivateStaffAccount(account model.StaffAccount) (model.StaffAccount, error) {
//	err := u.sql.Db.Model(&account).
//		Where("id=?", account.ID).Update("status_account_id", 1)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return account, err.Error
//		}
//		return account, err.Error
//	}
//	return account, nil
//}
//
//func (u *AdminRepoImpl) DeactivateStaffAccount(account model.StaffAccount) (model.StaffAccount, error) {
//	err := u.sql.Db.Model(&account).
//		Where("id=?", account.ID).Update("status_account_id", 0)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return account, err.Error
//		}
//		return account, err.Error
//	}
//	return account, nil
//}
//
//func (u *AdminRepoImpl) GetAccountStatusList() ([]model.StatusAccount, error) {
//	var listAccountStatus []model.StatusAccount
//	err := u.sql.Db.Find(&listAccountStatus)
//	if err != nil {
//		return listAccountStatus, err.Error
//	}
//	return listAccountStatus, err.Error
//}
//
//func (u *AdminRepoImpl) GetWorkStatusList() ([]model.StatusWork, error) {
//	var listWorkStatus []model.StatusWork
//	err := u.sql.Db.Find(&listWorkStatus)
//	if err != nil {
//		return listWorkStatus, err.Error
//	}
//	return listWorkStatus, err.Error
//}
//
//func (u *AdminRepoImpl) GetStatisticRevenueByDay(condition map[string]interface{}) ([]query.StatisticRevenueByTimeQuery, error) {
//	var result []query.StatisticRevenueByTimeQuery
//	//queryString := fmt.Sprintf("select "+
//	//	"tb1.payment_time,(select type_room_code from room_types where id in (tb2.room_type_id)) as room_type_code,sum(tb1.amount) "+
//	//	"from payments as tb1 LEFT JOIN bookings as tb2 on tb1.booking_id=tb2.id "+
//	//	"where tb1.status_payment_code='2' AND (tb1.payment_time BETWEEN '%s'::date AND '%s'::date) "+
//	//	"group by tb1.payment_time,tb2.room_type_id",
//	//	condition["time_start"], condition["time_end"])
//	err := u.sql.Db.Raw("select "+
//		"tb1.payment_time,(select type_room_code from room_types where id in (tb2.room_type_id)) as room_type_code,sum(tb1.amount) "+
//		"from payments as tb1 LEFT JOIN bookings as tb2 on tb1.booking_id=tb2.id "+
//		"where tb1.status_payment_code='2' AND (tb1.payment_time BETWEEN ?::date AND ?::date) "+
//		"group by tb1.payment_time,tb2.room_type_id",
//		condition["time_start"], condition["time_end"]).Find(&result)
//	if err != nil {
//		logger.Error("Error statistic data", zap.Error(err.Error))
//		return result, nil
//	}
//	return result, nil
//}
//
//func (u *AdminRepoImpl) GetStatisticRevenueByRoomTypeCode(condition map[string]interface{}) ([]query.StatisticRevenueByTypeRoomCode, error) {
//	var result []query.StatisticRevenueByTypeRoomCode
//	//queryString := fmt.Sprintf("select "+
//	//	"(select type_room_code from room_types where id in (tb2.room_type_id)) as room_type_code, sum (tb1.amount) "+
//	//	"from payments as tb1 LEFT JOIN bookings as tb2 on tb1.booking_id=tb2.id where tb1.status_payment_code='2' "+
//	//	"AND  (tb1.payment_time BETWEEN '%s'::date AND '%s'::date) "+
//	//	"group by tb2.room_type_id",
//	//	condition["time_start"], condition["time_end"])
//	err := u.sql.Db.Raw("select "+
//		"(select type_room_code from room_types where id in (tb2.room_type_id)) as room_type_code, sum (tb1.amount) "+
//		"from payments as tb1 LEFT JOIN bookings as tb2 on tb1.booking_id=tb2.id where tb1.status_payment_code='2' "+
//		"AND  (tb1.payment_time BETWEEN ?::date AND ?::date) "+
//		"group by tb2.room_type_id", condition["time_start"], condition["time_end"]).Find(&result)
//	if err != nil {
//		logger.Error("Error statistic data", zap.Error(err.Error))
//		return result, nil
//	}
//	return result, nil
//}
//
//func (u *AdminRepoImpl) GetAllCustomer() ([]model.Customer, error) {
//	var listCustomer []model.Customer
//	err := u.sql.Db.Find(&listCustomer)
//	if err != nil {
//		return listCustomer, err.Error
//	}
//	return listCustomer, err.Error
//}
//
//func (u *AdminRepoImpl) GetAllStaff() ([]model.Staff, error) {
//	var listStaff []model.Staff
//	err := u.sql.Db.Preload("StatusWork").Find(&listStaff)
//	if err != nil {
//		return listStaff, err.Error
//	}
//	return listStaff, err.Error
//}
//
//func (u *AdminRepoImpl) GetAllStaffAccount() ([]model.StaffAccount, error) {
//	var listStaffAccount []model.StaffAccount
//	err := u.sql.Db.Preload("Staff").Preload("StatusAccount").Find(&listStaffAccount)
//	if err != nil {
//		return listStaffAccount, err.Error
//	}
//	return listStaffAccount, err.Error
//}
//
//func (u *AdminRepoImpl) GetAllCustomerAccount() ([]model.Account, error) {
//	var listAccount []model.Account
//	err := u.sql.Db.Preload("Customer").Preload("StatusAccount").Find(&listAccount)
//	if err != nil {
//		return listAccount, err.Error
//	}
//	return listAccount, err.Error
//}
