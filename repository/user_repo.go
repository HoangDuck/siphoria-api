package repository

import (
	"hotel-booking-api/model"
)

type UserRepo interface {
	//SaveCustomerProfile(user model.User) (model.User, error)
	UpdateProfileCustomer(user model.User) (model.User, error)
	UpdateRankCustomer(userRank model.UserRank) (model.UserRank, error)
	GetProfileCustomer(user model.User) (model.User, error)
	GetUserRank(user model.User) (model.UserRank, error)
	//GetUserCart(user model.User) (query.RoomAvailableQuery, error)
	//CheckProfileCustomerExistByIdentify(user model.User) (model.User, error)
}
