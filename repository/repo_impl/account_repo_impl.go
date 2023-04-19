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

func (accountReceiver *AccountRepoImpl) ActivateAccount(account model.User) (model.User, error) {
	err := accountReceiver.sql.Db.Model(&account).
		Where("id=?", account.ID).Update("status", 1)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return account, err.Error
		}
		return account, err.Error
	}
	return account, nil
}

func (accountReceiver *AccountRepoImpl) ResetPassword(email string, password string) (bool, error) {
	var account = model.User{}
	err := accountReceiver.sql.Db.Model(&account).Where("email=?", email).Update("password", password)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return false, err.Error
		}
		return false, err.Error
	}
	if err.RowsAffected == 0 {
		return false, err.Error
	}
	return true, nil
}

func (accountReceiver *AccountRepoImpl) GetCustomerPageUrl() (string, error) {
	var customerPageConfig model.ConfigurationUrlDefine
	err := accountReceiver.sql.Db.Where("id=?", 2).Find(&customerPageConfig)
	logger.Error("Error get template email url ", zap.Error(err.Error))
	if err.Error == gorm.ErrRecordNotFound {
		return "https://siphoria.com/auth/reset?token=", err.Error
	}
	if customerPageConfig.Key == "" && customerPageConfig.Value == "" {
		return "https://siphoria.com/auth/reset?token=", err.Error
	}
	return customerPageConfig.Value, nil
}

func (accountReceiver *AccountRepoImpl) GetCustomerActivatePageUrl() (string, error) {
	var customerPageConfig model.ConfigurationUrlDefine
	err := accountReceiver.sql.Db.Where("id=?", 6).Find(&customerPageConfig)
	logger.Error("Error get template email url ", zap.Error(err.Error))
	if err.Error == gorm.ErrRecordNotFound {
		return "https://siphoria.com/auth/verifyemail/", err.Error
	}
	if customerPageConfig.Key == "" && customerPageConfig.Value == "" {
		return "https://siphoria.com/auth/verifyemail/", err.Error
	}
	return customerPageConfig.Value, nil
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
