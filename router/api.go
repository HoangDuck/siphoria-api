package router

import (
	"github.com/labstack/echo/v4"
	"hotel-booking-api/controller"
	middleware "hotel-booking-api/middlewares"
	"hotel-booking-api/services"
)

type API struct {
	Echo                   *echo.Echo
	AuthController         controller.AuthController
	UserController         controller.UserController
	LogsHandler            controller.LogsHandler
	RatePlanController     controller.RatePlanController
	AdminController        controller.AdminController
	HotelController        controller.HotelController
	UploadFileService      services.FileUploadService
	NotificationService    services.NotificationService
	NotificationController controller.NotificationController
	RoomController         controller.RoomController
	VoucherController      controller.VoucherController
	PaymentController      controller.PaymentController
}

func (api *API) SetupRouter() {
	api.Echo.GET("", func(context echo.Context) error {
		return context.String(200, "Welcome TLCN K19 Tran Kien Khang & Hoang Huu Duc!")
	})

	request := api.Echo.Group("/api")

	request.Group("/map").GET("/test/:id", controller.GetEmbeddedMap)

	auth := request.Group("/auth")
	authLogin := auth.Group("/login")
	authLogin.POST("/general", api.AuthController.HandleSignIn)
	//authLogin.POST("/gg", api.AuthController.HandleSignInGoogle)
	auth.POST("/signup", api.AuthController.HandleRegister)
	auth.POST("/change-pwd", api.AuthController.HandleChangePassword, middleware.JWTMiddleWare())
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
	user.GET("/payments", api.UserController.HandleGetPayments, middleware.JWTMiddleWare())
	user.DELETE("/carts/:id", api.UserController.HandleDeleteCart, middleware.JWTMiddleWare())
	//user.PATCH("/cancel-payment", api.PaymentController.HandleCancelPayment, middleware.JWTMiddleWare())
	//user.POST("/create-payment-session", api.PaymentController.HandleCreatePayment, middleware.JWTMiddleWare())
	user.GET("/notifications", api.UserController.HandleGetUserNotifications, middleware.JWTMiddleWare())
	user.POST("/payments", api.UserController.HandleCreatePaymentFromCart, middleware.JWTMiddleWare())
	user.PUT("/payments", api.UserController.HandleUpdateStatusPayment, middleware.JWTMiddleWare())
	user.GET("/payments/history", api.UserController.HandleGetPaymentsHistory, middleware.JWTMiddleWare())
	user.GET("/payments/pending-checkin", api.UserController.HandleGetPaymentsHistory, middleware.JWTMiddleWare())

	hotel := request.Group("/hotels")
	hotel.GET("/", api.HotelController.HandleGetHotelPartner, middleware.JWTMiddleWare())
	hotel.POST("/", api.HotelController.HandleCreateHotel, middleware.JWTMiddleWare())
	hotel.GET("/search", api.HotelController.HandleGetHotelSearchMobile)
	hotel.POST("/search", api.HotelController.HandleSearchHotel)
	hotel.GET("/:id", api.HotelController.HandleGetHotelById)
	hotel.PATCH("/:id", api.HotelController.HandleUpdateHotel, middleware.JWTMiddleWare())
	hotel.POST("/:id/photos", api.HotelController.HandleUpdateHotelPhoto, middleware.JWTMiddleWare())
	hotel.PATCH("/:id/business-license", api.HotelController.HandleUpdateHotelBusinessLicense, middleware.JWTMiddleWare())
	hotel.DELETE("/:id/photos", api.HotelController.HandleDeleteHotelBusinessLicense, middleware.JWTMiddleWare())
	hotel.POST("/:id/payout", api.HotelController.HandleSendRequestPaymentHotel, middleware.JWTMiddleWare())
	hotel.GET("/:id/rooms", api.HotelController.HandleGetRoomTypeByHotel, middleware.JWTMiddleWare())

	room := request.Group("/rooms")
	room.GET("/welcome-room", func(context echo.Context) error {
		return context.String(200, "Welcome TLCN K19 Tran Kien Khang & Hoang Huu Duc!")
	})
	room.POST("/", api.RoomController.HandleSaveRoomType, middleware.JWTMiddleWare())
	room.GET("/search/:id", api.RoomController.HandleGetHotelSearchById)
	room.POST("/roomnights", api.RoomController.HandleUpdateRoomNight, middleware.JWTMiddleWare())
	room.POST("/ratepackages", api.RoomController.HandleUpdateRatePackages, middleware.JWTMiddleWare())
	room.PATCH("/:id", api.RoomController.HandleUpdateRoomType, middleware.JWTMiddleWare())
	room.PATCH("/:id/photos", api.RoomController.HandleUpdateRoomPhotos, middleware.JWTMiddleWare())
	room.GET("/:id", api.RoomController.HandleGetRoomTypeDetail, middleware.JWTMiddleWare())

	ratePlan := request.Group("/rateplans")
	ratePlan.POST("/", api.RatePlanController.HandleSaveRatePlan, middleware.JWTMiddleWare())
	ratePlan.POST("/rateplan-info", api.RatePlanController.HandleGetRatePlanInfo)
	ratePlan.POST("/:rate_plan_id", api.RatePlanController.HandleUpdateRatePlan, middleware.JWTMiddleWare())

	booking := request.Group("/booking")
	booking.GET("/welcome-booking", func(context echo.Context) error {
		return context.String(200, "Welcome TLCN K19 Tran Kien Khang & Hoang Huu Duc!")
	})

	booking.GET("/welcome-payment", func(context echo.Context) error {
		return context.String(200, "Welcome TLCN K19 Tran Kien Khang & Hoang Huu Duc!")
	})

	payment := request.Group("/payments")
	payment.GET("/test/momo", api.PaymentController.CreatePaymentWithMomo)

	admin := request.Group("/admin")
	admin.POST("/create-account", api.AdminController.HandleCreateAccount, middleware.JWTMiddleWare())
	admin.PATCH("/update-account", api.AdminController.HandleUpdateAccount, middleware.JWTMiddleWare())
	admin.GET("/accounts", api.AdminController.HandleGetAccountByAdmin, middleware.JWTMiddleWare())
	admin.GET("/hotels", api.AdminController.HandleGetHotelByAdmin, middleware.JWTMiddleWare())
	admin.PATCH("/accept/:id", api.AdminController.HandleAcceptHotel, middleware.JWTMiddleWare())
	admin.PATCH("/update-rating/:id", api.AdminController.HandleUpdateRatingHotel, middleware.JWTMiddleWare())
	admin.PATCH("/update-cmsrate/:id", api.AdminController.HandleUpdateCommissionRateHotel, middleware.JWTMiddleWare())
	admin.PATCH("/payouts/:id", api.AdminController.HandleApprovePayoutHotel, middleware.JWTMiddleWare())
	admin.GET("/works/:id", api.AdminController.HandleGetAccountByAdmin, middleware.JWTMiddleWare())
	admin.DELETE("/works", api.AdminController.HandleDeleteHotelWorkByEmployee, middleware.JWTMiddleWare())
	admin.POST("/works", api.AdminController.HandleSaveHotelWorkByEmployee, middleware.JWTMiddleWare())
	log := api.Echo.Group("/manager/log")
	log.GET("/checkViewLogs", api.LogsHandler.CheckLogs)
	log.GET("/checkViewLogsSystem", api.LogsHandler.CheckLogsSystem)

	file := api.Echo.Group("/upload")
	file.POST("/files", api.UploadFileService.UploadMultipleFilesAPI)

	notification := api.Echo.Group("/notification")
	notification.POST("/push-test", api.NotificationService.PushNotificationMessageAPITest)
	admin.POST("/push-notification", api.NotificationController.PushNotificationMessageAPIAdmin, middleware.JWTMiddleWare())

	voucher := request.Group("/vouchers")
	voucher.POST("/", api.VoucherController.HandleSaveVoucher, middleware.JWTMiddleWare())
	voucher.PATCH("/:id", api.VoucherController.HandleUpdateVoucher, middleware.JWTMiddleWare())
}
