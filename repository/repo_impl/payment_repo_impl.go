package repo_impl

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/model/query"
	"hotel-booking-api/model/req"
	"hotel-booking-api/model/res"
	"hotel-booking-api/repository"
	"time"
)

type PaymentRepoImpl struct {
	sql *db.Sql
}

func NewPaymentRepo(sql *db.Sql) repository.PaymentRepo {
	return &PaymentRepoImpl{
		sql: sql,
	}
}

func (paymentReceiver *PaymentRepoImpl) ApplyVoucherPayments(requestApplyVoucher req.RequestApplyVoucher) (bool, error) {
	err := paymentReceiver.sql.Db.Exec("call sp_applyvoucher(?,?);",
		requestApplyVoucher.Code, requestApplyVoucher.SessionId)
	if err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

func (paymentReceiver *PaymentRepoImpl) GetUserWalletInfo(userId string) (model.Wallet, error) {
	var wallet model.Wallet
	err := paymentReceiver.sql.Db.Where("user_id=?", userId).Find(&wallet)
	logger.Debug("Get wallet failed", zap.Error(err.Error))
	if err.Error == gorm.ErrRecordNotFound {
		return wallet, err.Error
	}
	return wallet, nil
}

func (paymentReceiver *PaymentRepoImpl) UpdateWalletTransactionStatus(walletTransaction model.WalletTransaction) (model.WalletTransaction, error) {
	err := paymentReceiver.sql.Db.Model(&walletTransaction).Where("id = ?", walletTransaction.ID).Updates(walletTransaction)
	if err.Error != nil {
		logger.Error("Error update wallet transaction payment", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return walletTransaction, custom_error.PaymentNotFound
		}

		return walletTransaction, custom_error.PaymentNotUpdated
	}
	return walletTransaction, nil
}

func (paymentReceiver *PaymentRepoImpl) SaveWalletTransaction(walletTransaction model.WalletTransaction) (model.WalletTransaction, error) {
	result := paymentReceiver.sql.Db.Create(&walletTransaction)
	if result.Error != nil {
		return walletTransaction, result.Error
	}
	return walletTransaction, nil
}

func (paymentReceiver *PaymentRepoImpl) GetWalletTransactionsFilter(queryModel *query.DataQueryModel) ([]model.WalletTransaction, error) {
	var listWalletTransactions []model.WalletTransaction
	err := GenerateQueryGetData(paymentReceiver.sql, queryModel, &model.WalletTransaction{}, queryModel.ListIgnoreColumns)
	err = err.Where("wallet_id in (Select id from wallets where user_id = ?)", queryModel.UserId)
	err = err.Find(&listWalletTransactions)
	if err.Error != nil {
		logger.Error("Error get list wallet transaction url ", zap.Error(err.Error))
		return listWalletTransactions, err.Error
	}
	return listWalletTransactions, nil
}

func (paymentReceiver *PaymentRepoImpl) GetPaymentDetail(payment model.Payment) (res.PaymentResponse, error) {
	var tempPayment model.Payment
	var paymentResponse res.PaymentResponse
	var listTempPaymentDetail []model.PaymentDetail
	var listPaymentDetailResponse []res.PaymentDetailResponse
	err := paymentReceiver.sql.Db.Where("id = ?", payment.ID).Find(&tempPayment)
	if err.Error != nil {
		return paymentResponse, err.Error
	}
	paymentResponse = res.PaymentResponse{
		ID:             tempPayment.ID,
		PaymentMethod:  tempPayment.PaymentMethod,
		RankPrice:      tempPayment.RankPrice,
		ConvertedPrice: tempPayment.ConvertedPrice,
		VoucherPrice:   tempPayment.VoucherPrice,
		TotalPrice:     tempPayment.TotalPrice,
		StartAt:        tempPayment.StartAt,
		EndAt:          tempPayment.EndAt,
		TotalDay:       tempPayment.TotalDay,
		UpdatedAt:      tempPayment.UpdatedAt,
		User:           &tempPayment.User,
		RoomType:       tempPayment.RoomType,
		Hotel:          tempPayment.Hotel,
	}
	err = paymentReceiver.sql.Db.Where("payment_id = ?", paymentResponse.ID).Find(&listTempPaymentDetail)
	if err.Error != nil {
		logger.Error("Error get list payment url ", zap.Error(err.Error))
		return paymentResponse, err.Error
	}

	for indexDetail := 0; indexDetail < len(listTempPaymentDetail); indexDetail++ {
		listPaymentDetailResponse = append(listPaymentDetailResponse, res.PaymentDetailResponse{
			ID:          listTempPaymentDetail[indexDetail].ID,
			Price:       listTempPaymentDetail[indexDetail].Price,
			DayOff:      listTempPaymentDetail[indexDetail].DayOff,
			AdultNum:    listTempPaymentDetail[indexDetail].AdultNum,
			ChildrenNum: listTempPaymentDetail[indexDetail].ChildrenNum,
		})
	}
	paymentResponse.Details = listPaymentDetailResponse
	return paymentResponse, nil
}

func (paymentReceiver *PaymentRepoImpl) UpdateAddMoneyToTopUp(wallet model.Wallet) (model.Wallet, error) {
	err := paymentReceiver.sql.Db.Model(&wallet).Where("user_id=?", wallet.UserId).Updates(wallet)
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return wallet, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return wallet, err.Error
		}
		return wallet, err.Error
	}
	return wallet, nil
}

