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
	"log"
	"net/http"
	"os"
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
	redirectMomoUrl, err := paymentReceiver.PaymentRepo.GetRedirectPaymentUrl()
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
	redirectMomoUrl, err := paymentReceiver.PaymentRepo.GetRedirectPaymentUrl()
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
	isAddMoneyToWallet := false
	isUpdateRank := false
	var payment model.Payment
	var walletTransaction model.WalletTransaction
	var userRank model.UserRank
	dataFromVNPay := c.QueryParams()
	vnpayService := services.GetVNPayServiceInstance()
	resultCheckSignature := vnpayService.CheckSignatureResultVNPay(c)
	redirectMomoUrl, err := paymentReceiver.PaymentRepo.GetRedirectPaymentUrl()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	if resultCheckSignature {
		resultCode := fmt.Sprint(dataFromVNPay.Get("vnp_TransactionStatus"))
		orderId := fmt.Sprint(dataFromVNPay.Get("vnp_TxnRef"))
		logger.Info("Order id return from momo: " + orderId)
		arraySplitOrderId := strings.Split(orderId, "_")
		paymentID := fmt.Sprint(arraySplitOrderId[1])
		if arraySplitOrderId[2] == "add-siphoria-wallet" {
			isAddMoneyToWallet = true
			walletTransaction = model.WalletTransaction{
				ID: paymentID,
			}
		} else if arraySplitOrderId[2] == "update-rank" {
			isUpdateRank = true
			userRank = model.UserRank{
				ID:        paymentID,
				UserId:    arraySplitOrderId[4],
				RankId:    arraySplitOrderId[3],
				BeginAt:   time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
		} else {
			payment = model.Payment{
				SessionId: paymentID,
				Status:    "paid",
			}
		}
		if resultCode == "00" {
			if isAddMoneyToWallet {
				walletTransaction.Status = "success"
				_, err := paymentReceiver.PaymentRepo.UpdateWalletTransactionStatus(walletTransaction)
				if err != nil {
					logger.Info("redirect result payment vnpay update status payment")
					http.Redirect(c.Response(), c.Request(), redirectMomoUrl, http.StatusTemporaryRedirect)
					return c.Redirect(http.StatusInternalServerError, redirectMomoUrl)
				}
			} else if isUpdateRank {
				_, err := paymentReceiver.PaymentRepo.UpdateUserRank(userRank)
				if err != nil {
					http.Redirect(c.Response(), c.Request(), redirectMomoUrl, http.StatusTemporaryRedirect)
					return c.Redirect(http.StatusInternalServerError, redirectMomoUrl)
				}
			} else {
				_, err := paymentReceiver.PaymentRepo.UpdatePaymentStatusByBookingID(payment)
				if err != nil {
					logger.Info("redirect result payment vnpay update status payment")
					http.Redirect(c.Response(), c.Request(), redirectMomoUrl, http.StatusTemporaryRedirect)
					return c.Redirect(http.StatusInternalServerError, redirectMomoUrl)
				}
			}
		} else {
			if isAddMoneyToWallet {
				walletTransaction.Status = "failed"
				_, err := paymentReceiver.PaymentRepo.UpdateWalletTransactionStatus(walletTransaction)
				if err != nil {
					logger.Info("redirect result payment vnpay update status payment")
				}
			} else {
				payment.Status = "failed"
				_, err := paymentReceiver.PaymentRepo.UpdatePaymentStatusFailed(payment)
				if err != nil {
					logger.Info("redirect result payment vnpay update status payment failed")
				}
			}
			logger.Info("redirect result payment vnpay update status payment failed")
			http.Redirect(c.Response(), c.Request(), redirectMomoUrl, http.StatusTemporaryRedirect)
			return response.InternalServerError(c, "Thanh toán thất bại", nil)
		}
		logger.Info("redirect result payment vnpay")
		http.Redirect(c.Response(), c.Request(), redirectMomoUrl, http.StatusTemporaryRedirect)
		return response.Ok(c, "Thanh toán thành công", "")
	}
	http.Redirect(c.Response(), c.Request(), redirectMomoUrl, http.StatusTemporaryRedirect)
	return response.InternalServerError(c, "Thanh toán thất bại", nil)
}

