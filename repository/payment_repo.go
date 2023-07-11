package repository

import (
	"github.com/labstack/echo/v4"
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
	"hotel-booking-api/model/req"
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
	CancelBooking(payment model.Payment) (bool, error)
	GetWalletTopUp(userId string) (model.Wallet, error)
	UpdateAddMoneyToTopUp(wallet model.Wallet) (model.Wallet, error)
	GetPaymentDetail(payment model.Payment) (res.PaymentResponse, error)
	GetWalletTransactionsFilter(queryModel *query.DataQueryModel) ([]model.WalletTransaction, error)
	SaveWalletTransaction(walletTransaction model.WalletTransaction) (model.WalletTransaction, error)
	UpdateWalletTransactionStatus(walletTransaction model.WalletTransaction) (model.WalletTransaction, error)
	GetUserWalletInfo(userId string) (model.Wallet, error)
	ApplyVoucherPayments(requestApplyVoucher req.RequestApplyVoucher) (bool, error)
	UpdateUserRank(userRank model.UserRank) (bool, error)
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
