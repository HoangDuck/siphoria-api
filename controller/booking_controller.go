package controller

//
//import (
//	"fmt"
//	"github.com/golang-jwt/jwt"
//	"github.com/google/uuid"
//	"github.com/labstack/echo/v4"
//	"go.uber.org/zap"
//	"hotel-booking-api/custom_error"
//	"hotel-booking-api/logger"
//	"hotel-booking-api/model"
//	"hotel-booking-api/model/model_func"
//	"hotel-booking-api/model/req"
//	"hotel-booking-api/model/res"
//	"hotel-booking-api/repository"
//	"net/http"
//	"strings"
//	"time"
//)
//
//type BookingController struct {
//	BookingRepo repository.BookingRepo
//}
//
//// HandleCreateBooking godoc
//// @Summary Create booking
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddBooking true "booking"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /booking/create-booking [post]
//func (bookingReceiver *BookingController) HandleCreateBooking(c echo.Context) error {
//	reqAddBooking := req.RequestAddBooking{}
//	//binding
//	if err := c.Bind(&reqAddBooking); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.CUSTOMER.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	//check existed customer profile
//	customer := model.Customer{
//		FirstName:        reqAddBooking.FirstName,
//		LastName:         reqAddBooking.LastName,
//		Phone:            reqAddBooking.Phone,
//		Email:            reqAddBooking.Email,
//		Gender:           reqAddBooking.Gender,
//		IdentifierNumber: reqAddBooking.IdentifierNumber,
//	}
//	customerResult, err := model.Customer{}, custom_error.UserConflict
//	if reqAddBooking.TypeBooking == "offline" {
//		customerResult, err = bookingReceiver.BookingRepo.CheckProfileCustomerExistByIdentify(customer)
//	} else {
//		customerResult, err = bookingReceiver.BookingRepo.CheckProfileCustomerExistByEmail(customer)
//	}
//	if err == custom_error.UserNotFound {
//		customerId, err := uuid.NewUUID()
//		if err != nil {
//			return response.InternalServerError(c, err.Error(), nil)
//		}
//		customer.CustomerID = customerId.String()
//		customer.ID = customerId.String()
//		customerResult, err = bookingReceiver.BookingRepo.SaveCustomerProfile(customer)
//	}
//	logger.Errorf("Error query data", "sad")
//	if err != nil && err != custom_error.UserNotFound {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	//save info
//	bookingId, err := uuid.NewUUID()
//	//find type room id
//	roomTypeInfo := model.RoomType{
//		//TypeRoomCode: reqAddBooking.RoomTypeCode,
//	}
//	roomTypeInfo, err = bookingReceiver.BookingRepo.GetRoomTypeInfo(roomTypeInfo)
//	timeCheckIn, err := time.Parse("2006-01-02", reqAddBooking.CheckInTime)
//	if err != nil {
//		return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//	}
//	timeCheckOut, err := time.Parse("2006-01-02", reqAddBooking.CheckOutTime)
//	if err != nil {
//		return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//	}
//	var numberDays float32
//	numberDays = float32(timeCheckOut.Sub(timeCheckIn).Hours() / 24)
//	var totalCost float32
//	totalCost = (reqAddBooking.CostRatePlan + reqAddBooking.CostRoomType) * float32(reqAddBooking.NumberRoom) * numberDays
//	booking := model.Booking{
//		ID:              bookingId.String(),
//		CustomerID:      customerResult.CustomerID,
//		CheckInTime:     timeCheckIn,
//		CheckOutTime:    timeCheckOut,
//		TotalCost:       totalCost,
//		Tax:             totalCost * 0.1,
//		StatusBookingID: 1,
//		PaymentStatusID: 1,
//		RoomTypeID:      roomTypeInfo.ID,
//		NumberRoom:      reqAddBooking.NumberRoom,
//		RatePlanID:      reqAddBooking.RatePlanID,
//	}
//
//	resultSave, err := bookingReceiver.BookingRepo.SaveBooking(booking)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       resultSave,
//		})
//	}
//	//update booking status here
//
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lưu thành công",
//		Data:       resultSave,
//	})
//}
//
//// HandleGetListHistoryBooking godoc
//// @Summary Get list history booking by customer
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetHistoryBooking true "booking"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /booking/history-bookings [post]
//func (bookingReceiver *BookingController) HandleGetListHistoryBooking(c echo.Context) error {
//	reqGetListHistoryBooking := req.RequestGetHistoryBooking{}
//	if err := c.Bind(&reqGetListHistoryBooking); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.CUSTOMER.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	//find customer_ID by jwt user id
//	//process api here
//	var listBooking []model.Booking
//	listBooking, err := bookingReceiver.BookingRepo.GetListHistoryBooking(reqGetListHistoryBooking.CustomerID)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listBooking,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy lịch sử đặt phòng thành công",
//		Data:       listBooking,
//	})
//}
//
//// HandleGetListBookingByCondition godoc
//// @Summary Get list booking by condition
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetListBooking true "booking"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /booking/bookings-filter [post]
//func (bookingReceiver *BookingController) HandleGetListBookingByCondition(c echo.Context) error {
//	reqGetListBooking := req.RequestGetListBooking{}
//	if err := c.Bind(&reqGetListBooking); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	var listBooking []model.Booking
//	condition := map[string]interface{}{
//		"isGetAll":       "false",
//		"customer_name":  strings.ToLower(reqGetListBooking.CustomerName),
//		"room_type_code": strings.ToLower(reqGetListBooking.RoomTypeCode),
//		"status_booking": reqGetListBooking.StatusBooking,
//		"total_cost":     reqGetListBooking.TotalCost,
//		"status_payment": reqGetListBooking.StatusPayment,
//		"check_in_time":  reqGetListBooking.CheckInTime,
//		"check_out_time": reqGetListBooking.CheckOutTime,
//	}
//	listBooking, err := bookingReceiver.BookingRepo.GetListBookingByCondition(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listBooking,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy đặt phòng thành công",
//		Data:       listBooking,
//	})
//}
//
//// HandleCancelBooking godoc
//// @Summary Cancel booking
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestCancelBooking true "booking"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 403 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /booking/cancel-booking [post]
//func (bookingReceiver *BookingController) HandleCancelBooking(c echo.Context) error {
//	reqCancelBooking := req.RequestCancelBooking{}
//	//binding
//	if err := c.Bind(&reqCancelBooking); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.CUSTOMER.String() || claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	booking := model.Booking{
//		ID:              reqCancelBooking.BookingID,
//		StatusBookingID: 3,
//	}
//	bookingCheckTime, err := bookingReceiver.BookingRepo.GetBookingInfo(booking)
//	if bookingCheckTime.CreatedAt.Add(24 * time.Hour).Before(time.Now()) {
//		return c.JSON(http.StatusForbidden, res.Response{
//			StatusCode: http.StatusForbidden,
//			Message:    "Bạn không thể hủy booking vì quá 24h",
//			Data:       nil,
//		})
//	}
//	bookingResult, err := bookingReceiver.BookingRepo.CancelBooking(booking)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Hủy đặt phòng thành công",
//		Data:       bookingResult,
//	})
//}
//
//// HandleGetAllBooking godoc
//// @Summary Get all booking list
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /booking/bookings [get]
//func (bookingReceiver *BookingController) HandleGetAllBooking(c echo.Context) error {
//	var listBooking []model.Booking
//	condition := map[string]interface{}{
//		"isGetAll": "true",
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	listBooking, err := bookingReceiver.BookingRepo.GetListBookingByCondition(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listBooking,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy đặt phòng thành công",
//		Data:       listBooking,
//	})
//}
//
//// HandleGetAllBookingStatus godoc
//// @Summary Get all booking status list
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /booking/booking-statuses [get]
//func (bookingReceiver *BookingController) HandleGetAllBookingStatus(c echo.Context) error {
//	var listBookingStatus []model.StatusBooking
//	listBookingStatus, err := bookingReceiver.BookingRepo.GetBookingStatusList()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listBookingStatus,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách trạng thái đặt phòng thành công",
//		Data:       listBookingStatus,
//	})
//}
//
//// HandleGetBookingInfo godoc
//// @Summary Get booking info
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetBookingInfo true "booking"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /booking/booking-info [post]
//func (bookingReceiver *BookingController) HandleGetBookingInfo(c echo.Context) error {
//	reqGetBookingInfo := req.RequestGetBookingInfo{}
//	//binding
//	if err := c.Bind(&reqGetBookingInfo); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.CUSTOMER.String() || claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	booking := model.Booking{
//		ID: reqGetBookingInfo.BookingID,
//	}
//	bookingResult, err := bookingReceiver.BookingRepo.GetBookingInfo(booking)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy thông tin đặt phòng thành công",
//		Data:       bookingResult,
//	})
//}
//
//// HandleCheckInBooking godoc
//// @Summary Check in room
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestMultiCheckInBooking true "booking"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /booking/check-in [post]
//func (bookingReceiver *BookingController) HandleCheckInBooking(c echo.Context) error {
//	reqCheckInBooking := req.RequestMultiCheckInBooking{}
//	//binding
//	if err := c.Bind(&reqCheckInBooking); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	var listBookingDetail []model.BookingDetail
//	for _, element := range reqCheckInBooking.ListCheckIn {
//		tempID, _ := uuid.NewUUID()
//		tempBookingDetail := model.BookingDetail{
//			ID:        tempID.String(),
//			BookingID: element.BookingID,
//			RoomCode:  element.RoomCode,
//			//Floor:     element.Floor,
//			//Note:      element.Note,
//		}
//		listBookingDetail = append(listBookingDetail, tempBookingDetail)
//	}
//
//	bookingResult, err := bookingReceiver.BookingRepo.SaveMultipleBookingDetails(listBookingDetail)
//	if err != nil || !bookingResult {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    "Check in thất bại",
//		})
//	}
//	var listRoomBusyStatus []model.RoomBusyStatusDetail
//	for _, element := range reqCheckInBooking.ListCheckIn {
//		temp2ID, err := uuid.NewUUID()
//		timeCheckIn, err := time.Parse("2006-01-02", element.CheckInTime)
//		if err != nil {
//			return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//		}
//		timeCheckOut, err := time.Parse("2006-01-02", element.CheckOutTime)
//		if err != nil {
//			return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//		}
//		tempRoomBusyDetail := model.RoomBusyStatusDetail{
//			ID:                       temp2ID.String(),
//			RoomID:                   element.RoomID,
//			BookingID:                element.BookingID,
//			RoomBusyStatusCategoryID: "2",
//			FromTime:                 timeCheckIn,
//			ToTime:                   timeCheckOut,
//		}
//		listRoomBusyStatus = append(listRoomBusyStatus, tempRoomBusyDetail)
//	}
//
//	_, err = bookingReceiver.BookingRepo.SaveRoomBusyStatusDetail(listRoomBusyStatus)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    "Check in thất bại",
//		})
//	}
//	var tempListRoomID []string
//	for _, element := range reqCheckInBooking.ListCheckIn {
//		tempListRoomID = append(tempListRoomID, fmt.Sprintf("'%s'", element.RoomID))
//	}
//	_, err = bookingReceiver.BookingRepo.UpdateRoomByIDs(tempListRoomID)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    "Check in thất bại",
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Check in thành công",
//	})
//}
//
//// HandleGetBookingCheckIn godoc
//// @Summary Get list booking check in
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /booking/get-bookings-check-in [get]
//func (bookingReceiver *BookingController) HandleGetBookingCheckIn(c echo.Context) error {
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	bookingResult, err := bookingReceiver.BookingRepo.GetListBookingCheckInQuery()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    "Lấy danh sách booking thất bại",
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách booking thành công",
//		Data:       bookingResult,
//	})
//}
