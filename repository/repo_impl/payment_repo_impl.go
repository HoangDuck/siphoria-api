package repo_impl

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/repository"
)

type PaymentRepoImpl struct {
	sql *db.Sql
}

func NewPaymentRepo(sql *db.Sql) repository.PaymentRepo {
	return &PaymentRepoImpl{
		sql: sql,
	}
}

func (paymentReceiver *PaymentRepoImpl) UpdatePaymentMethodForPending(sessionId string, paymentMethod string) (bool, error) {
	var payment = model.Payment{
		PaymentMethod: paymentMethod,
	}
	err := paymentReceiver.sql.Db.Where("session_id=? AND status = ?", sessionId, "pending").Updates(payment)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return false, custom_error.PaymentNotFound
		}
		return false, custom_error.PaymentNotUpdated
	}
	return true, nil
}

func (paymentReceiver *PaymentRepoImpl) GetVNPayHostingUrl() (string, error) {
	var vnpayConfig model.ConfigurationUrlDefine
	err := paymentReceiver.sql.Db.Where("id=?", 5).Find(&vnpayConfig)
	logger.Debug("Get momo url", zap.Error(err.Error))
	if err.Error == gorm.ErrRecordNotFound {
		return "noUrl", err.Error
	}
	return vnpayConfig.Value, nil
}

func (paymentReceiver *PaymentRepoImpl) GetPaymentListByCondition(sessionId string) ([]model.Payment, error) {
	var listPayment []model.Payment
	err := paymentReceiver.sql.Db.Where("session_id = ?", sessionId).Find(&listPayment)
	if err.Error != nil {
		return listPayment, err.Error
	}
	return listPayment, nil
}

//	func (paymentReceiver *PaymentRepoImpl) GetPaymentHistoryList(customerID string) ([]model.Payment, error) {
//		var listHistoryPayment []model.Payment
//		err := paymentReceiver.sql.Db.Where("customer_id=?", customerID).Preload("Booking").Preload("PaymentMethod").Find(&listHistoryPayment)
//		logger.Debug("Get payment data", zap.Error(err.Error))
//		if err.Error == gorm.ErrRecordNotFound {
//			return listHistoryPayment, custom_error.PaymentsEmpty
//		}
//		return listHistoryPayment, nil
//	}
//
//	func (paymentReceiver *PaymentRepoImpl) GetBillPayment(payment model.Payment) (model.Payment, error) {
//		var paymentResult = model.Payment{}
//		err := paymentReceiver.sql.Db.Model(&payment).Preload("Booking").Preload("PaymentMethod").Find(&paymentResult)
//		logger.Debug("Get payment data", zap.Error(err.Error))
//		if err.Error == gorm.ErrRecordNotFound {
//			return paymentResult, custom_error.PaymentNotFound
//		}
//		return paymentResult, nil
//	}
//
//	func (paymentReceiver *PaymentRepoImpl) SavePayment(payment model.Payment) (model.Payment, error) {
//		result := paymentReceiver.sql.Db.Preload("Booking").Preload("PaymentMethod").Create(&payment)
//		if result.Error != nil {
//			logger.Error("Get payment data", zap.Error(result.Error))
//			return payment, custom_error.PaymentNotSaved
//		}
//		return payment, nil
//	}
//
//	func (paymentReceiver *PaymentRepoImpl) UpdatePayment(payment model.Payment) (model.Payment, error) {
//		var paymentResult model.Payment
//		err := paymentReceiver.sql.Db.Model(&paymentResult).Where("id=?", payment.ID).Preload("Booking").Preload("PaymentMethod").Updates(payment)
//		if err.Error == gorm.ErrRecordNotFound {
//			return paymentResult, custom_error.PaymentNotFound
//		}
//		return paymentResult, nil
//	}
//
//	func (paymentReceiver *PaymentRepoImpl) DeletePayment(payment model.Payment) (bool, error) {
//		err := paymentReceiver.sql.Db.Model(&payment).Where("id=?", payment.ID).Delete(payment)
//		if err != nil {
//			logger.Error("Get delete data", zap.Error(err.Error))
//			return false, custom_error.PaymentDeleteFailed
//		}
//		return true, nil
//	}
//
//	func (paymentReceiver *PaymentRepoImpl) SavePaymentMethod(paymentMethod model.PaymentMethod) (model.PaymentMethod, error) {
//		result := paymentReceiver.sql.Db.Create(&paymentMethod)
//		if result.Error != nil {
//			return paymentMethod, custom_error.PaymentNotSaved
//		}
//		return paymentMethod, nil
//	}
//
//	func (paymentReceiver *PaymentRepoImpl) GetPaymentOnlineDetail(payment model.Payment) (model.PaymentDetailOnline, error) {
//		var paymentResult = model.PaymentDetailOnline{}
//		err := paymentReceiver.sql.Db.Model(&payment).Where("payment_id=?", payment.ID).Preload("Payment").Find(&paymentResult)
//		if err != nil {
//			return paymentResult, custom_error.PaymentNotFound
//		}
//		return paymentResult, nil
//	}
//
//	func (paymentReceiver *PaymentRepoImpl) GetPaymentOfflineDetail(payment model.Payment) (model.PaymentDetailOffline, error) {
//		var paymentResult = model.PaymentDetailOffline{}
//		err := paymentReceiver.sql.Db.Model(&payment).Where("payment_id=?", payment.ID).Preload("Payment").Find(&paymentResult)
//		if err != nil {
//			return paymentResult, custom_error.PaymentNotFound
//		}
//		return paymentResult, nil
//	}
//
//	func (paymentReceiver *PaymentRepoImpl) UpdatePaymentStatusBooking(booking model.Booking) (model.Booking, error) {
//		var bookingResult model.Booking
//		err := paymentReceiver.sql.Db.Model(&bookingResult).Where("ID=?", booking.ID).Updates(booking)
//		if err.Error != nil {
//			logger.Debug("Get update payment data", zap.Error(err.Error))
//			if err.Error == gorm.ErrRecordNotFound {
//				return bookingResult, custom_error.PaymentNotFound
//			}
//
//			return bookingResult, custom_error.PaymentNotUpdated
//		}
//		return bookingResult, nil
//	}
func (paymentReceiver *PaymentRepoImpl) UpdatePaymentStatusFailed(payment model.Payment) (model.Payment, error) {
	var paymentResult model.Payment
	tempPaymentFailed := model.Payment{
		Status: "refunding",
	}
	err := paymentReceiver.sql.Db.Model(&paymentResult).Where("session_id = ?", payment.SessionId).Updates(tempPaymentFailed)
	if err.Error != nil {
		logger.Error("Error update momo payment", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return paymentResult, custom_error.PaymentNotFound
		}

		return paymentResult, custom_error.PaymentNotUpdated
	}
	return paymentResult, nil
}

