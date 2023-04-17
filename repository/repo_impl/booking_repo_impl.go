package repo_impl

//
//import (
//	"fmt"
//	"go.uber.org/zap"
//	"gorm.io/gorm"
//	"hotel-booking-api/custom_error"
//	"hotel-booking-api/db"
//	"hotel-booking-api/logger"
//	"hotel-booking-api/model"
//	"hotel-booking-api/repository"
//)
//
//type BookingRepoImpl struct {
//	sql *db.Sql
//}
//
//func NewBookingRepo(sql *db.Sql) repository.BookingRepo {
//	return &BookingRepoImpl{
//		sql: sql,
//	}
//}
//
//func (bookingReceiver *BookingRepoImpl) SaveCustomerProfile(user model.Customer) (model.Customer, error) {
//	result := bookingReceiver.sql.Db.Create(&user)
//	if result.Error != nil {
//		return user, result.Error
//	}
//	return user, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) CheckProfileCustomerExistByEmail(user model.Customer) (model.Customer, error) {
//	err := bookingReceiver.sql.Db.First(&user, "email = ?", user.Email)
//	if err.Error == gorm.ErrRecordNotFound {
//		return user, custom_error.UserNotFound
//	}
//	return user, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) CheckProfileCustomerExistByIdentify(user model.Customer) (model.Customer, error) {
//	err := bookingReceiver.sql.Db.First(&user, "identifier_number = ?", user.IdentifierNumber)
//	//err := bookingReceiver.sql.Db.Where("identifier_number = ?", user.IdentifierNumber).Find(&user)
//	//if err != nil {
//	//	logger.Error("Error query data", zap.Error(err.Error))
//	//	if err.Error == gorm.ErrRecordNotFound {
//	//		return user, custom_error.UserNotFound
//	//	}
//	//	return user, err.Error
//	//}
//	if err.Error == gorm.ErrRecordNotFound {
//		return user, custom_error.UserNotFound
//	}
//	return user, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) SaveMultipleBookingDetails(listBookingDetails []model.BookingDetail) (bool, error) {
//	result := bookingReceiver.sql.Db.CreateInBatches(listBookingDetails, len(listBookingDetails))
//	if result.Error != nil {
//		logger.Error("Error save data", zap.Error(result.Error))
//		return false, custom_error.BookingNotSaved
//	}
//	return true, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) SaveBooking(booking model.Booking) (model.Booking, error) {
//	result := bookingReceiver.sql.Db.Create(&booking) //.Preload("Customer").Preload("RoomType").Preload("RatePlan")
//	if result.Error != nil {
//		logger.Error("Error save data", zap.Error(result.Error))
//		return booking, custom_error.BookingNotSaved
//	}
//	return booking, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) CancelBooking(booking model.Booking) (model.Booking, error) {
//	var bookingResult = model.Booking{}
//	err := bookingReceiver.sql.Db.Model(&bookingResult).Where("ID=?", booking.ID).Preload("Customer").Preload("RoomType").Preload("RatePlan").Updates(booking)
//	if err.Error != nil {
//		logger.Debug("Save data", zap.Error(err.Error))
//		if err.Error == gorm.ErrRecordNotFound {
//			logger.Error("Error save data", zap.Error(err.Error))
//			return bookingResult, custom_error.BookingNotFound
//		}
//		logger.Error("Error save data", zap.Error(err.Error))
//		return bookingResult, custom_error.BookingNotUpdated
//	}
//	return bookingResult, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) GetListHistoryBooking(customerID string) ([]model.Booking, error) {
//	var listHistoryBooking []model.Booking
//	err := bookingReceiver.sql.Db.Where("customer_id=?", customerID).Preload("Customer").Preload("RoomType").Preload("RatePlan").Preload("StatusBooking").Find(&listHistoryBooking)
//	if err.Error == gorm.ErrRecordNotFound {
//		return listHistoryBooking, custom_error.BookingsEmpty
//	}
//	return listHistoryBooking, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) UpdateBooking(booking model.Booking) (model.Booking, error) {
//	var bookingResult model.Booking
//	err := bookingReceiver.sql.Db.Model(&bookingResult).Where("ID=?", booking.ID).Preload("Customer").Preload("RoomType").Preload("RatePlan").Preload("StatusBooking").Updates(booking)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return bookingResult, custom_error.BookingNotFound
//		}
//		return bookingResult, custom_error.BookingNotUpdated
//	}
//	return bookingResult, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) UpdateBookingDetail(bookingDetail model.BookingDetail) (model.BookingDetail, error) {
//	var bookingResult model.BookingDetail
//	err := bookingReceiver.sql.Db.Model(&bookingResult).Where("ID=?", bookingResult.BookingID).Preload("Booking").Updates(bookingDetail)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return bookingResult, custom_error.BookingNotFound
//		}
//		return bookingResult, custom_error.BookingNotUpdated
//	}
//	return bookingResult, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) GetListBookingByCondition(condition map[string]interface{}) ([]model.Booking, error) {
//	var listBooking []model.Booking
//	if condition["isGetAll"] == "true" {
//		err := bookingReceiver.sql.Db.Preload("Customer").Preload("RoomType").Preload("RatePlan").Find(&listBooking)
//		if err != nil {
//			return listBooking, err.Error
//		}
//	} else {
//		queryString := fmt.Sprintf("Select * from bookings "+
//			"where customer_id in (select ID from customers where lower(full_name) LIKE '%%%s%%') "+
//			"AND status_booking_id = '%s' "+
//			"AND room_type_id in (select ID from room_types where lower(type_room_code) LIKE '%%%s%%') "+
//			"AND total_cost>=%v "+
//			"AND payment_status_id = %d "+
//			"AND '%s'::date <= check_out_time "+
//			"AND '%s'::date >= check_in_time",
//			condition["customer_name"],
//			condition["status_booking"],
//			condition["room_type_code"],
//			condition["total_cost"],
//			condition["status_payment"],
//			condition["check_in_time"],
//			condition["check_out_time"])
//		err := bookingReceiver.sql.Db.Raw(queryString).Preload("RoomType").Preload("RatePlan").Find(&listBooking)
//		if err != nil {
//			logger.Error("Error get data", zap.Error(err.Error))
//			return listBooking, err.Error
//		}
//	}
//	return listBooking, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) GetRoomTypeInfo(roomType model.RoomType) (model.RoomType, error) {
//	var roomTypeResult = model.RoomType{}
//	err := bookingReceiver.sql.Db.Where("type_room_code =?", roomType.Name).Find(&roomTypeResult)
//	if err != nil {
//		return roomTypeResult, err.Error
//	}
//	return roomTypeResult, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) GetBookingStatusList() ([]model.StatusBooking, error) {
//	var listBookingStatus []model.StatusBooking
//	err := bookingReceiver.sql.Db.Find(&listBookingStatus)
//	if err != nil {
//		return listBookingStatus, err.Error
//	}
//	return listBookingStatus, err.Error
//}
//
//func (bookingReceiver *BookingRepoImpl) GetBookingInfo(booking model.Booking) (model.Booking, error) {
//	var bookingResult = model.Booking{}
//	err := bookingReceiver.sql.Db.Where("id=?", booking.ID).Preload("Customer").Preload("RoomType").Preload("RatePlan").Preload("StatusBooking").Find(&bookingResult)
//	if err.Error == gorm.ErrRecordNotFound {
//		return bookingResult, custom_error.BookingNotFound
//	}
//	return bookingResult, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) SaveRoomBusyStatusDetail(listRoomBusy []model.RoomBusyStatusDetail) (bool, error) {
//	result := bookingReceiver.sql.Db.CreateInBatches(listRoomBusy, len(listRoomBusy))
//	if result.Error != nil {
//		return false, custom_error.BookingNotSaved
//	}
//	return true, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) GetListBookingCheckInQuery() ([]model.Booking, error) {
//	var bookingResult []model.Booking
//	err := bookingReceiver.sql.Db.Where("status_booking_id=? AND payment_status_id=?", "1", "2").Preload("Customer").Preload("RoomType").Preload("RatePlan").Preload("StatusBooking").Find(&bookingResult)
//	if err.Error == gorm.ErrRecordNotFound {
//		return bookingResult, custom_error.BookingNotFound
//	}
//	return bookingResult, nil
//}
//
//func (bookingReceiver *BookingRepoImpl) UpdateRoomByIDs(roomIDs []string) (model.Room, error) {
//	var roomResult = model.Room{}
//	err := bookingReceiver.sql.Db.Table("rooms").Where("id IN ?", roomIDs).Updates(map[string]interface{}{"room_busy_status_category_id": "e5d5bfbb-64dd-11ed-837f-089798c34e0e"})
//	//err := bookingReceiver.sql.Db.Model(&roomResult).Where("room_code=?", room.RoomCode).Preload("RoomType").Updates(room)
//	if err.Error != nil {
//		if err.Error == gorm.ErrRecordNotFound {
//			return roomResult, err.Error
//		}
//
//		return roomResult, custom_error.RoomNotUpdated
//	}
//	return roomResult, nil
//}
