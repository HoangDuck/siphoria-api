package router

import (
	"github.com/labstack/echo/v4"
	"hotel-booking-api/controller"
	middleware "hotel-booking-api/middlewares"
	"hotel-booking-api/services"
)

type API struct {
	Echo           *echo.Echo
	AuthController controller.AuthController
	UserController controller.UserController
	LogsHandler    controller.LogsHandler
	//PaymentController   controller.PaymentController
	//RoomController      controller.RoomController
	//BookingController   controller.BookingController
	RatePlanController controller.RatePlanController
	//AdminController     controller.AdminController
	//StaffController     controller.StaffController
	HotelController     controller.HotelController
	UploadFileService   services.FileUploadService
	NotificationService services.NotificationService
}

func (api *API) SetupRouter() {
	api.Echo.GET("", func(context echo.Context) error {
		return context.String(200, "Welcome TLCN K19 Tran Kien Khang & Hoang Huu Duc!")
	})

	request := api.Echo.Group("/api")

	auth := request.Group("/auth")
	authLogin := auth.Group("/login")
	authLogin.POST("/general", api.AuthController.HandleSignIn)
	//authLogin.POST("/gg", api.AuthController.HandleSignInGoogle)
	auth.POST("/signup", api.AuthController.HandleRegister)
	//auth.POST("/change-pwd", api.AuthController.HandleChangePassword, middleware.JWTMiddleWare())
	auth.POST("/forgot", api.AuthController.HandleSendEmailResetPassword)
	auth.POST("/reset", api.AuthController.HandleResetPassword)
	auth.GET("/verifyemail/:code", api.AuthController.HandleActivateAccount)
	//auth.GET("/deactive/by/account-id", api.AuthController.HandleDeactivateAccount, middleware.JWTMiddleWare())

	//auth.GET("/signGoogle", api.AuthController.HandleAuthenticateWithGoogle)
	//auth.GET("/google-callback", api.AuthController.HandleAuthenticateWithGoogleCallBack)
	//auth.POST("/sign-in-google", api.AuthController.HandleSignInGoogleToken, middleware.JWTMiddleWare())
	//auth.POST("/sign-in-oauth-info", api.AuthController.HandleAuthenticationGoogleWithInfo)
	//auth.GET("/send-email", api.AuthController.TestSendEmail)
	//auth.POST("/check-email", api.AuthController.CheckEmailExisted)
	auth.POST("/refresh-token", api.AuthController.HandleGenerateNewAccessToken)

	user := request.Group("/users")
	user.GET("/details", api.UserController.HandleGetCustomerProfileInfo, middleware.JWTMiddleWare())
	user.POST("/update-avatar", api.UserController.HandleUpdateAvatar, middleware.JWTMiddleWare())
	user.GET("/get-rank-available", api.UserController.HandleGetUserRank, middleware.JWTMiddleWare())
	user.POST("/update-rank", api.UserController.HandleUpdateRank, middleware.JWTMiddleWare())
	user.PATCH("/update-profile", api.UserController.HandleUpdateProfile, middleware.JWTMiddleWare())
	user.POST("/add-to-cart", api.UserController.HandleAddToCart, middleware.JWTMiddleWare())
	user.GET("/carts", api.UserController.HandleGetCart, middleware.JWTMiddleWare())
	//user.PATCH("/cancel-payment", api.PaymentController.HandleCancelPayment, middleware.JWTMiddleWare())
	//user.POST("/create-payment-session", api.PaymentController.HandleCreatePayment, middleware.JWTMiddleWare())

	hotel := request.Group("/hotels")
	hotel.POST("/search", api.HotelController.HandleSearchHotel)
	hotel.GET("/:id", api.HotelController.HandleGetHotelById)

	//////OLD API
	room := request.Group("/room")
	room.GET("/welcome-room", func(context echo.Context) error {
		return context.String(200, "Welcome TLCN K19 Tran Kien Khang & Hoang Huu Duc!")
	})
	//room.GET("/roomtypes", api.RoomController.HandleGetListRoomType)
	//room.POST("/room-type-info", api.RoomController.HandleGetRoomTypeInfo)
	//room.GET("/rooms", api.RoomController.HandleGetRoomList, middleware.JWTMiddleWare())
	//room.POST("/delete-room", api.RoomController.HandleDeleteRoomByCode, middleware.JWTMiddleWare())
	//room.POST("/room-info", api.RoomController.HandleGetRoomInfoByCode)
	//room.POST("/save-room", api.RoomController.HandleSaveRoom, middleware.JWTMiddleWare())
	//room.POST("/update-room", api.RoomController.HandleUpdateRoom, middleware.JWTMiddleWare())
	//room.POST("/checkRoomCountValid", api.RoomController.HandleCheckRoomCountAvailable)
	//room.POST("/check-room-valid", api.RoomController.HandleCheckRoomAvailable)
	//room.POST("/rooms-valid-roomcode", api.RoomController.HandleGetRoomListAvailable)
	//room.POST("/rooms-valid", api.RoomController.HandleSearchGetRoomListAvailable)
	//room.POST("/rooms-by-filter", api.RoomController.HandleGetRoomListFilterSearch)
	//room.POST("/add-room-type", api.RoomController.HandleSaveRoomType, middleware.JWTMiddleWare())
	//room.POST("/add-room-status-cate", api.RoomController.HandleSaveRoomBusyStatusCategory, middleware.JWTMiddleWare())
	//room.POST("/search-num-room-valid", api.RoomController.HandleSearchRoomAvailableByCapacityAndTimeCheckNumberRoomV2)
	//room.POST("/get-room-code", api.RoomController.HandleGetRoomListAtReception, middleware.JWTMiddleWare())
	//room.POST("/add-status-detail", api.RoomController.HandleAddRoomStatusDetails, middleware.JWTMiddleWare())
	//room.POST("/update-status-detail", api.RoomController.HandleUpdateStatusRoomDetail, middleware.JWTMiddleWare())
	//room.GET("/get-rooms-checkout", api.RoomController.HandleGetRoomListCheckOut, middleware.JWTMiddleWare())
	//room.POST("/get-room-status", api.RoomController.HandleGetRoomStatusInfo, middleware.JWTMiddleWare())
	//room.POST("/check-out", api.RoomController.HandleCheckOutBooking, middleware.JWTMiddleWare())
	//room.POST("/get-room-rateplan", api.RoomController.HandleGetRoomRatePlanByTypeRoomCode, middleware.JWTMiddleWare())
	//room.POST("/get-room-check-in", api.RoomController.HandleGetRoomCheckIn, middleware.JWTMiddleWare())

	ratePlan := request.Group("/ratePlan")
	ratePlan.GET("/rateplans", api.RatePlanController.HandleGetListRatePlan)
	ratePlan.POST("/rateplan-info", api.RatePlanController.HandleGetRatePlanInfo)
	ratePlan.POST("/update-rateplan", api.RatePlanController.HandleUpdateRatePlan, middleware.JWTMiddleWare())
	ratePlan.POST("/delete-rateplan", api.RatePlanController.HandleDeleteRatePlan, middleware.JWTMiddleWare())
	ratePlan.POST("/save-rateplan", api.RatePlanController.HandleSaveRatePlan, middleware.JWTMiddleWare())

	booking := request.Group("/booking")
	booking.GET("/welcome-booking", func(context echo.Context) error {
		return context.String(200, "Welcome TLCN K19 Tran Kien Khang & Hoang Huu Duc!")
	})
	//booking.POST("/create-booking", api.BookingController.HandleCreateBooking, middleware.JWTMiddleWare())
	//booking.POST("/history-bookings", api.BookingController.HandleGetListHistoryBooking, middleware.JWTMiddleWare())
	//booking.POST("/bookings-filter", api.BookingController.HandleGetListBookingByCondition, middleware.JWTMiddleWare())
	//booking.POST("/cancel-booking", api.BookingController.HandleCancelBooking, middleware.JWTMiddleWare())
	//booking.GET("/bookings", api.BookingController.HandleGetAllBooking, middleware.JWTMiddleWare())
	//booking.GET("/booking-statuses", api.BookingController.HandleGetAllBookingStatus)
	//booking.POST("/booking-info", api.BookingController.HandleGetBookingInfo, middleware.JWTMiddleWare())
	//booking.POST("/check-in", api.BookingController.HandleCheckInBooking, middleware.JWTMiddleWare())
	//booking.GET("/get-bookings-check-in", api.BookingController.HandleGetBookingCheckIn, middleware.JWTMiddleWare())

	//payment := request.Group("/payment")
	booking.GET("/welcome-payment", func(context echo.Context) error {
		return context.String(200, "Welcome TLCN K19 Tran Kien Khang & Hoang Huu Duc!")
	})
	//payment.POST("/create-payment-momo", api.PaymentController.CreatePaymentWithMomo, middleware.JWTMiddleWare())
	//payment.POST("/result-momo", api.PaymentController.GetResultPaymentMomo)
	//payment.POST("/list-payment", api.PaymentController.HandleGetListPaymentCondition, middleware.JWTMiddleWare())
	//payment.POST("/history-payment", api.PaymentController.HandleGetListHistoryPayment, middleware.JWTMiddleWare())
	//payment.POST("/update-payment", api.PaymentController.HandleUpdatePayment, middleware.JWTMiddleWare())
	//payment.POST("/cancel-payment", api.PaymentController.HandleCancelPayment, middleware.JWTMiddleWare())
	//payment.POST("/delete-payment", api.PaymentController.HandleDeletePayment, middleware.JWTMiddleWare())
	//payment.POST("/create-payment-offline", api.PaymentController.HandleCreatePayment, middleware.JWTMiddleWare())
	//payment.POST("/add-payment-method", api.PaymentController.HandleSavePaymentMethod, middleware.JWTMiddleWare())
	//payment.POST("/create-bill", api.PaymentController.HandleCreatePaymentBill, middleware.JWTMiddleWare())
	//payment.GET("/payments", api.PaymentController.HandleGetAllPayments, middleware.JWTMiddleWare())
	//payment.GET("/payment-statuses", api.PaymentController.HandleGetAllPaymentStatus)

	//staff := request.Group("/staff")
	//staff.POST("/sign-in", api.StaffController.HandleSignIn)
	//staff.POST("/change-pwd", api.StaffController.HandleChangePassword, middleware.JWTMiddleWare())
	//staff.POST("/account-info", api.StaffController.HandleGetAccountInfo, middleware.JWTMiddleWare())
	//staff.POST("/update-profile", api.StaffController.HandleUpdateStaffProfile, middleware.JWTMiddleWare())
	//staff.POST("/change-avatar", api.StaffController.HandleUpdateAvatarStaffProfile, middleware.JWTMiddleWare())

	//admin := request.Group("/admin")
	//admin.POST("/sign-in", api.AdminController.HandleSignIn)
	//admin.POST("/account-info", api.AdminController.HandleGetAccountInfo, middleware.JWTMiddleWare())
	//admin.POST("/staff-info", api.AdminController.HandleGetStaffProfileInfo, middleware.JWTMiddleWare())
	//admin.POST("/change-pwd", api.AdminController.HandleChangePassword, middleware.JWTMiddleWare())
	//admin.POST("/create-staff-profile", api.AdminController.HandleSaveStaffProfile, middleware.JWTMiddleWare())
	//admin.POST("/create-staff-account", api.AdminController.HandleCreateStaffAccount, middleware.JWTMiddleWare())
	//admin.POST("/update-staff-account", api.AdminController.HandleUpdateStaffAccount, middleware.JWTMiddleWare())
	//admin.POST("/activate-staff-acc", api.AdminController.HandleActivateAccount, middleware.JWTMiddleWare())
	//admin.POST("/deactivate-staff-acc", api.AdminController.HandleDeactivateAccount, middleware.JWTMiddleWare())
	//admin.POST("/update-staff-profile", api.AdminController.HandleUpdateStaffProfile, middleware.JWTMiddleWare())
	//admin.POST("/change-role", api.AdminController.HandleChangeRoleName, middleware.JWTMiddleWare())
	//admin.POST("/reset-pwd", api.AdminController.HandleResetPassword, middleware.JWTMiddleWare())
	//admin.POST("/add-status-booking", api.AdminController.HandleSaveBookingStatus, middleware.JWTMiddleWare())
	//admin.POST("/add-status-work", api.AdminController.HandleSaveWorkStatus, middleware.JWTMiddleWare())
	//admin.POST("/add-status-account", api.AdminController.HandleSaveAccountStatus, middleware.JWTMiddleWare())
	//admin.POST("/add-status-payment", api.AdminController.HandleSavePaymentStatus, middleware.JWTMiddleWare())
	//admin.GET("/account-statuses", api.AdminController.HandleGetAllAccountStatus, middleware.JWTMiddleWare())
	//admin.GET("/work-statuses", api.AdminController.HandleGetAllWorkStatus, middleware.JWTMiddleWare())
	//admin.POST("/revenue-statistic-day", api.AdminController.HandleStatisticRevenueDay, middleware.JWTMiddleWare())
	//admin.POST("/revenue-statistic-type-code", api.AdminController.HandleStatisticRevenueTypeRoomCode, middleware.JWTMiddleWare())
	//admin.GET("/get-all-staff", api.AdminController.HandleGetAllStaff, middleware.JWTMiddleWare())
	//admin.GET("/get-all-customer", api.AdminController.HandleGetAllCustomer, middleware.JWTMiddleWare())
	//admin.GET("/get-all-account", api.AdminController.HandleGetAllAccount, middleware.JWTMiddleWare())
	//admin.GET("/get-all-staff-account", api.AdminController.HandleGetAllStaffAccount, middleware.JWTMiddleWare())

	log := api.Echo.Group("/manager/log")
	log.GET("/checkViewLogs", api.LogsHandler.CheckLogs)
	log.GET("/checkViewLogsSystem", api.LogsHandler.CheckLogsSystem)

	file := api.Echo.Group("/upload")
	file.POST("/files", api.UploadFileService.UploadMultipleFilesAPI)

	notification := api.Echo.Group("/notification")
	notification.POST("/push", api.NotificationService.PushNotificationMessageAPI)
}
