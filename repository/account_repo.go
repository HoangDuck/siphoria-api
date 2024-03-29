package repository

import (
	"hotel-booking-api/model"
	"hotel-booking-api/model/req"
)

type AccountRepo interface {
	CheckLogin(request req.RequestSignIn) (model.User, error)
	GetAccountById(userId string) (model.User, error)
	SaveAccount(account model.User) (model.User, error)
	SaveAccountRank(accountRank model.UserRank) (model.UserRank, error)
	CheckEmailExisted(email string) (model.User, error)
	GetCustomerPageUrl() (string, error)
	GetCustomerActivatePageUrl() (string, error)
	ResetPassword(userEmail string, hashedPassword string) (bool, error)
	ActivateAccount(account model.User) (model.User, error)
	UpdatePassword(userId string, newPassword string) (bool, error)
}
