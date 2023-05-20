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
	"hotel-booking-api/utils"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func (userReceiver *UserRepoImpl) GetUserPaymentPendingCheckin(user model.User) ([]model.Payment, error) {
	var listPaymentUser []model.Payment
	err := userReceiver.sql.Db.Where("user_id = ? AND status = ?", user.ID, "paid").Preload("User").
		Preload("RatePlan").Preload("RoomType").Preload("Hotel").Find(&listPaymentUser)
	if err.Error != nil {
		logger.Error("Error get list cart url ", zap.Error(err.Error))
		return listPaymentUser, err.Error
	}
	for index := 0; index < len(listPaymentUser); index++ {
		var listPaymentDetail []model.PaymentDetail
		err = userReceiver.sql.Db.Where("payment_id = ?", listPaymentUser[index].ID).Find(&listPaymentDetail)
		if err.Error != nil {
			logger.Error("Error get list payment url ", zap.Error(err.Error))
			continue
		}
		listPaymentUser[index].PaymentDetail = listPaymentDetail
	}
	return listPaymentUser, nil
}

func (userReceiver *UserRepoImpl) GetUserPaymentHistory(user model.User) ([]model.Payment, error) {
	var listPaymentUser []model.Payment
	err := userReceiver.sql.Db.Where("user_id = ?", user.ID).Preload("User").
		Preload("RatePlan").Preload("RoomType").Preload("Hotel").Find(&listPaymentUser)
	if err.Error != nil {
		logger.Error("Error get list cart url ", zap.Error(err.Error))
		return listPaymentUser, err.Error
	}
	for index := 0; index < len(listPaymentUser); index++ {
		var listPaymentDetail []model.PaymentDetail
		err = userReceiver.sql.Db.Where("payment_id = ?", listPaymentUser[index].ID).Find(&listPaymentDetail)
		if err.Error != nil {
			logger.Error("Error get list payment url ", zap.Error(err.Error))
			continue
		}
		listPaymentUser[index].PaymentDetail = listPaymentDetail
	}
	return listPaymentUser, nil
}

func (userReceiver *UserRepoImpl) UpdatePaymentStatus(payment model.Payment) (bool, error) {
	err := userReceiver.sql.Db.Model(&payment).Where("session_id=?", payment.SessionId).Updates(payment)
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return false, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return false, err.Error
		}
		return false, err.Error
	}
	return true, nil
}

func (userReceiver *UserRepoImpl) GetUserPayment(user model.User) ([]model.Payment, error) {
	var listPaymentUser []model.Payment
	err := userReceiver.sql.Db.Where("user_id = ? AND status = ?", user.ID, "pending").Preload("User").
		Preload("RatePlan").Preload("RoomType").Preload("Hotel").Find(&listPaymentUser)
	if err.Error != nil {
		logger.Error("Error get list cart url ", zap.Error(err.Error))
		return listPaymentUser, err.Error
	}
	for index := 0; index < len(listPaymentUser); index++ {
		var listPaymentDetail []model.PaymentDetail
		err = userReceiver.sql.Db.Where("payment_id = ?", listPaymentUser[index].ID).Find(&listPaymentDetail)
		if err.Error != nil {
			logger.Error("Error get list payment url ", zap.Error(err.Error))
			continue
		}
		listPaymentUser[index].PaymentDetail = listPaymentDetail
	}
	return listPaymentUser, nil
}

func (userReceiver *UserRepoImpl) CreatePaymentFromCart(user model.User) (bool, error) {
	err := userReceiver.sql.Db.Exec("call sp_addpayment(?);",
		user.ID)
	if err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

func (userReceiver *UserRepoImpl) GetUserCart(user model.User) ([]model.Cart, error) {
	var listCartUser []model.Cart
	err := userReceiver.sql.Db.Exec("call sp_clearcart(?);", user.ID)
	if err.Error != nil {
		return listCartUser, err.Error
	}
	err = userReceiver.sql.Db.Where("user_id = ?", user.ID).Preload("RatePlan").Preload("RoomType").Preload("Hotel").Find(&listCartUser)
	if err.Error != nil {
		logger.Error("Error get list cart url ", zap.Error(err.Error))
		return listCartUser, err.Error
	}
	for index := 0; index < len(listCartUser); index++ {
		var listCartItemDetail []model.CartDetail
		err = userReceiver.sql.Db.Where("cart_id = ?", listCartUser[index].ID).Find(&listCartItemDetail)
		if err.Error != nil {
			logger.Error("Error get list cart url ", zap.Error(err.Error))
			continue
		}
		listCartUser[index].Details = listCartItemDetail
	}
	return listCartUser, nil
}

func (userReceiver *UserRepoImpl) DeleteCart(cartId string) (bool, error) {
	var cartDelete model.Cart
	err := userReceiver.sql.Db.Where("id=?", cartId).Delete(cartDelete)
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return false, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return false, err.Error
		}
		return false, err.Error
	}
	return true, nil
}

