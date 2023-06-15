package controller

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/model/req"
	"hotel-booking-api/model/res"
	"hotel-booking-api/repository"
	"hotel-booking-api/security"
	"hotel-booking-api/utils"
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
		FirstName: reqRegister.FirstName,
		LastName:  reqRegister.LastName,
		FullName:  reqRegister.FirstName + " " + reqRegister.LastName,
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
	if !(security.CheckRole(claims, model.HOTELIER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	hash := security.HashAndSalt([]byte(reqChangeAccount.Password))
	account := model.User{
		ID:        reqChangeAccount.ID,
		FirstName: reqChangeAccount.FirstName,
		LastName:  reqChangeAccount.LastName,
		FullName:  reqChangeAccount.FirstName + " " + reqChangeAccount.LastName,
		Password:  hash,
		Status:    reqChangeAccount.Status,
	}
	account, err = adminController.AdminRepo.UpdateAccount(account)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật tài khoản người dùng thành công", account)
}

// HandleGetHotelWorkByEmployee godoc
// @Summary Get list hotel employee works
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/works/:id [get]
func (adminController *AdminController) HandleGetHotelWorkByEmployee(c echo.Context) error {
	var listHotel []model.Hotel
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"hotelier", "created_at", "updated_at", "", "overview", "rating", "commission_rate", "status", "activate", "province_code", "district_code", "ward_core", "raw_address", "hotel_photos", "bank_account", "bank_beneficiary", "bank_name", "business_licence", "hotelier_id", "price_hotel", "discount_price", "discount_hotel", "hotel_type", "hotel_facility",
	}, &model.Hotel{})
	dataQueryModel.UserId = claims.UserId
	listHotel, err := adminController.AdminRepo.GetHotelWorkByEmployee(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Lấy danh sách khách sạn nhân viên làm việc thành công", struct {
		UserId  string        `json:"user_id"`
		HotelId []model.Hotel `json:"hotel_id"`
	}{
		UserId:  dataQueryModel.UserId,
		HotelId: listHotel,
	})
}

// HandleGetAccountByAdmin godoc
// @Summary Get account admin (Sort: pass sort=field&order=desc, Filter: field_you_want_to_pass=value, Paging: page=0&offset=3, Search: search=8)
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/accounts [get]
func (adminController *AdminController) HandleGetAccountByAdmin(c echo.Context) error {
	var listUser []model.User
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"token", "created_at", "updated_at", "", "avatar", "email", "phone", "gender", "gender", "role", "status", "-", "user_rank",
	}, &model.User{})
	listUser, err := adminController.AdminRepo.GetAccountFilter(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listUser)
	}
	return response.Ok(c, "Lấy danh sách tài khoản thành công", struct {
		Data   []model.User    `json:"data"`
		Paging res.PagingModel `json:"paging"`
	}{
		Data: listUser,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
}

// HandleGetHotelByAdmin godoc
// @Summary Get hotel list
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/hotels [get]
func (adminController *AdminController) HandleGetHotelByAdmin(c echo.Context) error {
	var listHotel []model.Hotel
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"hotelier", "created_at", "updated_at", "", "overview", "rating", "commission_rate", "status", "activate", "province_code", "district_code", "ward_core", "raw_address", "hotel_photos", "bank_account", "bank_beneficiary", "bank_name", "business_licence", "hotelier_id", "price_hotel", "discount_price", "discount_hotel", "hotel_type", "hotel_facility",
	}, &model.Hotel{})
	listHotel, err := adminController.AdminRepo.GetHotelFilter(&dataQueryModel)

	fmt.Println(listHotel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listHotel)
	}

	return response.Ok(c, "Lấy danh sách khách sạn thành công", struct {
		Data   []model.Hotel   `json:"data"`
		Paging res.PagingModel `json:"paging"`
	}{
		Data: listHotel,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
}

// HandleAcceptHotel godoc
// @Summary Accept
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/accept/:id [patch]
func (adminController *AdminController) HandleAcceptHotel(c echo.Context) error {
	var hotel model.Hotel
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	hotel = model.Hotel{
		ID:     c.Param("id"),
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
// @Router /admin/update-rating/:id [patch]
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
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	hotel = model.Hotel{
		ID:     c.Param("id"),
		Rating: reqUpdateRating.Rating,
	}

	hotel, err := adminController.AdminRepo.UpdateRatingHotel(hotel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), hotel)
	}
	return response.Ok(c, "Cập nhật số sao thành công", hotel)
}

// HandleUpdateCommissionRateHotel godoc
// @Summary Update commission rate hotel
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/update-cmsrate/:id [patch]
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
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	hotel = model.Hotel{
		ID:             c.Param("id"),
		CommissionRate: reqUpdateCommissionRating.CommissionRate,
	}

	hotel, err := adminController.AdminRepo.UpdateCommissionRatingHotel(hotel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), hotel)
	}
	return response.Ok(c, "Cập nhật tỷ lệ hoa hồng thành công", hotel)
}

// HandleApprovePayoutHotel godoc
// @Summary Approve request payout
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/payouts/:id [patch]
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
	if !(security.CheckRole(claims, model.ACCOUNTANT, false)) {
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

// HandleDeleteHotelWorkByEmployee godoc
// @Summary Delete HotelWork By Employee
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestDeleteHotelWorkByEmployee true "staff"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/works [delete]
func (adminController *AdminController) HandleDeleteHotelWorkByEmployee(c echo.Context) error {
	reqDeleteHotelWorkByEmployee := req.RequestDeleteHotelWorkByEmployee{}
	//binding
	if err := c.Bind(&reqDeleteHotelWorkByEmployee); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	result, err := adminController.AdminRepo.DeleteHotelWorkByEmployee(reqDeleteHotelWorkByEmployee)
	if !result {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Xoá nhân viên thành công", nil)
}

// HandleSaveHotelWorkByEmployee godoc
// @Summary  HotelWork By Employee
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestDeleteHotelWorkByEmployee true "staff"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/works [post]
func (adminController *AdminController) HandleSaveHotelWorkByEmployee(c echo.Context) error {
	reqSaveHotelWorkByEmployee := req.RequestSaveHotelWorkByEmployee{}
	//binding
	if err := c.Bind(&reqSaveHotelWorkByEmployee); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	hotelWork := model.HotelWork{
		HotelId:   reqSaveHotelWorkByEmployee.HotelId,
		UserId:    reqSaveHotelWorkByEmployee.UserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}

	result, err := adminController.AdminRepo.SaveHotelWorkByEmployee(hotelWork)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Xoá nhân viên thành công", result)
}

// HandleGetPayoutByAdmin godoc
// @Summary Get payout by admin
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/payouts [get]
func (adminController *AdminController) HandleGetPayoutByAdmin(c echo.Context) error {
	var listPayout []model.PayoutRequest
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ACCOUNTANT, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"pettioner", "hotel", "payer", "", "created_at", "updated_at", "-", "open_at", "close_at",
	}, &model.PayoutRequest{})
	listPayout, err := adminController.AdminRepo.GetPayoutRequest(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listPayout)
	}
	return response.Ok(c, "Lấy danh sách yêu cầu thanh toán thành công", struct {
		Data   []model.PayoutRequest `json:"data"`
		Paging res.PagingModel       `json:"paging"`
	}{
		Data: listPayout,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
}
