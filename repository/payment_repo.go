package repository

import (
	"github.com/labstack/echo/v4"
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
	"hotel-booking-api/model/res"
)

type PaymentRepo interface {
	GetMomoHostingUrl() (string, error)
	GetVNPayHostingUrl() (string, error)
	GetRedirectPaymentUrl() (string, error)
	GetPaymentListByCondition(sessionId string) ([]model.Payment, error)
	UpdatePaymentStatusByBookingID(payment model.Payment) (model.Payment, error)
	UpdatePaymentStatusFailed(payment model.Payment) (model.Payment, error)
	UpdatePaymentMethodForPending(sessionId string, paymentMethod string) (bool, error)
	CancelSessionPayment(userId string) (bool, error)
	GetPaymentFilter(context echo.Context, queryModel *query.DataQueryModel) ([]res.PaymentResponse, error)
	//GetPaymentHistoryList(customerID string) ([]model.Payment, error)
	//GetBillPayment(payment model.Payment) (model.Payment, error)
	//SavePayment(payment model.Payment) (model.Payment, error)
	//UpdatePayment(payment model.Payment) (model.Payment, error)
	//DeletePayment(payment model.Payment) (bool, error)
	//SavePaymentMethod(paymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	//GetPaymentOnlineDetail(payment model.Payment) (model.PaymentDetailOnline, error)
	//GetPaymentOfflineDetail(payment model.Payment) (model.PaymentDetailOffline, error)
	//UpdatePaymentStatusBooking(booking model.Booking) (model.Booking, error)
	//GetPaymentStatusList() ([]model.PaymentStatus, error)
	//SavePaymentOnlineDetail(paymentDetail model.PaymentDetailOnline) (model.PaymentDetailOnline, error)
	//DeletePaymentByBookingID(payment model.Payment) (bool, error)
}
