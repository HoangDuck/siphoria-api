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
	"hotel-booking-api/utils"
	"strconv"
	"strings"
)

type HotelRepoImpl struct {
	sql *db.Sql
}

func (hotelReceiver *HotelRepoImpl) GetListCheckInByHotel(context echo.Context, queryModel *query.DataQueryModel) ([]model.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (hotelReceiver *HotelRepoImpl) GetTotalReviewByHotel(context echo.Context) (res.TotalReviews, error) {
	hotelId := context.QueryParam("id")
	if hotelId == "" {
		return res.TotalReviews{}, nil
	}
	var resultReviews res.TotalReviews
	var resultAvgRating float32
	var totalRatingCount, totalRatingFiveStart, totalRatingFourStart, totalRatingThreeStart, totalRatingTwoStart, totalRatingOneStart int64
	err := hotelReceiver.sql.Db.Model(&model.Review{}).Select("coalesce(AVG(rating),0) as rating_avg").Where("hotel_id = ?", hotelId).Find(&resultAvgRating)
	err = hotelReceiver.sql.Db.Model(&model.Review{}).Where("hotel_id = ?", hotelId).Count(&totalRatingCount)
	err = hotelReceiver.sql.Db.Model(&model.Review{}).Where("hotel_id = ? AND rating = 5", hotelId).Count(&totalRatingFiveStart)
	err = hotelReceiver.sql.Db.Model(&model.Review{}).Where("hotel_id = ? AND rating = 4", hotelId).Count(&totalRatingFourStart)
	err = hotelReceiver.sql.Db.Model(&model.Review{}).Where("hotel_id = ? AND rating = 3", hotelId).Count(&totalRatingThreeStart)
	err = hotelReceiver.sql.Db.Model(&model.Review{}).Where("hotel_id = ? AND rating = 2", hotelId).Count(&totalRatingTwoStart)
	err = hotelReceiver.sql.Db.Model(&model.Review{}).Where("hotel_id = ? AND rating = 1", hotelId).Count(&totalRatingOneStart)
	if err.Error != nil {
		return resultReviews, err.Error
	}
	if totalRatingCount == 0 {
		totalRatingCount = 1
	}
	resultReviews = res.TotalReviews{
		Average:       resultAvgRating,
		OneStarRate:   float32(totalRatingOneStart / totalRatingCount),
		TwoStarRate:   float32(totalRatingTwoStart / totalRatingCount),
		ThreeStarRate: float32(totalRatingThreeStart / totalRatingCount),
		FourStarRate:  float32(totalRatingFourStart / totalRatingCount),
		FiveStarRate:  float32(totalRatingFiveStart / totalRatingCount),
	}
	return resultReviews, nil
}

func (hotelReceiver *HotelRepoImpl) GetVoucherByHotelFilter(queryModel *query.DataQueryModel) ([]model.Voucher, error) {
	var listVoucher []model.Voucher
	err := GenerateQueryGetData(hotelReceiver.sql, queryModel, &model.Voucher{}, queryModel.ListIgnoreColumns)
	err = err.Preload("Hotel").Where("hotel_id = ?", queryModel.DataId)
	var countTotalRows int64
	err.Model(model.Voucher{}).Count(&countTotalRows)
	queryModel.TotalRows = int(countTotalRows)
	err = err.Find(&listVoucher)
	for index := 0; index < len(listVoucher); index++ {
		var tempListVoucherExcept []model.VoucherExcept
		errVoucherExcept := hotelReceiver.sql.Db.Where("voucher_id = ? AND is_deleted = false", listVoucher[index].ID).Find(&tempListVoucherExcept)
		if errVoucherExcept.Error != nil {
			logger.Error("Error get list voucher except url ", zap.Error(err.Error))
			continue
		}
		listVoucher[index].Excepts = tempListVoucherExcept
	}
	if err.Error != nil {
		logger.Error("Error get list voucher url ", zap.Error(err.Error))
		return listVoucher, err.Error
	}
	return listVoucher, nil
}

func (hotelReceiver *HotelRepoImpl) GetHotelById(context echo.Context) (model.Hotel, error) {
	var hotel model.Hotel
	err := hotelReceiver.sql.Db.Preload("HotelFacility").Where("id = ?", context.Param("id")).Find(&hotel)
	if err.Error != nil {
		return hotel, err.Error
	}
	return hotel, nil
}

func (hotelReceiver *HotelRepoImpl) GetReviewsByHotel(queryModel *query.DataQueryModel) ([]model.Review, error) {
	var listReview []model.Review
	err := GenerateQueryGetData(hotelReceiver.sql, queryModel, &model.Review{}, queryModel.ListIgnoreColumns)
	err = err.Preload("User").Preload("User").Where("hotel_id = ?", queryModel.DataId)
	var countTotalRows int64
	err.Model(model.Review{}).Count(&countTotalRows)
	queryModel.TotalRows = int(countTotalRows)
	err = err.Find(&listReview)
	if err.Error != nil {
		logger.Error("Error get list review url ", zap.Error(err.Error))
		return listReview, err.Error
	}
	return listReview, nil
}

func (hotelReceiver *HotelRepoImpl) GetListHotelSearch(context echo.Context) ([]model.HotelSearch, error) {
	var listHotelData []model.HotelSearch
	from := context.QueryParam("from")
	to := context.QueryParam("to")
	city := context.QueryParam("city")
	rating := context.QueryParam("rating")
	min := context.QueryParam("min")
	max := context.QueryParam("max")
	n_o_r := 1
	if context.QueryParam("n_o_r") != "" {
		temp_n_o_r, err := strconv.ParseInt(context.QueryParam("n_o_r"), 10, 32)
		if err != nil {
			temp_n_o_r = 1
		}
		n_o_r = int(temp_n_o_r)
	}
	n_o_a := 1
	if context.QueryParam("n_o_a") != "" {
		temp_n_o_a, err := strconv.ParseInt(context.QueryParam("n_o_a"), 10, 32)
		if err != nil {
			temp_n_o_a = 1
		}
		n_o_a = int(temp_n_o_a)
	}
	n_o_c := 1
	if context.QueryParam("n_o_c") != "" {
		temp_n_o_c, err := strconv.ParseInt(context.QueryParam("n_o_c"), 10, 32)
		if err != nil {
			temp_n_o_c = 1
		}
		n_o_c = int(temp_n_o_c)
	}
	if rating == "" {
		rating = "1,2,3,4,5"
	}
	err := hotelReceiver.sql.Db.Raw("select * from fn_searchhotel(?,?,?,?,?,?,?::text,?,?)", from, to, n_o_r, n_o_a, n_o_c, city, rating, min, max).Scan(&listHotelData)
	if err.Error != nil {
		return listHotelData, err.Error
	}
	return listHotelData, nil
}

func (hotelReceiver *HotelRepoImpl) GetPayoutRequestByHotel(queryModel *query.DataQueryModel) ([]model.PayoutRequest, error) {
	var listPayoutRequest []model.PayoutRequest
	err := GenerateQueryGetData(hotelReceiver.sql, queryModel, &model.RoomType{}, queryModel.ListIgnoreColumns)
	err = err.Preload("Hotel").Preload("User").Preload("User").Where("hotel_id = ?", queryModel.DataId)
	var countTotalRows int64
	err.Model(model.PayoutRequest{}).Count(&countTotalRows)
	queryModel.TotalRows = int(countTotalRows)
	err = err.Find(&listPayoutRequest)
	if err.Error != nil {
		logger.Error("Error get list room type url ", zap.Error(err.Error))
		return listPayoutRequest, err.Error
	}
	return listPayoutRequest, nil
}

func (hotelReceiver *HotelRepoImpl) GetHotelMobile() ([]model.Hotel, error) {
	var listHotel []model.Hotel
	err := hotelReceiver.sql.Db
	err = err.Order("id desc")
	err = err.Find(&listHotel)
	if err.Error != nil {
		logger.Error("Error get list hotel url ", zap.Error(err.Error))
		return listHotel, err.Error
	}
	return listHotel, nil
}

func (hotelReceiver *HotelRepoImpl) GetRoomTypeFilter(queryModel *query.DataQueryModel) ([]model.RoomType, error) {
	var listRoomType []model.RoomType
	err := GenerateQueryGetData(hotelReceiver.sql, queryModel, &model.RoomType{}, queryModel.ListIgnoreColumns)
	err = err.Preload("RoomTypeFacility").Preload("RoomTypeViews").Where("hotel_id = ?", queryModel.DataId).Preload("RoomTypeFacility").Preload("RoomTypeViews")
	var countTotalRows int64
	err.Model(model.Hotel{}).Count(&countTotalRows)
	queryModel.TotalRows = int(countTotalRows)
	err = err.Find(&listRoomType)
	if err.Error != nil {
		logger.Error("Error get list room type url ", zap.Error(err.Error))
		return listRoomType, err.Error
	}
	return listRoomType, nil
}

func (hotelReceiver *HotelRepoImpl) GetHotelFilter(queryModel *query.DataQueryModel) ([]model.Hotel, error) {
	var listHotel []model.Hotel
	//err :=
	err := GenerateQueryGetData(hotelReceiver.sql, queryModel, &model.Hotel{}, queryModel.ListIgnoreColumns)
	err = err.Where("id in (Select hotel_id from hotel_works where user_id = ?)", queryModel.UserId)
	var countTotalRows int64
	err.Model(model.Hotel{}).Count(&countTotalRows)
	queryModel.TotalRows = int(countTotalRows)
	err = err.Find(&listHotel)
	if err.Error != nil {
		logger.Error("Error get list hotel url ", zap.Error(err.Error))
		return listHotel, err.Error
	}
	return listHotel, nil
}

func (hotelReceiver *HotelRepoImpl) UpdateHotel(requestUpdateHotel req.RequestUpdateHotel, idHotel string) (model.Hotel, error) {
	hotel := model.Hotel{
		ID:              idHotel,
		Name:            requestUpdateHotel.Name,
		Overview:        requestUpdateHotel.Overview,
		Activate:        requestUpdateHotel.Activate,
		ProvinceCode:    requestUpdateHotel.Province,
		DistrictCode:    requestUpdateHotel.District,
		WardCode:        requestUpdateHotel.Ward,
		RawAddress:      requestUpdateHotel.RawAddress,
		BankAccount:     requestUpdateHotel.BankAccount,
		BankBeneficiary: requestUpdateHotel.BankBeneficiary,
		BankName:        requestUpdateHotel.BankName,
	}
	hotelType := model.HotelType{
		HotelId:   idHotel,
		Hotel:     requestUpdateHotel.Hotel,
		Apartment: requestUpdateHotel.Apartment,
		Resort:    requestUpdateHotel.Resort,
		Villa:     requestUpdateHotel.Villa,
		Camping:   requestUpdateHotel.Camping,
		Motel:     requestUpdateHotel.Motel,
		Homestay:  requestUpdateHotel.Homestay,
	}
	hotelFacility := model.HotelFacility{
		HotelId:       idHotel,
		Beach:         requestUpdateHotel.Beach,
		Pool:          requestUpdateHotel.Pool,
		Bar:           requestUpdateHotel.Bar,
		NoSmokingRoom: requestUpdateHotel.NoSmokingRoom,
		Fitness:       requestUpdateHotel.Fitness,
		Spa:           requestUpdateHotel.Spa,
		Bath:          requestUpdateHotel.Bath,
		Wifi:          requestUpdateHotel.Wifi,
		Breakfast:     requestUpdateHotel.Breakfast,
		Casino:        requestUpdateHotel.Casio,
		Parking:       requestUpdateHotel.Parking,
	}

	err := hotelReceiver.sql.Db.Model(&hotel).Updates(hotel)
	err = hotelReceiver.sql.Db.Select("activate").Model(&hotel).Updates(hotel)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}

		return hotel, custom_error.HotelNotUpdated
	}
	err = hotelReceiver.sql.Db.Model(&hotelType).Updates(utils.ConvertStructToMap(&hotelType, []string{
		"hotel_id", "created_at", "updated_at", "-",
	}))
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}

		return hotel, custom_error.HotelNotUpdated
	}
	err = hotelReceiver.sql.Db.Model(&hotelFacility).Updates(utils.ConvertStructToMap(&hotelFacility, []string{
		"hotelier", "created_at", "updated_at", "",
	}))
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}

		return hotel, custom_error.HotelNotUpdated
	}
	return hotel, nil
}

