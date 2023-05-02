package repo_impl

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/repository"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func (u *UserRepoImpl) GetUserNotifications() (model.Notification, error) {
	//TODO implement me
	panic("implement me")
}

//func (u *UserRepoImpl) CheckProfileCustomerExistByIdentify(user model.Customer) (model.Customer, error) {
//	err := u.sql.Db.Where("identifier_number=?", user.IdentifierNumber).Find(&user)
//	if err != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return user, custom_error.UserNotFound
//		}
//		return user, err.Error
//	}
//	return user, nil
//}

//func (u *UserRepoImpl) SaveCustomerProfile(user model.Customer) (model.Customer, error) {
//	result := u.sql.Db.Create(&user)
//	if result.Error != nil {
//		return user, result.Error
//	}
//	return user, nil
//}

//func (u *UserRepoImpl) GetUserCart(customer model.User) (query.RoomAvailableQuery, error) {
//	var cartData = model.User{}
//	err := u.sql.Db.Where("id=?", customer.ID).Find(&cartData)
//	if err != nil {
//		logger.Error("Error query data", zap.Error(err.Error))
//		if err.Error == gorm.ErrRecordNotFound {
//			return cartData, err.Error
//		}
//		if err.Error == gorm.ErrInvalidTransaction {
//			return cartData, err.Error
//		}
//		return cartData, err.Error
//	}
//	return cartData, nil
//}

func (u *UserRepoImpl) GetProfileCustomer(customer model.User) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.Where("id=?", customer.ID).Find(&user)
	if err != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return user, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return user, err.Error
		}
		return user, err.Error
	}
	return user, nil
}

func (u *UserRepoImpl) UpdateRankCustomer(userRank model.UserRank) (model.UserRank, error) {
	err := u.sql.Db.Model(&userRank).Where("user_id=?", userRank.UserId).Updates(userRank)
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return userRank, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return userRank, err.Error
		}
		return userRank, err.Error
	}
	return userRank, nil
}

func (u *UserRepoImpl) UpdateProfileCustomer(user model.User) (model.User, error) {
	err := u.sql.Db.Model(&user).Updates(user)
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return user, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return user, err.Error
		}
		return user, err.Error
	}
	return user, nil
}

func (u *UserRepoImpl) GetUserRank(customer model.User) (model.UserRank, error) {
	var userRank = model.UserRank{}
	err := u.sql.Db.Where("id=?", customer.ID).Find(&userRank)
	if err != nil {
		logger.Error("Error query database", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return userRank, custom_error.UserNotFound
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return userRank, err.Error
		}
		return userRank, err.Error
	}
	return userRank, nil
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}
