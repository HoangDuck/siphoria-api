package controller

import (
	"github.com/labstack/echo/v4"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/repository"
)

type HotelController struct {
	HotelRepo repository.HotelRepo
}

// HandleSearchHotel godoc
// @Summary Search Hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateProfile true "hotel"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/search [post]
func (hotelReceiver *HotelController) HandleSearchHotel(c echo.Context) error {
	return response.Ok(c, "Cập nhật thành công", nil)
}

// HandleGetHotelById godoc
// @Summary Get hotel by Id
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id [get]
func (hotelReceiver *HotelController) HandleGetHotelById(c echo.Context) error {
	return response.Ok(c, "Cập nhật thành công", nil)
}
