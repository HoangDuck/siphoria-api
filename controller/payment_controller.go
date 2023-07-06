package controller

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/model/res"
	"hotel-booking-api/repository"
	"hotel-booking-api/services"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaymentController struct {
	PaymentRepo repository.PaymentRepo
}

// CreatePaymentWithVNPay godoc
// @Summary Create payment vnpay
// @Tags payment-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestCreatePayment true "payment"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /payment/create-vnpay [get]
func (paymentReceiver *PaymentController) CreatePaymentWithVNPay(c echo.Context) error {
	vnpayService := services.GetVNPayServiceInstance()
	//momoUrl := "https://momo.vn"
	vnpayUrl, err := paymentReceiver.PaymentRepo.GetVNPayHostingUrl()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	//redirectMomoUrl := "https://momo.vn"
	redirectMomoUrl, err := paymentReceiver.PaymentRepo.GetRedirectMomoUrl()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	paymentId, err := uuid.NewUUID()
	condition := map[string]interface{}{
		"booking-info":        "VNPay",
		"amount":              50000,
		"booking-description": "asdas",
		"ipn-url":             vnpayUrl,
		"redirect-url":        redirectMomoUrl,
		"payment_id":          paymentId.String(),
	}
	dataResponse := vnpayService.VNPayPaymentService(condition)
	if err != nil {
		return response.BadRequest(c, "nil", nil)
	}
	return c.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "Tạo thanh toán thành công",
		Data:       services.ConfigInfo.VNPay.VNPUrl + "?" + dataResponse,
	})
}

// CreatePaymentWithMomo godoc
// @Summary Create payment momo
// @Tags payment-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestCreatePayment true "payment"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /payment/create-momo [get]
func (paymentReceiver *PaymentController) CreatePaymentWithMomo(c echo.Context) error {
	momoService := services.GetMomoServiceInstance()
	//momoUrl := "https://momo.vn"
	momoUrl, err := paymentReceiver.PaymentRepo.GetMomoHostingUrl()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	logger.Info(momoUrl)
	//redirectMomoUrl := "https://momo.vn"
	redirectMomoUrl, err := paymentReceiver.PaymentRepo.GetRedirectMomoUrl()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	logger.Info(redirectMomoUrl)
	paymentId, err := uuid.NewUUID()
	condition := map[string]interface{}{
		"booking-info":        "MOMO",
		"amount":              50000,
		"booking-description": "asdasfsdgfsdgsd",
		"ipn-url":             momoUrl,
		"redirect-url":        redirectMomoUrl,
		"payment_id":          paymentId.String() + "_" + strconv.FormatInt(time.Now().Unix(), 10),
	}
	dataResponse := momoService.PaymentService(condition)
	var tempResultCode = fmt.Sprint(dataResponse["resultCode"])
	if tempResultCode == "0" {

	} else if tempResultCode == "41" {
		logger.Error("Error update momo payment " + tempResultCode)
		return c.JSON(http.StatusInternalServerError, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Tạo thanh toán thất bại",
			Data:       dataResponse,
		})
	} else {
		logger.Error("Error update momo payment " + tempResultCode)
		return c.JSON(http.StatusInternalServerError, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Tạo thanh toán thất bại",
			Data:       dataResponse,
		})
	}
	if err != nil {
		return response.BadRequest(c, "nil", nil)
	}
	return c.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "Tạo thanh toán thành công",
		Data:       dataResponse,
	})
}

func (paymentReceiver *PaymentController) GetResultPaymentVNPay(c echo.Context) error {
	dataFromVNPay := c.QueryParams()

	logger.Info(dataFromVNPay.Encode())
	dataFromVNPay.Del("vnp_SecureHash")

	return c.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "Thanh toán thành công",
		Data:       dataFromVNPay.Encode(),
	})
}

