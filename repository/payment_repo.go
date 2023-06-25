package repository

type PaymentRepo interface {
	GetMomoHostingUrl() (string, error)
	GetRedirectMomoUrl() (string, error)
	//GetPaymentListByCondition(condition map[string]interface{}) ([]model.Payment, error)
	//GetPaymentHistoryList(customerID string) ([]model.Payment, error)
	//GetBillPayment(payment model.Payment) (model.Payment, error)
	//SavePayment(payment model.Payment) (model.Payment, error)
	//UpdatePayment(payment model.Payment) (model.Payment, error)
	//DeletePayment(payment model.Payment) (bool, error)
	//SavePaymentMethod(paymentMethod model.PaymentMethod) (model.PaymentMethod, error)
	//GetPaymentOnlineDetail(payment model.Payment) (model.PaymentDetailOnline, error)
	//GetPaymentOfflineDetail(payment model.Payment) (model.PaymentDetailOffline, error)
	//UpdatePaymentStatusBooking(booking model.Booking) (model.Booking, error)
	//UpdatePaymentStatusByBookingID(payment model.Payment) (model.Payment, error)
	//GetPaymentStatusList() ([]model.PaymentStatus, error)
	//SavePaymentOnlineDetail(paymentDetail model.PaymentDetailOnline) (model.PaymentDetailOnline, error)
	//DeletePaymentByBookingID(payment model.Payment) (bool, error)
	//UpdatePaymentStatusFailed(payment model.Payment) (model.Payment, error)
}
