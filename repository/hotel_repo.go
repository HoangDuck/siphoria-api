package repository

import (
	"github.com/labstack/echo/v4"
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
	"hotel-booking-api/model/req"
)

type HotelRepo interface {
	UpdateHotelPhotos(hotel model.Hotel) (model.Hotel, error)
	UpdateHotel(requestUpdateHotel req.RequestUpdateHotel, idHotel string) (model.Hotel, error)
	UpdateHotelBusinessLicensePhotos(hotel model.Hotel) (model.Hotel, error)
	CreateRequestPayout(payoutRequest model.PayoutRequest, paymentIds []string) (model.PayoutRequest, error)
	SaveHotel(requestAddHotel req.RequestCreateHotel) (model.Hotel, error)
	GetHotelFilter(queryModel *query.DataQueryModel) ([]model.Hotel, error)
	GetRoomTypeFilter(queryModel *query.DataQueryModel) ([]model.RoomType, error)
	GetHotelMobile() ([]model.Hotel, error)
	GetPayoutRequestByHotel(queryModel *query.DataQueryModel) ([]model.PayoutRequest, error)
	GetListHotelSearch(context echo.Context) ([]model.HotelSearch, error)
	GetReviewsByHotel(queryModel *query.DataQueryModel) ([]model.Review, error)
	GetHotelById(context echo.Context) (model.Hotel, error)
}
