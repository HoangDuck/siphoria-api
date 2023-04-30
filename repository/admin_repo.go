package repository

import "hotel-booking-api/model"

type AdminRepo interface {
	CheckEmail(email string) (model.User, error)
	SaveAccount(account model.User) (model.User, error)
	UpdateAccount(accountStaff model.User) (model.User, error)
	GetAccountFilter() ([]model.User, error)
	GetHotelFilter() ([]model.Hotel, error)
	AcceptHotel(hotel model.Hotel) (model.Hotel, error)
	ApprovePayoutRequestHotel(hotelPayoutRequest model.PayoutRequest) (model.PayoutRequest, error)
}
