package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
	"hotel-booking-api/security"
	"time"
)

type AdminController struct {
	AdminRepo repository.AdminRepo
}

// HandleCreateAccount godoc
// @Summary create account by admin
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestAddStatusBooking true "status booking"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/create-account [post]
func (adminController *AdminController) HandleCreateAccount(c echo.Context) error {
	reqRegister := req.RequestCreateAccountByAdmin{}
	//binding
	if err := c.Bind(&reqRegister); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	//Validate
	if err := c.Validate(reqRegister); err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.SUPERADMIN.String() || claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	//validate existed email
	_, err := adminController.AdminRepo.CheckEmail(reqRegister.Email)
	if err != nil {
		return response.Conflict(c, "Email đã tồn tại", nil)
	}
	//Generate UUID
	accountId, err := uuid.NewUUID()
	if err != nil {
		logger.Error("Error uuid data", zap.Error(err))
		return response.Forbidden(c, err.Error(), nil)
	}
	//create password
	hash := security.HashAndSalt([]byte(reqRegister.Password))
	//Init account
	account := model.User{
		ID:        accountId.String(),
		Email:     reqRegister.Email,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      reqRegister.Role,
		Status:    1,
	}

	//Save account
	account, err = adminController.AdminRepo.SaveAccount(account)
	if err != nil {
		return response.Conflict(c, err.Error(), nil)
	}
	return response.Ok(c, "Đăng ký thành công", nil)
}

// HandleUpdateAccount godoc
// @Summary update account
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateAccount true "staffaccount"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/update-account [patch]
func (adminController *AdminController) HandleUpdateAccount(c echo.Context) error {
	reqChangeAccount := req.RequestUpdateAccount{}
	if err := c.Bind(&reqChangeAccount); err != nil {
		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
	}
	err := c.Validate(reqChangeAccount)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String() || claims.Role == model.HOTELIER.String() || claims.Role == model.SUPERADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	account := model.User{
		ID:        reqChangeAccount.ID,
		FirstName: reqChangeAccount.FirstName,
		LastName:  reqChangeAccount.LastName,
		Password:  reqChangeAccount.Password,
		Status:    reqChangeAccount.Status,
	}
	account, err = adminController.AdminRepo.UpdateAccount(account)
	if err != nil {
		return response.UnprocessableEntity(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật cài đặt thành công", account)
}

// HandleGetAccountByAdmin godoc
// @Summary Get all account status list
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/accounts?s= [get]
func (adminController *AdminController) HandleGetAccountByAdmin(c echo.Context) error {
	var listUser []model.User
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String() || claims.Role == model.SUPERADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	listUser, err := adminController.AdminRepo.GetAccountFilter()
	if err != nil {
		return response.InternalServerError(c, err.Error(), listUser)
	}
	return response.Ok(c, "Lấy danh sách tài khoản thành công", listUser)
}

// HandleGetHotelByAdmin godoc
// @Summary Get all hotel list
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/hotels [post]
func (adminController *AdminController) HandleGetHotelByAdmin(c echo.Context) error {
	var listHotel []model.Hotel
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	listHotel, err := adminController.AdminRepo.GetHotelFilter()
	if err != nil {
		return response.InternalServerError(c, err.Error(), listHotel)
	}
	return response.Ok(c, "Lấy danh sách tài khoản thành công", listHotel)
}

// HandleAcceptHotel godoc
// @Summary Accept
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/accept/:hotel-id [patch]
func (adminController *AdminController) HandleAcceptHotel(c echo.Context) error {
	var hotel model.Hotel
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	hotel = model.Hotel{
		ID:     c.Param("hotel-id"),
		Status: 1,
	}

	hotel, err := adminController.AdminRepo.AcceptHotel(hotel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), hotel)
	}
	return response.Ok(c, "Duyệt khách sạn thành công", hotel)
}

// HandleUpdateRatingHotel godoc
// @Summary Update rating hotel
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/update-rating/:hotel-id [patch]
func (adminController *AdminController) HandleUpdateRatingHotel(c echo.Context) error {
	reqUpdateRating := req.RequestUpdateRating{}
	//binding
	if err := c.Bind(&reqUpdateRating); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	//Validate
	if err := c.Validate(reqUpdateRating); err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	var hotel model.Hotel
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	hotel = model.Hotel{
		ID:     c.Param("hotel-id"),
		Rating: reqUpdateRating.Rating,
	}

	hotel, err := adminController.AdminRepo.AcceptHotel(hotel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), hotel)
	}
	return response.Ok(c, "Cập nhật khách sạn thành công", hotel)
}

// HandleUpdateCommissionRateHotel godoc
// @Summary Update commission rate hotel
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/update-cmsrate/:hotel-id [patch]
func (adminController *AdminController) HandleUpdateCommissionRateHotel(c echo.Context) error {
	reqUpdateCommissionRating := req.RequestUpdateCommissionRating{}
	//binding
	if err := c.Bind(&reqUpdateCommissionRating); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	//Validate
	if err := c.Validate(reqUpdateCommissionRating); err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	var hotel model.Hotel
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	hotel = model.Hotel{
		ID:             c.Param("hotel-id"),
		CommissionRate: reqUpdateCommissionRating.CommissionRate,
	}

	hotel, err := adminController.AdminRepo.AcceptHotel(hotel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), hotel)
	}
	return response.Ok(c, "Cập nhật khách sạn thành công", hotel)
}

// HandleApprovePayoutHotel godoc
// @Summary Approve request payout
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/payouts/:payout-request-id [patch]
func (adminController *AdminController) HandleApprovePayoutHotel(c echo.Context) error {
	reqApprovePayout := req.RequestApprovePayout{}
	//binding
	if err := c.Bind(&reqApprovePayout); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	var hotelPayoutRequest model.PayoutRequest
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	hotelPayoutRequest = model.PayoutRequest{
		PayerId: claims.UserId,
		Resolve: reqApprovePayout.Resolve,
	}

	hotelPayoutRequest, err := adminController.AdminRepo.ApprovePayoutRequestHotel(hotelPayoutRequest)
	if err != nil {
		return response.InternalServerError(c, err.Error(), hotelPayoutRequest)
	}
	return response.Ok(c, "Duyệt yêu cầu thanh toán thành công", hotelPayoutRequest)
}
