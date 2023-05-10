package repository

import (
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
)

type UserRepo interface {
	UpdateProfileCustomer(user model.User) (model.User, error)
	UpdateRankCustomer(userRank model.UserRank) (model.UserRank, error)
	GetProfileCustomer(user model.User) (model.User, error)
	GetUserRank(user model.User) (model.UserRank, error)
	GetUserNotifications(queryModel query.DataQueryModel) ([]model.Notification, error)
	//GetUserCart(user model.User) (query.RoomAvailableQuery, error)
	//CheckProfileCustomerExistByIdentify(user model.User) (model.User, error)
}