func (hotelReceiver *HotelRepoImpl) SaveHotel(requestAddHotel req.RequestCreateHotel) (model.Hotel, error) {
	err := hotelReceiver.sql.Db.Exec("call sp_addhotel(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);",
		requestAddHotel.ID,
		requestAddHotel.HotelierID,
		requestAddHotel.Name,
		requestAddHotel.Overview,
		requestAddHotel.Activate,
		requestAddHotel.Photos,
		requestAddHotel.RawAddress,
		requestAddHotel.BankAccount,
		requestAddHotel.BankName,
		requestAddHotel.BankBeneficiary,
		requestAddHotel.Hotel,
		requestAddHotel.Apartment,
		requestAddHotel.Resort,
		requestAddHotel.Villa,
		requestAddHotel.Camping,
		requestAddHotel.Motel,
		requestAddHotel.Homestay,
		requestAddHotel.Beach,
		requestAddHotel.Spa,
		requestAddHotel.Pool,
		requestAddHotel.Bar,
		requestAddHotel.NoSmokingRoom,
		requestAddHotel.Fitness,
		requestAddHotel.Bath,
		requestAddHotel.Wifi,
		requestAddHotel.Breakfast,
		requestAddHotel.Casino,
		requestAddHotel.Parking,
		requestAddHotel.District,
		requestAddHotel.Province,
		requestAddHotel.Ward,
		requestAddHotel.BusinessLicense)
	var hotel model.Hotel
	if err.Error != nil {
		return hotel, err.Error
	} else {
		hotel = model.Hotel{
			ID:              requestAddHotel.ID,
			Name:            requestAddHotel.Name,
			Overview:        requestAddHotel.Overview,
			Activate:        requestAddHotel.Activate,
			ProvinceCode:    requestAddHotel.Province,
			DistrictCode:    requestAddHotel.District,
			WardCode:        requestAddHotel.Ward,
			RawAddress:      requestAddHotel.RawAddress,
			HotelPhotos:     requestAddHotel.Photos,
			BankAccount:     requestAddHotel.BankAccount,
			BankBeneficiary: requestAddHotel.BankBeneficiary,
			BankName:        requestAddHotel.BankName,
			BusinessLicence: requestAddHotel.BusinessLicense,
			HotelierId:      requestAddHotel.HotelierID,
		}
	}
	return hotel, nil
}

