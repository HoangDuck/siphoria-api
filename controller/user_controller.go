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
	"hotel-booking-api/model/res"
	"hotel-booking-api/repository"
	"hotel-booking-api/security"
	"hotel-booking-api/services"
	"hotel-booking-api/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	UserRepo    repository.UserRepo
	PaymentRepo repository.PaymentRepo
}

// HandleGetNotifications godoc
// @Summary Get notification
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/notifications [get]
func (userReceiver *UserController) HandleGetNotifications(c echo.Context) error {
	var listNotifications []model.Notification
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"user", "created_at", "updated_at", "",
	}, &model.Notification{})
	dataQueryModel.Role = claims.Role
	listNotifications, err := userReceiver.UserRepo.GetUserNotifications(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listNotifications)
	}
	return response.Ok(c, "Lấy danh sách thông báo thành công", struct {
		Data   []model.Notification `json:"data"`
		Paging res.PagingModel      `json:"paging"`
	}{
		Data: listNotifications,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
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
	if !(security.CheckRole(claims, model.CUSTOMER, false)) {
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
		Avatar: urls[0],
	}
	customer, err := userReceiver.UserRepo.UpdateProfileCustomer(customer)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật avatar thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", []string{
		urls[0],
	})
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
// @Router /users/rank-available [get]
func (userReceiver *UserController) HandleGetUserRank(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.CUSTOMER, false)) {
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
// @Param data body req.RequestAddToCart true "user"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/add-to-cart [post]
func (userReceiver *UserController) HandleAddToCart(c echo.Context) error {
	reqAddToCart := req.RequestAddToCart{}
	//binding
	if err := c.Bind(&reqAddToCart); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.CUSTOMER.String()) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	reqAddToCart.UserId = claims.UserId
	result, err := userReceiver.UserRepo.AddToCart(reqAddToCart)
	if err != nil || !result {
		return response.InternalServerError(c, "Thêm giỏ hàng thất bại", nil)
	}
	return response.Ok(c, "Thêm giỏ hàng thành công", result)
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
	user := model.User{
		ID: claims.UserId,
	}
	listCartUser, err := userReceiver.UserRepo.GetUserCart(user)
	if err != nil {
		logger.Error("Error query data", zap.Error(err))
		return response.InternalServerError(c, "Lấy danh sách giỏ hàng thành công", nil)
	}
	return response.Ok(c, "Lấy danh sách giỏ hàng", listCartUser)
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
	customer := model.User{
		ID: claims.UserId,
	}
	customerRankResult, err := userReceiver.UserRepo.GetUserRank(customer)
	customerResult, err := userReceiver.UserRepo.GetProfileCustomer(customer)
	customerResult.UserRank = &customerRankResult
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

// HandleDeleteCart godoc
// @Summary Delete user cart
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/carts/:id [delete]
func (userReceiver *UserController) HandleDeleteCart(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.CUSTOMER.String()) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	result, err := userReceiver.UserRepo.DeleteCart(c.Param("id"))
	if err != nil || !result {
		return response.InternalServerError(c, "Xoá giỏ hàng thất bại", nil)
	}
	return response.Ok(c, "Xoá giỏ hàng thành công", result)
}

// HandleCreatePaymentFromCart godoc
// @Summary Create payment
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/payments [post]
func (userReceiver *UserController) HandleCreatePaymentFromCart(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	customer := model.User{
		ID: claims.UserId,
	}
	_, err := userReceiver.PaymentRepo.CancelSessionPayment(claims.UserId)
	if err != nil {
		return response.InternalServerError(c, "Tạo thanh toán thất bại", nil)
	}
	customerResult, err := userReceiver.UserRepo.CreatePaymentFromCart(customer)
	if err != nil {
		logger.Error("Error get profile data", zap.Error(err))
		return response.InternalServerError(c, "Tải dữ liệu thất bại", nil)
	}

	return response.Ok(c, "Tạo thanh toán thành công", customerResult)
}

// HandleGetPayments godoc
// @Summary Get payment by user
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/payments [get]
func (userReceiver *UserController) HandleGetPayments(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.CUSTOMER.String()) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	user := model.User{
		ID: claims.UserId,
	}
	listPaymentUser, err := userReceiver.UserRepo.GetUserPayment(user)
	if err != nil {
		logger.Error("Error query data", zap.Error(err))
		return response.InternalServerError(c, "Lấy danh sách thanh toán thành công", nil)
	}
	return response.Ok(c, "Lấy danh sách thanh toán thành công", listPaymentUser)
}

