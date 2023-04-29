package repository

import "hotel-booking-api/model"

type HotelRepo interface {
	UpdateHotelPhotos(hotel model.Hotel) (model.Hotel, error)
	UpdateHotelBusinessLicensePhotos(hotel model.Hotel) (model.Hotel, error)
	CreateRequestPayout(payoutRequest model.PayoutRequest, paymentIds string) (model.PayoutRequest, error)
}