func (paymentReceiver *PaymentRepoImpl) UpdatePaymentStatusByBookingID(payment model.Payment) (model.Payment, error) {
	//update payment success
	var paymentResult model.Payment
	err := paymentReceiver.sql.Db.Model(&paymentResult).Where("session_id=?", payment.SessionId).Updates(payment)
	if err.Error != nil {
		logger.Error("Error update momo payment", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return paymentResult, custom_error.PaymentNotFound
		}

		return paymentResult, custom_error.PaymentNotUpdated
	}
	return paymentResult, nil
}

//
//func (paymentReceiver *PaymentRepoImpl) GetPaymentStatusList() ([]model.PaymentStatus, error) {
//	var listPayment []model.PaymentStatus
//	err := paymentReceiver.sql.Db.Find(&listPayment)
//	if err != nil {
//		return listPayment, err.Error
//	}
//	return listPayment, err.Error
//}

func (paymentReceiver *PaymentRepoImpl) GetMomoHostingUrl() (string, error) {
	var momoConfig model.ConfigurationUrlDefine
	err := paymentReceiver.sql.Db.Where("id=?", 3).Find(&momoConfig)
	logger.Debug("Get momo url", zap.Error(err.Error))
	if err.Error == gorm.ErrRecordNotFound {
		return "noUrl", err.Error
	}
	return momoConfig.Value, nil
}

//func (paymentReceiver *PaymentRepoImpl) SavePaymentOnlineDetail(paymentDetail model.PaymentDetailOnline) (model.PaymentDetailOnline, error) {
//	result := paymentReceiver.sql.Db.Create(&paymentDetail)
//	if result.Error != nil {
//		return paymentDetail, custom_error.PaymentNotSaved
//	}
//	return paymentDetail, nil
//}

func (paymentReceiver *PaymentRepoImpl) GetRedirectMomoUrl() (string, error) {
	var momoConfig model.ConfigurationUrlDefine
	err := paymentReceiver.sql.Db.Where("id=?", 4).Find(&momoConfig)
	logger.Debug("Get momo url", zap.Error(err.Error))
	if err.Error == gorm.ErrRecordNotFound {
		return "noUrl", err.Error
	}
	return momoConfig.Value, nil
}

//func (paymentReceiver *PaymentRepoImpl) DeletePaymentByBookingID(payment model.Payment) (bool, error) {
//	err := paymentReceiver.sql.Db.Model(&payment).Where("booking_id=?", payment.BookingID).Delete(payment)
//	if err != nil {
//		logger.Error("Get delete data", zap.Error(err.Error))
//		return false, custom_error.PaymentDeleteFailed
//	}
//	return true, nil
//}