// HandleUpdateStatusPayment godoc
// @Summary Update status payment
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/payments [put]
func (userReceiver *UserController) HandleUpdateStatusPayment(c echo.Context) error {
	reqUpdatePaymentStatus := req.RequestUpdatePaymentStatus{}
	if err := c.Bind(&reqUpdatePaymentStatus); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return err
	}
	err := c.Validate(reqUpdatePaymentStatus)
	if err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, "Thông tin không hợp lệ", nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	payment := model.Payment{
		UserId:    claims.UserId,
		SessionId: reqUpdatePaymentStatus.SessionId,
		Status:    "paid",
	}
	customerResult, err := userReceiver.UserRepo.UpdatePaymentStatus(payment)
	if err != nil {
		logger.Error("Error get profile data", zap.Error(err))
		return response.InternalServerError(c, "Tải dữ liệu thất bại", nil)
	}
	return response.Ok(c, "Cập nhật trạng thái thanh toán thành công", customerResult)
}

// HandleGetPaymentsHistory godoc
// @Summary Get payment by user
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/payments/history [get]
func (userReceiver *UserController) HandleGetPaymentsHistory(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)

	if !(claims.Role == model.CUSTOMER.String()) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	user := model.User{
		ID: claims.UserId,
	}
	listPaymentUser, err := userReceiver.UserRepo.GetUserPaymentHistory(c, user)
	if err != nil {
		logger.Error("Error query data", zap.Error(err))
		return response.InternalServerError(c, "Lấy danh sách thanh toán thành công", nil)
	}
	return response.Ok(c, "Lấy danh sách thanh toán", listPaymentUser)
}

// HandleGetPaymentsPendingCheckin godoc
// @Summary Get payment by user
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/payments/pending-checkin [get]
func (userReceiver *UserController) HandleGetPaymentsPendingCheckin(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.CUSTOMER.String()) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	user := model.User{
		ID: claims.UserId,
	}
	listPaymentUser, err := userReceiver.UserRepo.GetUserPaymentPendingCheckin(user)
	if err != nil {
		logger.Error("Error query data", zap.Error(err))
		return response.InternalServerError(c, "Lấy danh sách thanh toán thành công", nil)
	}
	return c.JSON(http.StatusOK, listPaymentUser)
}

// HandleSaveReview godoc
// @Summary Save review
// @Tags review-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestAddRatePlan true "review"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /users/reviews [post]
func (userReceiver *UserController) HandleSaveReview(c echo.Context) error {
	reqAddReview := req.RequestAddReview{}
	if err := c.Bind(&reqAddReview); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return err
	}
	err := c.Validate(reqAddReview)
	if err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, "Thông tin không hợp lệ", nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	reviewId, err := utils.GetNewId()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	review := model.Review{
		ID:        reviewId,
		UserId:    claims.UserId,
		HotelId:   reqAddReview.HotelId,
		Content:   reqAddReview.Content,
		Rating:    reqAddReview.Rating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}
	result, err := userReceiver.UserRepo.SaveReview(review)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Lưu thành công", result)
}

