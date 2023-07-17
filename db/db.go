package db

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
)

type Sql struct {
	Db       *gorm.DB
	UserName string
	Password string
	DbName   string
}

func (s *Sql) Connect(config *model.Config) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Database.DbHost,
		config.Database.DbPort,
		config.Database.DbUserName,
		config.Database.DbPassword,
		config.Database.DbName)

	var err error
	s.Db, err = gorm.Open(postgres.Open(dataSource), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		defer logger.ViewLoggerDB(s.Db)
	} else {
		fmt.Println("Connected database!")
	}

}

func (s *Sql) SetupDB() {
	err := s.Db.AutoMigrate(
		//model.RoleModel{},
		//model.ConfigurationUrlDefine{},
		//model.User{},
		//model.PaymentStatus{},
		//model.PaymentMethod{},
		//model.StatusUser{},
		//model.Hotel{},
		//model.HotelWork{},
		//model.HotelType{},
		//model.HotelFacility{},
		model.PayoutRequest{},
		//model.PaymentMethod{},
		//model.Payment{},
		//model.RoomType{},
		//model.RoomTypeViews{},
		//model.RoomTypeFacility{},
		//model.RoomNights{},
		//model.LockRoom{},
		//model.Voucher{},
		//model.VoucherExcept{},
		////model.RatePlan{},
		//model.Wallet{},
		//model.WalletTransaction{},
		//model.RatePackage{},
		//model.Review{},
		//model.UserRank{},
		//model.PaymentDetail{},
		//model.Cart{},
		//model.CartDetail{},
		//model.Notification{},
	)
	if err != nil {
		logger.Error("Error migrate DB", zap.Error(err))
	}
}
