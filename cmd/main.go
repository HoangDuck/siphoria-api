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
	//set variable from cmd and check dev, fast or production
	flag.StringVar(&env, "env", "pro", "env = pro | dev")
	flag.Parse()
	//load config info from file
	config.LoadConfig(&ConfigInfo, &env)
	config.SetupEnv(&ConfigInfo)
	services.ConfigInfo = &ConfigInfo
}

// @title HOTEL BOOKING API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host https://siphoria-api-production.up.railway.app
// @BasePath /api
func main() {
	services.InitFirebase()
	//initialize DB
	sql := new(db.Sql)
	sql.Connect(&ConfigInfo)
	boolEnableMigrate := env == "dbdev"
	if boolEnableMigrate {
		sql.SetupDB()
	}

	//init echo framework
	echoInstance := echo.New()
	echoInstance.GET("/swagger/*", echoSwagger.WrapHandler)
	ConfigureMiddlewaresAndValidator(echoInstance)

	//init routers
	var api *router.API
	InitController(api, sql, echoInstance)
	//init go cron (scheduler)
	InitSchedulerGoCron(sql, echoInstance)

	//log export port running app
	echoInstance.Logger.Fatal(echoInstance.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func InitSchedulerGoCron(sql *db.Sql, echoInstance *echo.Echo) {
	//init go cron (scheduler)
	cronInstance := cron.New()
	schedulerGoCronService := services.SchedulerGoCronService{
		Echo:     echoInstance,
		Cron:     cronInstance,
		RoomRepo: repo_impl.NewRoomRepo(sql),
	}
	schedulerGoCronService.InitSchedulerGoCronLockRoom()
}

func InitController(api *router.API, sql *db.Sql, echoInstance *echo.Echo) {
	accountController := controller.AuthController{
		AccountRepo: repo_impl.NewAccountRepo(sql),
	}
	userController := controller.UserController{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	logsHandler := controller.LogsHandler{}

	adminController := controller.AdminController{
		AdminRepo: repo_impl.NewAdminRepo(sql),
	}
	notificationController := controller.NotificationController{
		NotificationRepo: repo_impl.NewNotificationRepo(sql),
	}
	hotelController := controller.HotelController{
		HotelRepo: repo_impl.NewHotelRepo(sql),
		RoomRepo:  repo_impl.NewRoomRepo(sql),
	}

	uploadService := services.FileUploadService{
		Echo: echoInstance,
	}

	pushNotificationService := services.NotificationService{
		Echo: echoInstance,
	}

	roomController := controller.RoomController{
		RoomRepo: repo_impl.NewRoomRepo(sql),
	}

	ratePlanController := controller.RatePlanController{
		RatePlanRepo: repo_impl.NewRatePlanRepo(sql),
	}

	voucherController := controller.VoucherController{
		VoucherRepo: repo_impl.NewVoucherRepo(sql),
	}

	paymentController := controller.PaymentController{
		PaymentRepo: repo_impl.NewPaymentRepo(sql),
	}

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

	api.SetupRouter()
}

func ConfigureMiddlewaresAndValidator(echoInstance *echo.Echo) { //init validation for app
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