// HandleUpdateReview godoc
// @Summary Update review
// @Tags review-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateReview true "review"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /users/reviews/:id [patch]
func (userReceiver *UserController) HandleUpdateReview(c echo.Context) error {
	reqUpdateReview := req.RequestUpdateReview{}
	if err := c.Bind(&reqUpdateReview); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return err
	}
	err := c.Validate(reqUpdateReview)
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
	review := model.Review{
		ID:      c.Param("id"),
		UserId:  claims.UserId,
		Content: reqUpdateReview.Content,
		Rating:  reqUpdateReview.Rating,
	}
	review, err = userReceiver.UserRepo.UpdateReview(review)
	if err != nil {
		logger.Error("Error query data", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", review)
}

// HandleDeleteReview godoc
// @Summary Delete review
// @Tags review-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /users/reviews/:id [delete]
func (userReceiver *UserController) HandleDeleteReview(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.CUSTOMER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	reviewId := c.Param("id")
	review := model.Review{
		ID:        reviewId,
		IsDeleted: true,
	}
	result, err := userReceiver.UserRepo.DeleteReview(review)
	if !result {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Xoá thành công", nil)
}

// HandleCreatePayment godoc
// @Summary Create payment momo
// @description Choose payment method by add query param payment_method (?payment_method=momo,?payment_method=vnpay)
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestCreatePaymentModel true "payment"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /users/pay [post]
func (userReceiver *UserController) HandleCreatePayment(c echo.Context) error {
	reqCreatePayment := req.RequestCreatePaymentModel{}
	if err := c.Bind(&reqCreatePayment); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return err
	}
	paymentMethod := strings.ToLower(reqCreatePayment.PaymentMethod)

	listPaymentSessionId, err := userReceiver.PaymentRepo.GetPaymentListByCondition(reqCreatePayment.SessionID)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	var totalPrice float32
	for i := 0; i < len(listPaymentSessionId); i++ {
		totalPrice += listPaymentSessionId[i].TotalPrice
	}
	if paymentMethod == "momo" {
		momoService := services.GetMomoServiceInstance()
		//momoUrl := "https://momo.vn"
		momoUrl, err := userReceiver.PaymentRepo.GetMomoHostingUrl()
		if err != nil {
			return response.InternalServerError(c, err.Error(), nil)
		}
		logger.Info(momoUrl)
		//redirectMomoUrl := "https://momo.vn"
		redirectMomoUrl, err := userReceiver.PaymentRepo.GetRedirectMomoUrl()
		if err != nil {
			return response.InternalServerError(c, err.Error(), nil)
		}
		logger.Info(redirectMomoUrl)
		condition := map[string]interface{}{
			"booking-info":        "MOMO",
			"amount":              totalPrice,
			"booking-description": "Payment room Siphoria",
			"ipn-url":             momoUrl,
			"redirect-url":        redirectMomoUrl,
			"payment_id":          reqCreatePayment.SessionID,
		}
		dataResponse := momoService.PaymentService(condition)
		var tempResultCode = fmt.Sprint(dataResponse["resultCode"])
		if tempResultCode == "0" {
			_, err := userReceiver.PaymentRepo.UpdatePaymentMethodForPending(reqCreatePayment.SessionID, "Momo")
			if err != nil {
				return response.InternalServerError(c, "Tạo thanh toán thất bại", err.Error())
			}
		} else if tempResultCode == "41" {
			logger.Error("Error update momo payment " + tempResultCode)
			return response.InternalServerError(c, "Tạo thanh toán thất bại", dataResponse)
		} else {
			logger.Error("Error update momo payment " + tempResultCode)
			return response.InternalServerError(c, "Tạo thanh toán thất bại", dataResponse)
		}
		if err != nil {
			return response.BadRequest(c, err.Error(), nil)
		}
		return response.Ok(c, "Tạo thanh toán thành công", dataResponse)
	} else if paymentMethod == "vnpay" {
		vnpayService := services.GetVNPayServiceInstance()
		//momoUrl := "https://momo.vn"
		vnpayUrl, err := userReceiver.PaymentRepo.GetVNPayHostingUrl()
		if err != nil {
			return response.InternalServerError(c, err.Error(), nil)
		}
		if err != nil {
			return response.InternalServerError(c, err.Error(), nil)
		}
		condition := map[string]interface{}{
			"booking-info":        "VNPay",
			"amount":              int(totalPrice) * 100,
			"booking-description": "paymentsiphoria",
			"ipn-url":             vnpayUrl,
			"redirect-url":        "",
			"payment_id":          reqCreatePayment.SessionID + "_" + strconv.FormatInt(time.Now().Unix(), 10),
		}
		dataResponse := vnpayService.VNPayPaymentService(condition)
		if err != nil {
			return response.BadRequest(c, err.Error(), nil)
		}
		_, err = userReceiver.PaymentRepo.UpdatePaymentMethodForPending(reqCreatePayment.SessionID, "VNPay")
		if err != nil {
			return response.InternalServerError(c, "Tạo thanh toán thất bại", err.Error())
		}
		requestId, _ := utils.GetNewId()
		return response.Ok(c, "Tạo thanh toán thành công", res.DataPaymentRes{
			Amount:       int(totalPrice),
			Message:      "Tạo thanh toán thành công",
			OrderID:      "VNPay" + "_" + fmt.Sprint(condition["payment_id"]),
			PartnerCode:  "",
			PayURL:       services.ConfigInfo.VNPay.VNPUrl + "?" + dataResponse,
			RequestID:    requestId,
			ResponseTime: time.Now().Unix(),
			ResultCode:   0,
		})
	}
	return response.BadRequest(c, "Phương thức thanh toán chưa được hỗ trợ", nil)
}