func (paymentReceiver *PaymentController) GetResultPaymentMomo(c echo.Context) error {
	logger.Info("Receive result from momo")
	jsonRequestMomo := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonRequestMomo)
	if err != nil {
		logger.Error("Error get momo body request", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Thanh toán thất bại",
			Data:       nil,
		})
	} else {
		resultCode := fmt.Sprint(jsonRequestMomo["resultCode"])
		orderId := fmt.Sprint(jsonRequestMomo["orderId"])
		logger.Info("Order id return from momo: " + orderId)
		arraySplitOrderId := strings.Split(orderId, "_")
		//bookingID := fmt.Sprint(arraySplitOrderId[0])
		paymentID := fmt.Sprint(arraySplitOrderId[1])
		//booking := model.Booking{
		//	ID:              bookingID,
		//	PaymentStatusID: 2,
		//}
		payment := model.Payment{
			SessionId: paymentID,
			Status:    "paid",
		}
		if resultCode == "0" {
			//check payment existed
			//payment.DueTimePayment = time.Now()
			//payment.PaymentTime = time.Now()
			_, err := paymentReceiver.PaymentRepo.UpdatePaymentStatusByBookingID(payment)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, res.Response{
					StatusCode: http.StatusInternalServerError,
					Message:    "Thanh toán thất bại",
					Data:       nil,
				})
			} else {
				if err != nil {
					return c.JSON(http.StatusInternalServerError, res.Response{
						StatusCode: http.StatusInternalServerError,
						Message:    "Thanh toán thất bại",
						Data:       nil,
					})
				}
			}
		} else {
			logger.Error("Error get momo result code request", zap.Error(err))
			payment.Status = "4"
			_, err := paymentReceiver.PaymentRepo.UpdatePaymentStatusFailed(payment)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, res.Response{
					StatusCode: http.StatusInternalServerError,
					Message:    "Thanh toán thất bại",
					Data:       nil,
				})
			}
			return c.JSON(http.StatusInternalServerError, res.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Thanh toán thất bại",
				Data:       nil,
			})
		}
	}
	return c.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "Thanh toán thành công",
		Data:       "",
	})
}