func (paymentReceiver *PaymentRepoImpl) GetWalletTopUp(userId string) (model.Wallet, error) {
	wallet := model.Wallet{}
	err := paymentReceiver.sql.Db.Where("user_id = ?", userId).Find(&wallet)
	if err.Error != nil {
		logger.Error("Error get wallet", zap.Error(err.Error))
		return wallet, err.Error
	}
	return wallet, nil
}

func (paymentReceiver *PaymentRepoImpl) CancelBooking(payment model.Payment) (bool, error) {
	err := paymentReceiver.sql.Db.Model(&payment).Where("payment_id=?", payment.ID).Updates(payment)
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return false, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return false, err.Error
		}
		return false, err.Error
	}
	return true, nil
}

func (paymentReceiver *PaymentRepoImpl) GetPaymentFilter(context echo.Context, queryModel *query.DataQueryModel) ([]res.PaymentResponse, error) {
	var listPayment []res.PaymentResponse
	var listTempPayment []model.Payment

	err := GenerateQueryGetData(paymentReceiver.sql, queryModel, &model.Payment{}, queryModel.ListIgnoreColumns)
	err = err.Preload("User").Preload("RoomType").Preload("Voucher").Preload("PayoutRequest").
		Preload("Hotel").Preload("RatePlan")
	err = err.Where("user_id = ?", queryModel.UserId)

	if context.QueryParam("state") == "" {
		logger.Error(context.QueryParam("state"))
		queryModel.DataId = "paid"
	} else {
		logger.Error(context.QueryParam("state"))
		queryModel.DataId = context.QueryParam("state")
	}
	switch queryModel.DataId {
	case "paid":
		{
			err = err.Where("end_at > ? AND status = ?", time.Now(), queryModel.DataId)
			break
		}
	case "refunded":
		{
			err = err.Where("status = ?", queryModel.DataId)
			break
		}
	case "cancel":
		{
			err = err.Where("status = ?", queryModel.DataId)
			break
		}
	case "history":
		{
			err = err.Where("end_at < ? AND (status = ? OR status = ?)", time.Now(), "paid", "checked")
			break
		}
	default:
		err = err.Where("end_at > ? AND status = ?", time.Now(), queryModel.DataId)
	}
	var countTotalRows int64
	err.Model(model.Payment{}).Count(&countTotalRows)
	queryModel.TotalRows = int(countTotalRows)
	err = err.Find(&listTempPayment)
	for index := 0; index < len(listTempPayment); index++ {
		listPayment = append(listPayment, res.PaymentResponse{
			ID:             listTempPayment[index].ID,
			PaymentMethod:  listTempPayment[index].PaymentMethod,
			RankPrice:      listTempPayment[index].RankPrice,
			ConvertedPrice: listTempPayment[index].ConvertedPrice,
			VoucherPrice:   listTempPayment[index].VoucherPrice,
			TotalPrice:     listTempPayment[index].TotalPrice,
			StartAt:        listTempPayment[index].StartAt,
			EndAt:          listTempPayment[index].EndAt,
			TotalDay:       listTempPayment[index].TotalDay,
			UpdatedAt:      listTempPayment[index].UpdatedAt,
			User:           &listTempPayment[index].User,
			RoomType:       listTempPayment[index].RoomType,
			RatePlan:       listTempPayment[index].RatePlan,
			Hotel:          listTempPayment[index].Hotel,
		})
	}
	for index := 0; index < len(listPayment); index++ {
		var listTempPaymentDetail []model.PaymentDetail
		var listPaymentDetail []res.PaymentDetailResponse

		err = paymentReceiver.sql.Db.Where("payment_id = ?", listPayment[index].ID).Find(&listTempPaymentDetail)
		if err.Error != nil {
			logger.Error("Error get list payment url ", zap.Error(err.Error))
			continue
		}

		for indexDetail := 0; indexDetail < len(listTempPaymentDetail); indexDetail++ {
			listPaymentDetail = append(listPaymentDetail, res.PaymentDetailResponse{
				ID:          listTempPaymentDetail[indexDetail].ID,
				Price:       listTempPaymentDetail[indexDetail].Price,
				DayOff:      listTempPaymentDetail[indexDetail].DayOff,
				AdultNum:    listTempPaymentDetail[indexDetail].AdultNum,
				ChildrenNum: listTempPaymentDetail[indexDetail].ChildrenNum,
			})
		}
		listPayment[index].Details = listPaymentDetail
	}

	if listPayment == nil {
		listPayment = []res.PaymentResponse{}
	}
	if err.Error != nil {
		logger.Error("Error get list payment url ", zap.Error(err.Error))
		return listPayment, err.Error
	}
	return listPayment, nil
}

func (paymentReceiver *PaymentRepoImpl) CancelSessionPayment(userId string) (bool, error) {
	var lockRoom = model.LockRoom{
		LockFrom: time.Now(),
		LockTo:   time.Now().Add(time.Minute * 5),
	}
	err := paymentReceiver.sql.Db.Where("user_id=? AND expired = ?", userId, false).Updates(lockRoom)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return false, custom_error.PaymentNotFound
		}
		return false, custom_error.PaymentNotUpdated
	}
	return true, nil
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

func (paymentReceiver *PaymentRepoImpl) GetRedirectPaymentUrl() (string, error) {
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
