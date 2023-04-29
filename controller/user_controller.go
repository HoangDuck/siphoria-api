package controller

import (
	"fmt"
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
	"strings"
)

type UserController struct {
	UserRepo repository.UserRepo
}

// HandleUpdateAvatar godoc
// @Summary Update user's avatar
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/update-avatar [post]
func (userReceiver *UserController) HandleUpdateAvatar(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.CUSTOMER)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	urls := services.UploadMultipleFiles(c)
	if len(urls) == 0 {
		logger.Error("Error upload avatar to cloudinary failed", zap.Error(nil))
		return response.InternalServerError(c, "Cập nhật avatar thất bại", nil)
	}
	//find customer id by userid(account id)
	customer := model.User{
		ID:     claims.UserId,
		Avatar: fmt.Sprintf("[%s]", strings.Join(urls, ",")),
	}
	customer, err := userReceiver.UserRepo.UpdateProfileCustomer(customer)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật avatar thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", nil)
}

// HandleGetUserRank godoc
// @Summary Get user rank
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateProfile true "user"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/get-rank-available [get]
func (userReceiver *UserController) HandleGetUserRank(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.CUSTOMER)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	customer := model.User{
		ID: claims.UserId,
	}
	customerResult, err := userReceiver.UserRepo.GetUserRank(customer)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Lấy thông tin thành công", customerResult)
}

// HandleAddToCart godoc
// @Summary Add to cart
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateProfile true "user"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/add-to-cart [post]
func (userReceiver *UserController) HandleAddToCart(c echo.Context) error {
	return response.Ok(c, "Cập nhật thành công", nil)
}

// HandleGetCart godoc
// @Summary Get cart by user
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/carts [get]
func (userReceiver *UserController) HandleGetCart(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.CUSTOMER.String()) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	//find customer id by userid(account id)
	//customer := model.User{
	//	ID: claims.UserId,
	//}
	//cartData, err := userReceiver.UserRepo.GetUserCart(customer)
	//if err != nil {
	//	logger.Error("Error query data", zap.Error(err))
	//	return response.InternalServerError(c, "Tải dữ liệu thất bại", nil)
	//}
	return response.Ok(c, "Danh sách giỏ hàng", nil /*cartData*/)
}

// HandleUpdateRank godoc
// @Summary Update user's rank
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateProfile true "user"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/update-rank [post]
func (userReceiver *UserController) HandleUpdateRank(c echo.Context) error {
	reqUpdateRank := req.RequestUpdateRank{}
	//binding
	if err := c.Bind(&reqUpdateRank); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.CUSTOMER.String()) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	userRank := model.UserRank{
		ID:     claims.UserId,
		RankId: reqUpdateRank.RankTo,
	}
	userRank, err := userReceiver.UserRepo.UpdateRankCustomer(userRank)
	if err != nil {
		logger.Error("Error query data", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", userRank)
}

// HandleUpdateProfile godoc
// @Summary Update customer profile
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateProfile true "user"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/update-profile [patch]
func (userReceiver *UserController) HandleUpdateProfile(c echo.Context) error {
	reqUpdateProfile := req.RequestUpdateProfile{}
	if err := c.Bind(&reqUpdateProfile); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return err
	}
	err := c.Validate(reqUpdateProfile)
	if err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, "Thông tin không hợp lệ", nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if claims.Role != model.CUSTOMER.String() {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	//find customer id by userid(account id)
	customer := model.User{
		ID:        claims.UserId,
		FirstName: reqUpdateProfile.FirstName,
		LastName:  reqUpdateProfile.LastName,
		Phone:     reqUpdateProfile.Phone,
	}
	customer, err = userReceiver.UserRepo.UpdateProfileCustomer(customer)
	if err != nil {
		logger.Error("Error query data", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", customer)
}

// HandleGetCustomerProfileInfo godoc
// @Summary Get Customer Profile
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/details [get]
func (userReceiver *UserController) HandleGetCustomerProfileInfo(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.CUSTOMER.String()) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	customer := model.User{
		ID: claims.UserId,
	}
	customerResult, err := userReceiver.UserRepo.GetProfileCustomer(customer)
	if err != nil {
		logger.Error("Error get profile data", zap.Error(err))
		return response.InternalServerError(c, "Tải dữ liệu thất bại", nil)
	}
	return response.Ok(c, "Lấy thông tin thành công", customerResult)
}

// HandleGetUserNotifications godoc
// @Summary Get User Notifications
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/notifications [get]
func (userReceiver *UserController) HandleGetUserNotifications(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	customer := model.User{
		ID: claims.UserId,
	}
	customerResult, err := userReceiver.UserRepo.GetProfileCustomer(customer)
	if err != nil {
		logger.Error("Error get profile data", zap.Error(err))
		return response.InternalServerError(c, "Tải dữ liệu thất bại", nil)
	}
	return response.Ok(c, "Lấy thông tin thành công", customerResult)
}
