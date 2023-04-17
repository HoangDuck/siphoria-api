package repository

//import (
//	"hotel-booking-api/model"
//)
//
//type BookingRepo interface {
//	SaveBooking(booking model.Booking) (model.Booking, error)
//	SaveMultipleBookingDetails(listBookingDetails []model.BookingDetail) (bool, error)
//	UpdateBooking(booking model.Booking) (model.Booking, error)
//	UpdateBookingDetail(bookingDetail model.BookingDetail) (model.BookingDetail, error)
//	CancelBooking(booking model.Booking) (model.Booking, error)
//	GetListBookingByCondition(condition map[string]interface{}) ([]model.Booking, error)
//	GetListHistoryBooking(customerID string) ([]model.Booking, error)
//	CheckProfileCustomerExistByIdentify(user model.Customer) (model.Customer, error)
//	CheckProfileCustomerExistByEmail(user model.Customer) (model.Customer, error)
//	SaveCustomerProfile(user model.Customer) (model.Customer, error)
//	GetRoomTypeInfo(roomType model.RoomType) (model.RoomType, error)
//	GetBookingStatusList() ([]model.StatusBooking, error)
//	GetBookingInfo(booking model.Booking) (model.Booking, error)
//	SaveRoomBusyStatusDetail(listRoomBusy []model.RoomBusyStatusDetail) (bool, error)
//	GetListBookingCheckInQuery() ([]model.Booking, error)
//	UpdateRoomByIDs(roomIDs []string) (model.Room, error)
//}
