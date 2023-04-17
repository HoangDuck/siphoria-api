package repo_impl

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
)

type AccountRepoImpl struct {
	sql *db.Sql
}

func (accountReceiver *AccountRepoImpl) CheckEmailExisted(email string) (bool, error) {
	var user = model.User{}
	err := accountReceiver.sql.Db.Where("email in (?)", email).Find(&user)
	if err != nil {
		logger.Error("Error query database", zap.Error(err.Error))
		if err.RowsAffected != 0 {
			return true, custom_error.EmailAlreadyExists
		}
	}
	return false, nil
}

func (accountReceiver *AccountRepoImpl) SaveAccount(account model.User) (model.User, error) {
	result := accountReceiver.sql.Db.Create(&account)
	if result != nil {
		logger.Error("Error query database", zap.Error(result.Error))
		if result.Error != nil {
			logger.Error("Error insert", zap.Error(result.Error))
			return account, result.Error
		}
	}
	return account, nil
}

func (accountReceiver *AccountRepoImpl) GetAccountById(userId string) (model.User, error) {
	var user = model.User{}
	err := accountReceiver.sql.Db.Where("id=?", userId).Find(&user)
	if err != nil {
		logger.Error("Error query database", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return user, custom_error.UserNotFound
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return user, err.Error
		}
		return user, err.Error
	}
	return user, nil
}

func (accountReceiver *AccountRepoImpl) CheckLogin(request req.RequestSignIn) (model.User, error) {
	var user = model.User{}
	err := accountReceiver.sql.Db.Where("email=?", request.Email).Find(&user)
	if err != nil {
		logger.Error("Error query database", zap.Error(err.Error))

		if err.Error == gorm.ErrRecordNotFound {
			return user, custom_error.UserNotFound
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return user, err.Error
		}
		return user, err.Error
	}
	return user, nil
}

func NewAccountRepo(sql *db.Sql) repository.AccountRepo {
	return &AccountRepoImpl{
		sql: sql,
	}
}
