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
	"io"
	_ "math/rand"
	"strings"
)

type RoomController struct {
	RoomRepo repository.RoomRepo
}

// HandleSaveRoomType godoc
// @Summary Save room type
// @Tags room-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestCreateRoomType true "room"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms [post]
func (roomReceiver *RoomController) HandleSaveRoomType(c echo.Context) error {
	reqAddRoomType := req.RequestCreateRoomType{}
	//binding
	if err := c.Bind(&reqAddRoomType); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	roomTypeId, err := utils.GetNewId()
	if err != nil {
		return response.InternalServerError(c, "Đăng ký thất bại", nil)
	}
	reqAddRoomType.ID = roomTypeId
	result, err := roomReceiver.RoomRepo.SaveRoomType(reqAddRoomType)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Lưu room type thành công", result)
}

// HandleUpdateRoomNight godoc
// @Summary Update room nights
// @Tags room-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateRoomNight true "room"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms/roomnights [post]
func (roomReceiver *RoomController) HandleUpdateRoomNight(c echo.Context) error {
	reqUpdateRoomNight := req.RequestUpdateRoomNight{}
	//binding
	if err := c.Bind(&reqUpdateRoomNight); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	result, err := roomReceiver.RoomRepo.UpdateRoomNight(reqUpdateRoomNight)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật room nights thành công", result)
}

// HandleUpdateRatePackages godoc
// @Summary Update ratepackages
// @Tags room-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateRatePackage true "room"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms/ratepackages [post]
func (roomReceiver *RoomController) HandleUpdateRatePackages(c echo.Context) error {
	reqUpdateRatePackages := req.RequestUpdateRatePackage{}
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	reqUpdateRatePackages, err = req.UnmarshalRequestUpdateRatePackage(body)
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	result, err := roomReceiver.RoomRepo.UpdateRatePackages(reqUpdateRatePackages)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật rate packages thành công", result)
}

// HandleUpdateRoomType godoc
// @Summary update room type
// @Tags room-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateRoomType true "RoomType"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms/:id [patch]
func (roomReceiver *RoomController) HandleUpdateRoomType(c echo.Context) error {
	reqUpdateRoomType := req.RequestUpdateRoomType{}
	if err := c.Bind(&reqUpdateRoomType); err != nil {
		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) || security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	roomType, err := roomReceiver.RoomRepo.UpdateRoomType(reqUpdateRoomType, c.Param("id"))
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật thông tin room type thành công", roomType)
}

// HandleUpdateRoomPhotos godoc
// @Summary Update room photos
// @Tags room-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /rooms/:id/photos [patch]
func (roomReceiver *RoomController) HandleUpdateRoomPhotos(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) || security.CheckRole(claims, model.MANAGER, false)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	form, err := c.MultipartForm()
	if err != nil {
	}
	oldUrls := utils.DecodeJSONArray(form.Value["text"][0])
	urls := services.UploadMultipleFiles(c)
	if len(urls) == 0 {
		logger.Error("Error upload avatar to cloudinary failed", zap.Error(nil))
		return response.InternalServerError(c, "Cập nhật hình ảnh thất bại", nil)
	}
	urls = append(urls, oldUrls...)
	//find customer id by userid(account id)
	room := model.RoomType{
		ID:     c.Param("id"),
		Photos: strings.Join(urls, ";"),
	}
	room, err = roomReceiver.RoomRepo.UpdateRoomPhotos(room)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật hình ảnh thất bại", nil)
	}
	return response.Ok(c, "Cập nhật hình ảnh thành công", room)
}
