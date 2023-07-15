package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gopkg.in/robfig/cron.v2"
	"hotel-booking-api/config"
	"hotel-booking-api/controller"
	"hotel-booking-api/db"
	_ "hotel-booking-api/docs"
	"hotel-booking-api/model"
	"hotel-booking-api/repository/repo_impl"
	"hotel-booking-api/router"
	"hotel-booking-api/services"
	"hotel-booking-api/validator"
	"os"
)

var env string
var ConfigInfo model.Config

func init() {
	fmt.Println(">>>>", os.Getenv("APP_NAME"))
	err := os.Setenv("APP_NAME", "github")
	if err != nil {
		return
	}
	//set variable from cmd and check dev or production
	flag.StringVar(&env, "env", "pro", "env = pro | dev")
	flag.Parse()
	//load config info from file env.dev.yml or env.pro.yml
	config.LoadConfig(&ConfigInfo, &env)
	config.SetupEnv(&ConfigInfo)
	//set configure data to configInfo variable - global variable
	services.ConfigInfo = &ConfigInfo
}

// @title HOTEL BOOKING API
// @version 1.0
// @description This is a sample server Siphoria server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host https://siphoria-api-production.up.railway.app
// @BasePath /api
func main() {
	//initialize firebase instance
	services.InitFirebase()
	//initialize database instance
	sql := new(db.Sql)
	//connect service to database
	sql.Connect(&ConfigInfo)
	//if program run in "dbdev" environment, it runs migration to create, modify tables by definition of models
	//if program doesn't run in "dbdev" environment, it skips running migration
	boolEnableMigrate := env == "dbdev"
	if boolEnableMigrate {
		sql.SetupDB()
	}
	//initialize echo framework instance
	echoInstance := echo.New()
	//initialize swagger instance
	echoInstance.GET("/swagger/*", echoSwagger.WrapHandler)
	//Add validation, middlewares to program
	ConfigureMiddlewaresAndValidator(echoInstance)

	//initialize routers
	var api *router.API
	//initialize controllers
	InitController(api, sql, echoInstance)
	//init go cron (scheduler)
	InitSchedulerGoCron(sql, echoInstance)

	//log export port running app
	echoInstance.Logger.Fatal(echoInstance.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func InitSchedulerGoCron(sql *db.Sql, echoInstance *echo.Echo) {
	//initialize go cron instance (scheduler)
	cronInstance := cron.New()
	schedulerGoCronService := services.SchedulerGoCronService{
		Echo:     echoInstance,
		Cron:     cronInstance,
		RoomRepo: repo_impl.NewRoomRepo(sql),
	}
	schedulerGoCronService.InitSchedulerGoCronLockRoom()
}

func InitController(api *router.API, sql *db.Sql, echoInstance *echo.Echo) {
	//initialize authentication controller with account repository, feature: login, sign up ,...
	accountController := controller.AuthController{
		AccountRepo: repo_impl.NewAccountRepo(sql),
	}
	//initialize user controller with user repository, payment repository to get user, payment data.
	//feature: add to cart, edit profile,...
	userController := controller.UserController{
		UserRepo:    repo_impl.NewUserRepo(sql),
		PaymentRepo: repo_impl.NewPaymentRepo(sql),
	}
	//initialize logger handler
	//feature: view logger system
	logsHandler := controller.LogsHandler{}
	//initialize Administrator controller with Admin repository
	adminController := controller.AdminController{
		AdminRepo: repo_impl.NewAdminRepo(sql),
	}
	//initialize notification controller with notifications repository
	//feature: get user's notifications data
	notificationController := controller.NotificationController{
		NotificationRepo: repo_impl.NewNotificationRepo(sql),
	}
	//initialize hotel controller with hotel repository
	//feature: get data information details hotel
	hotelController := controller.HotelController{
		HotelRepo:   repo_impl.NewHotelRepo(sql),
		RoomRepo:    repo_impl.NewRoomRepo(sql),
		VoucherRepo: repo_impl.NewVoucherRepo(sql),
	}
	//initialize file upload service, using cloudinary
	//feature: upload images.
	uploadService := services.FileUploadService{
		Echo: echoInstance,
	}
	//initialize push notification service, using firebase
	//feature: push notifications, api push notifications to Firebase --> Firebase push to Frontend.
	pushNotificationService := services.NotificationService{
		Echo: echoInstance,
	}
	//initialize room controller with room repository
	roomController := controller.RoomController{
		RoomRepo: repo_impl.NewRoomRepo(sql),
	}
	//initialize rateplan controller with rateplan repository
	//rateplan is a package of services, using with booking a roomtype
	//For example: You book roomtype A with rateplan A (free buffet, refundable)
	//You book roomtype A with rateplan A (not refundable)
	ratePlanController := controller.RatePlanController{
		RatePlanRepo: repo_impl.NewRatePlanRepo(sql),
	}
	//initialize voucher controller with voucher repository
	voucherController := controller.VoucherController{
		VoucherRepo: repo_impl.NewVoucherRepo(sql),
	}
	//initialize payment controller with payment repository
	//a payment is a hotel booking
	paymentController := controller.PaymentController{
		PaymentRepo: repo_impl.NewPaymentRepo(sql),
	}
	//api variables
	api = &router.API{
		Echo:                   echoInstance,
		AuthController:         accountController,
		UserController:         userController,
		LogsHandler:            logsHandler,
		AdminController:        adminController,
		HotelController:        hotelController,
		UploadFileService:      uploadService,
		NotificationService:    pushNotificationService,
		NotificationController: notificationController,
		RoomController:         roomController,
		RatePlanController:     ratePlanController,
		VoucherController:      voucherController,
		PaymentController:      paymentController,
	}
	//set up routers with functions in controllers
	api.SetupRouter()
}

func ConfigureMiddlewaresAndValidator(echoInstance *echo.Echo) {
	//init validation for app
	structValidator := validator.NewStructValidator()
	structValidator.CustomValidate()
	echoInstance.Validator = structValidator
	// middleware accept cors
	echoInstance.Use(middleware.CORS())
	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// middleware loggers
	echoInstance.Use(middleware.Logger())
}