func (paymentReceiver *PaymentController) GetResultPaymentMomo(c echo.Context) error {
	logger.Info("Receive result from momo")
	isAddMoneyToWallet := false
	isUpdateRank := false
	var payment model.Payment
	var walletTransaction model.WalletTransaction
	var userRank model.UserRank
	jsonRequestMomo := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonRequestMomo)
	if err != nil {
		logger.Error("Error get momo body request", zap.Error(err))
		return response.InternalServerError(c, "Thanh toán thất bại", nil)
	} else {
		resultCode := fmt.Sprint(jsonRequestMomo["resultCode"])
		orderId := fmt.Sprint(jsonRequestMomo["orderId"])
		logger.Info("Order id return from momo: " + orderId)
		arraySplitOrderId := strings.Split(orderId, "_")
		paymentID := fmt.Sprint(arraySplitOrderId[1])
		if arraySplitOrderId[2] == "add-siphoria-wallet" {
			isAddMoneyToWallet = true
			walletTransaction = model.WalletTransaction{
				ID: paymentID,
			}
		} else if arraySplitOrderId[2] == "update-rank" {
			isUpdateRank = true
			userRank = model.UserRank{
				ID:        paymentID,
				UserId:    arraySplitOrderId[4],
				RankId:    arraySplitOrderId[3],
				BeginAt:   time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
		} else {
			payment = model.Payment{
				SessionId: paymentID,
				Status:    "paid",
			}
		}
		if resultCode == "0" {
			if isAddMoneyToWallet {
				walletTransaction.Status = "success"
				_, err := paymentReceiver.PaymentRepo.UpdateWalletTransactionStatus(walletTransaction)
				if err != nil {
					logger.Info("redirect result payment vnpay update status payment")
					return response.InternalServerError(c, "Thanh toán thất bại", nil)
				}
			} else if isUpdateRank {
				_, err := paymentReceiver.PaymentRepo.UpdateUserRank(userRank)
				if err != nil {
					return response.InternalServerError(c, "Thanh toán thất bại", nil)
				}
			} else {
				_, err := paymentReceiver.PaymentRepo.UpdatePaymentStatusByBookingID(payment)
				if err != nil {
					return response.InternalServerError(c, "Thanh toán thất bại", nil)
				}
			}
		} else {
			logger.Error("Error get momo result code request", zap.Error(err))
			if isAddMoneyToWallet {
				walletTransaction.Status = "failed"
				_, err := paymentReceiver.PaymentRepo.UpdateWalletTransactionStatus(walletTransaction)
				if err != nil {
					return response.InternalServerError(c, "Thanh toán thất bại", nil)
				}
			} else {
				payment.Status = "failed"
				_, err := paymentReceiver.PaymentRepo.UpdatePaymentStatusFailed(payment)
				if err != nil {
					return response.InternalServerError(c, "Thanh toán thất bại", nil)
				}
			}

			return response.InternalServerError(c, "Thanh toán thất bại", nil)
		}
	}
	return response.Ok(c, "Thanh toán thành công", "")
}

func (paymentReceiver *PaymentController) HandleWebHookStripe(c echo.Context) error {
	stripeService := services.GetStripeServiceInstance()
	eventType := stripeService.CheckWebHookEvent(c)
	logger.Infof("event type cong hoa", eventType)
	switch eventType {
	case "payment_intent.succeeded":

		log.Printf("Successful payment for %d.", int64(1234))
		// Then define and call a func to handle the successful payment intent.
		// handlePaymentIntentSucceeded(paymentIntent)
	case "payment_method.attached":
		return response.Ok(c, "Thanh toán thành công", eventType)
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handlePaymentMethodAttached(paymentMethod)
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", eventType)
	}
	return response.Ok(c, "Thanh toán thành công", "")
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
