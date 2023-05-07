package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
	"hotel-booking-api/security"
	"hotel-booking-api/services"
	"hotel-booking-api/utils"
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
func (hotelController *HotelController) HandleSearchHotel(c echo.Context) error {
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
func (hotelController *HotelController) HandleGetHotelById(c echo.Context) error {
	return response.Ok(c, "Cập nhật thành công", nil)
}

// HandleGetHotelPartner godoc
// @Summary Get hotel Controller
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels [get]
func (hotelController *HotelController) HandleGetHotelPartner(c echo.Context) error {
	var listHotel []model.Hotel
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) ||
		security.CheckRole(claims, model.STAFF, false) ||
		security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"hotelier", "created_at", "updated_at", "",
	}, &model.Hotel{})
	dataQueryModel.UserId = claims.UserId
	listHotel, err := hotelController.HotelRepo.GetHotelFilter(dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listHotel)
	}
	return response.Ok(c, "Lấy danh sách khách sạn thành công", listHotel)
}

// HandleCreateHotel godoc
// @Summary Create hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestCreateHotel true "hotel"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /hotels [post]
func (hotelController *HotelController) HandleCreateHotel(c echo.Context) error {
	reqCreateHotel := req.RequestCreateHotel{}
	//binding
	if err := c.Bind(&reqCreateHotel); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	hotelId, err := utils.GetNewId()
	if err != nil {
		return response.Forbidden(c, "Đăng ký thất bại", nil)
	}
	reqCreateHotel.ID = hotelId
	reqCreateHotel.HotelierID = claims.UserId
	result, err := hotelController.HotelRepo.SaveHotel(reqCreateHotel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Đăng ký khách sạn thành công", result)
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
func (hotelController *HotelController) HandleUpdateHotelPhoto(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
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
	hotel, err := hotelController.HotelRepo.UpdateHotelPhotos(hotel)
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
func (hotelController *HotelController) HandleUpdateHotelBusinessLicense(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
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
	hotel, err := hotelController.HotelRepo.UpdateHotelBusinessLicensePhotos(hotel)
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
func (hotelController *HotelController) HandleDeleteHotelBusinessLicense(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
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
	hotel, err := hotelController.HotelRepo.UpdateHotelPhotos(hotel)
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
func (hotelController *HotelController) HandleSendRequestPaymentHotel(c echo.Context) error {
	reqCreatePayout := req.RequestCreatePayout{}
	//binding
	if err := c.Bind(&reqCreatePayout); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	payoutRequestId, err := utils.GetNewId()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	payoutRequest := model.PayoutRequest{
		ID:          payoutRequestId,
		HotelId:     c.Param("hotel_id"),
		PettionerId: claims.UserId,
	}
	payoutRequestResult, err := hotelController.HotelRepo.CreateRequestPayout(payoutRequest, reqCreatePayout.Payments)

	if err != nil {
		logger.Error("Error uuid data", zap.Error(err))
		return response.InternalServerError(c, "Tạo yêu cầu thanh toán thất bại", nil)
	}
	return response.Ok(c, "Tạo yêu cầu thanh toán thành công", payoutRequestResult)
}

// HandleUpdateHotel godoc
// @Summary update hotel info
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateHotel true "hotel"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /hotel/:id [patch]
func (hotelController *HotelController) HandleUpdateHotel(c echo.Context) error {
	reqUpdateHotel := req.RequestUpdateHotel{}
	if err := c.Bind(&reqUpdateHotel); err != nil {
		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String() || claims.Role == model.HOTELIER.String() || claims.Role == model.SUPERADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	hotel, err := hotelController.HotelRepo.UpdateHotel(reqUpdateHotel, c.Param("id"))
	if err != nil {
		return response.UnprocessableEntity(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật thông tin khách sạn thành công", hotel)
}
