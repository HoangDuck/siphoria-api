package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/repository"
	"hotel-booking-api/security"
	"hotel-booking-api/services"
	"strings"
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

// HandleCreateHotel godoc
// @Summary Create hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels [post]
func (hotelReceiver *HotelController) HandleCreateHotel(c echo.Context) error {
	return response.Ok(c, "Cập nhật thành công", nil)
}

// HandleUpdateHotelPhoto godoc
// @Summary Update hotel photos
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id/photos [post]
func (hotelReceiver *HotelController) HandleUpdateHotelPhoto(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	urls := services.UploadMultipleFiles(c)
	if len(urls) == 0 {
		logger.Error("Error upload avatar to cloudinary failed", zap.Error(nil))
		return response.InternalServerError(c, "Cập nhật hình ảnh thất bại", nil)
	}
	//find customer id by userid(account id)
	hotel := model.Hotel{
		ID:          c.Param("id"),
		HotelPhotos: strings.Join(urls, ""),
	}
	hotel, err := hotelReceiver.HotelRepo.UpdateHotelPhotos(hotel)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật avatar thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", hotel)
}

// HandleUpdateHotelBusinessLicense godoc
// @Summary Update hotel business license photos
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id/business-license [post]
func (hotelReceiver *HotelController) HandleUpdateHotelBusinessLicense(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	urls := services.UploadMultipleFiles(c)
	if len(urls) == 0 {
		logger.Error("Error upload avatar to cloudinary failed", zap.Error(nil))
		return response.InternalServerError(c, "Cập nhật hình ảnh thất bại", nil)
	}
	//find customer id by userid(account id)
	hotel := model.Hotel{
		ID:              c.Param("id"),
		BusinessLicence: strings.Join(urls, ""),
	}
	hotel, err := hotelReceiver.HotelRepo.UpdateHotelBusinessLicensePhotos(hotel)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật avatar thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", hotel)
}

// HandleDeleteHotelBusinessLicense godoc
// @Summary Delete hotel photos
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id/photos [delete]
func (hotelReceiver *HotelController) HandleDeleteHotelBusinessLicense(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	urls := services.UploadMultipleFiles(c)
	if len(urls) == 0 {
		logger.Error("Error upload avatar to cloudinary failed", zap.Error(nil))
		return response.InternalServerError(c, "Cập nhật hình ảnh thất bại", nil)
	}
	//find customer id by userid(account id)
	hotel := model.Hotel{
		ID:          c.Param("id"),
		HotelPhotos: "",
	}
	hotel, err := hotelReceiver.HotelRepo.UpdateHotelPhotos(hotel)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật avatar thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", hotel)
}

// HandleSendRequestPaymentHotel godoc
// @Summary Send request payment hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:hotel_id/payout [post]
func (hotelReceiver *HotelController) HandleSendRequestPaymentHotel(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	return response.Ok(c, "Tạo yêu cầu thanh toán thành công", nil)
}