func (userReceiver *UserRepoImpl) AddToCart(requestAddCart req.RequestAddToCart) (bool, error) {
	for i := 0; i < requestAddCart.NumberOfRooms; i++ {
		cartId, _ := utils.GetNewId()
		dateBeginAt, _ := time.Parse("2006-01-02 15:04:05", requestAddCart.FromDate)
		dateEndAt, _ := time.Parse("2006-01-02 15:04:05", requestAddCart.ToDate)

		cart := model.AddCart{
			ID:          cartId,
			StartAt:     dateBeginAt,
			EndAt:       dateEndAt,
			AdultNum:    requestAddCart.NumberOfAdults,
			ChildrenNum: requestAddCart.NumberOfChildren,
			RoomTypeId:  requestAddCart.RoomTypeID,
			RatePlanId:  requestAddCart.RatePlanID,
			HotelId:     requestAddCart.HotelID,
			UserId:      requestAddCart.UserId,
		}
		err := userReceiver.sql.Db.Exec("call sp_addtocart(?,?,?,?,?,?,?,?,?);",
			cart.ID,
			cart.StartAt,
			cart.EndAt,
			cart.AdultNum,
			cart.ChildrenNum,
			cart.RatePlanId,
			cart.RoomTypeId,
			cart.HotelId,
			cart.UserId)
		if err.Error != nil {
			return false, err.Error
		}
	}

	return true, nil
}

func (userReceiver *UserRepoImpl) GetUserNotifications(queryModel query.DataQueryModel) ([]model.Notification, error) {
	var listNotifications []model.Notification
	err := GenerateQueryGetData(userReceiver.sql, queryModel, &model.Notification{}, queryModel.ListIgnoreColumns)
	err = err.Find(&listNotifications)
	if err.Error != nil {
		logger.Error("Error get list notifications url ", zap.Error(err.Error))
		return listNotifications, err.Error
	}
	return listNotifications, nil
}

//func (userReceiver *UserRepoImpl) CheckProfileCustomerExistByIdentify(user model.Customer) (model.Customer, error) {
//	err := userReceiver.sql.Db.Where("identifier_number=?", user.IdentifierNumber).Find(&user)
//	if err != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return user, custom_error.UserNotFound
//		}
//		return user, err.Error
//	}
//	return user, nil
//}

//func (userReceiver *UserRepoImpl) SaveCustomerProfile(user model.Customer) (model.Customer, error) {
//	result := userReceiver.sql.Db.Create(&user)
//	if result.Error != nil {
//		return user, result.Error
//	}
//	return user, nil
//}

//func (userReceiver *UserRepoImpl) GetUserCart(customer model.User) (query.RoomAvailableQuery, error) {
//	var cartData = model.User{}
//	err := userReceiver.sql.Db.Where("id=?", customer.ID).Find(&cartData)
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

func (userReceiver *UserRepoImpl) GetProfileCustomer(customer model.User) (model.User, error) {
	var user = model.User{}
	var userRank = model.UserRank{}
	err := userReceiver.sql.Db.Where("id=?", customer.ID).Find(&user)
	err = userReceiver.sql.Db.Where("user_id=?", customer.ID).Find(&userRank)
	user.UserRank = &userRank
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

func (userReceiver *UserRepoImpl) UpdateRankCustomer(userRank model.UserRank) (model.UserRank, error) {
	err := userReceiver.sql.Db.Model(&userRank).Where("user_id=?", userRank.UserId).Updates(userRank)
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

func (userReceiver *UserRepoImpl) UpdateProfileCustomer(user model.User) (model.User, error) {
	err := userReceiver.sql.Db.Model(&user).Updates(user)
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

func (userReceiver *UserRepoImpl) GetUserRank(customer model.User) (model.UserRank, error) {
	var userRank = model.UserRank{}
	err := userReceiver.sql.Db.Where("id=?", customer.ID).Find(&userRank)
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
