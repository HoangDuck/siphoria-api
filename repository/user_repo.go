package repository

import (
	"github.com/labstack/echo/v4"
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
	"hotel-booking-api/model/req"
)

type UserRepo interface {
	UpdateProfileCustomer(user model.User) (model.User, error)
	UpdateRankCustomer(userRank model.UserRank) (model.UserRank, error)
	GetProfileCustomer(user model.User) (model.User, error)
	GetUserRank(user model.User) (model.UserRank, error)
	GetUserNotifications(queryModel *query.DataQueryModel) ([]model.Notification, error)
	AddToCart(requestAddCart req.RequestAddToCart) (bool, error)
	DeleteCart(cartId string) (bool, error)
	GetUserCart(user model.User) ([]model.Cart, error)
	GetUserPayment(context echo.Context, user model.User) ([]model.Payment, error)
	CreatePaymentFromCart(user model.User) (string, error)
	UpdatePaymentStatus(payment model.Payment) (bool, error)
	GetUserPaymentHistory(context echo.Context, user model.User) ([]model.Payment, error)
	GetUserPaymentPendingCheckin(user model.User) ([]model.Payment, error)
	SaveReview(review model.Review) (model.Review, error)
	UpdateReview(review model.Review) (model.Review, error)
	DeleteReview(review model.Review) (bool, error)
	//CheckProfileCustomerExistByIdentify(user model.User) (model.User, error)
}
