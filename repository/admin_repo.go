package repository

import (
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
)

type AdminRepo interface {
	CheckEmail(email string) (model.User, error)
	SaveAccount(account model.User) (model.User, error)
	UpdateAccount(accountStaff model.User) (model.User, error)
	GetAccountFilter(queryModel query.DataQueryModel) ([]model.User, error)
	GetHotelFilter(queryModel query.DataQueryModel) ([]model.Hotel, error)
	AcceptHotel(hotel model.Hotel) (model.Hotel, error)
	ApprovePayoutRequestHotel(hotelPayoutRequest model.PayoutRequest) (model.PayoutRequest, error)
}