//// HandleGetListPaymentCondition godoc
//// @Summary Get list payment
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestPaymentList true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/list-payment [post]
//func (paymentReceiver *PaymentController) HandleGetListPaymentCondition(c echo.Context) error {
//	reqGetListPayment := req.RequestPaymentList{}
//	if err := c.Bind(&reqGetListPayment); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	var listPayment []model.Payment
//	condition := map[string]interface{}{
//		"isGetAll":       "false",
//		"full_name":      strings.ToLower(reqGetListPayment.CustomerName),
//		"payment_time":   reqGetListPayment.PaymentTime,
//		"status_payment": reqGetListPayment.StatusPayment,
//		"amount":         reqGetListPayment.Amount,
//		"payment_method": strings.ToLower(reqGetListPayment.PaymentMethod),
//	}
//
//	listPayment, err := paymentReceiver.PaymentRepo.GetPaymentListByCondition(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listPayment,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách thanh toán thành công",
//		Data:       listPayment,
//	})
//}
//
//// HandleGetListHistoryPayment godoc
//// @Summary Get list history payment
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetHistoryPaymentList true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/history-payment [post]
//func (paymentReceiver *PaymentController) HandleGetListHistoryPayment(c echo.Context) error {
//	reqGetListPayment := req.RequestGetHistoryPaymentList{}
//	if err := c.Bind(&reqGetListPayment); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.CUSTOMER.String() || claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	var listPayment []model.Payment
//	listPayment, err := paymentReceiver.PaymentRepo.GetPaymentHistoryList(reqGetListPayment.CustomerID)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listPayment,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy lịch sử thanh toán thành công",
//		Data:       listPayment,
//	})
//}
//
//// HandleUpdatePayment godoc
//// @Summary Update payment
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestUpdatePayment true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/update-payment [post]
//func (paymentReceiver *PaymentController) HandleUpdatePayment(c echo.Context) error {
//	reqUpdatePayment := req.RequestUpdatePayment{}
//	if err := c.Bind(&reqUpdatePayment); err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	err := c.Validate(reqUpdatePayment)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	timeDuePayment, err := time.Parse("2006-01-02", reqUpdatePayment.DueTimePayment)
//	if err != nil {
//		return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//	}
//	timePayment, err := time.Parse("2006-01-02", reqUpdatePayment.PaymentTime)
//	if err != nil {
//		return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//	}
//	payment := model.Payment{
//		ID:                reqUpdatePayment.PaymentID,
//		FineAmount:        reqUpdatePayment.FineAmount,
//		Amount:            reqUpdatePayment.Amount,
//		DueTimePayment:    timeDuePayment,
//		PaymentTime:       timePayment,
//		StatusPaymentCode: reqUpdatePayment.StatusPayment,
//	}
//	payment, err = paymentReceiver.PaymentRepo.UpdatePayment(payment)
//	if err != nil {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Cập nhật thành công",
//		Data:       payment,
//	})
//}
//
//// HandleCancelPayment godoc
//// @Summary Cancel payment
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestCancelPayment true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/cancel-payment [post]
//func (paymentReceiver *PaymentController) HandleCancelPayment(c echo.Context) error {
//	reqCancelPayment := req.RequestCancelPayment{}
//	if err := c.Bind(&reqCancelPayment); err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	err := c.Validate(reqCancelPayment)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.CUSTOMER.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	payment := model.Payment{
//		ID:                reqCancelPayment.PaymentID,
//		StatusPaymentCode: "3",
//	}
//	payment, err = paymentReceiver.PaymentRepo.UpdatePayment(payment)
//	if err != nil {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Hủy thanh toán thành công",
//		Data:       nil,
//	})
//}
//
//// HandleDeletePayment godoc
//// @Summary Delete payment
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestDeletePayment true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/delete-payment [post]
//func (paymentReceiver *PaymentController) HandleDeletePayment(c echo.Context) error {
//	reqDeletePayment := req.RequestDeletePayment{}
//	if err := c.Bind(&reqDeletePayment); err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	err := c.Validate(reqDeletePayment)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	payment := model.Payment{
//		ID: reqDeletePayment.PaymentID,
//	}
//	resultDeletePayment, err := paymentReceiver.PaymentRepo.DeletePayment(payment)
//	if err != nil || !resultDeletePayment {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Xóa thanh toán thành công",
//		Data:       nil,
//	})
//}
//
//// HandleCreatePayment godoc
//// @Summary Create payment
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestCreatePayment true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/create-payment-offline [post]
//func (paymentReceiver *PaymentController) HandleCreatePayment(c echo.Context) error {
//	reqCreatePayment := req.RequestCreatePayment{}
//	//binding
//	if err := c.Bind(&reqCreatePayment); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	paymentId, err := uuid.NewUUID()
//	payment := model.Payment{
//		ID:                paymentId.String(),
//		BookingID:         reqCreatePayment.BookingID,
//		CustomerID:        reqCreatePayment.CustomerID,
//		Amount:            reqCreatePayment.Amount,
//		DueTimePayment:    time.Now().Add(1 * time.Hour),
//		PaymentMethodID:   "2d81bba8-64e0-11ed-934f-089798c34e0e",
//		StatusPaymentCode: "1",
//	}
//	payment, err = paymentReceiver.PaymentRepo.SavePayment(payment)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lưu thành công",
//		Data:       payment,
//	})
//}
//
//// HandleSavePaymentMethod godoc
//// @Summary Save payment method
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddPaymentMethod true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/add-payment-method [post]
//func (paymentReceiver *PaymentController) HandleSavePaymentMethod(c echo.Context) error {
//	reqAddPaymentMethod := req.RequestAddPaymentMethod{}
//	//binding
//	if err := c.Bind(&reqAddPaymentMethod); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	paymentMethodId, err := uuid.NewUUID()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	paymentMethod := model.PaymentMethod{
//		ID:          paymentMethodId.String(),
//		MethodName:  reqAddPaymentMethod.MethodName,
//		Provider:    reqAddPaymentMethod.Provider,
//		Description: reqAddPaymentMethod.Description,
//	}
//	result, err := paymentReceiver.PaymentRepo.SavePaymentMethod(paymentMethod)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lưu thành công", result)
//}
//
//// HandleCreatePaymentBill godoc
//// @Summary Create payment bill
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestCreatePaymentBill true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/create-bill [post]
//func (paymentReceiver *PaymentController) HandleCreatePaymentBill(c echo.Context) error {
//	reqCreatePaymentBill := req.RequestCreatePaymentBill{}
//	if err := c.Bind(&reqCreatePaymentBill); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.CUSTOMER.String() || claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return c.JSON(http.StatusBadRequest, res.Response{
//			StatusCode: http.StatusBadRequest,
//			Message:    "Bạn không có quyền thực hiện chức năng này",
//			Data:       nil,
//		})
//	}
//	payment := model.Payment{
//		BookingID: reqCreatePaymentBill.BookingID,
//		ID:        reqCreatePaymentBill.PaymentID,
//	}
//	paymentResult, err := paymentReceiver.PaymentRepo.GetBillPayment(payment)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy thông tin hóa đơn thành công",
//		Data:       paymentResult,
//	})
//}
//
//// HandleGetAllPayments godoc
//// @Summary Get all payment list
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Success 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/payments [get]
//func (paymentReceiver *PaymentController) HandleGetAllPayments(c echo.Context) error {
//	var listPayment []model.Payment
//	condition := map[string]interface{}{
//		"isGetAll": "true",
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	listPayment, err := paymentReceiver.PaymentRepo.GetPaymentListByCondition(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listPayment,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách thanh toán thành công",
//		Data:       listPayment,
//	})
//}
//
//// HandleGetAllPaymentStatus godoc
//// @Summary Get all payment status list
//// @Tags payment-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /payment/payment-statuses [get]
//func (paymentReceiver *PaymentController) HandleGetAllPaymentStatus(c echo.Context) error {
//	var listPayment []model.PaymentStatus
//	listPayment, err := paymentReceiver.PaymentRepo.GetPaymentStatusList()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listPayment,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách trạng thái thanh toán thành công",
//		Data:       listPayment,
//	})
//}