func (hotelReceiver *HotelRepoImpl) CreateRequestPayout(payoutRequest model.PayoutRequest, paymentIds []string) (model.PayoutRequest, error) {
	var result query.ResultTotalPrice
	err := hotelReceiver.sql.Db.Raw("fn_calculateTotalPricePayment(?,?) as total_price", payoutRequest.HotelId, strings.Join(paymentIds, ",")).Scan(&result)
	payoutRequest.TotalPrice = result.Sum
	payoutRequest.PaymentList = strings.Join(paymentIds, ",")
	err = hotelReceiver.sql.Db.Model(&model.Payment{}).Where("id IN ?", paymentIds).Update("payout_request_id", payoutRequest.ID)
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return payoutRequest, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return payoutRequest, err.Error
		}
		return payoutRequest, err.Error
	}
	err = hotelReceiver.sql.Db.Create(payoutRequest)
	if err.Error != nil {
		logger.Error("Error save data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return payoutRequest, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return payoutRequest, err.Error
		}
		return payoutRequest, err.Error
	}
	return payoutRequest, nil
}

func (hotelReceiver *HotelRepoImpl) UpdateHotelBusinessLicensePhotos(hotel model.Hotel) (model.Hotel, error) {
	err := hotelReceiver.sql.Db.Model(&hotel).Select("business_licence").Updates(map[string]interface{}{"business_licence": hotel.BusinessLicence})
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return hotel, err.Error
		}
		return hotel, err.Error
	}
	return hotel, nil
}

func (hotelReceiver *HotelRepoImpl) UpdateHotelPhotos(hotel model.Hotel) (model.Hotel, error) {
	err := hotelReceiver.sql.Db.Model(&hotel).Select("hotel_photos").Updates(map[string]interface{}{"hotel_photos": hotel.HotelPhotos})
	if err.Error != nil {
		logger.Error("Error query data", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}
		if err.Error == gorm.ErrInvalidTransaction {
			return hotel, err.Error
		}
		return hotel, err.Error
	}
	return hotel, nil
}

func NewHotelRepo(sql *db.Sql) repository.HotelRepo {
	return &HotelRepoImpl{
		sql: sql,
	}
}
