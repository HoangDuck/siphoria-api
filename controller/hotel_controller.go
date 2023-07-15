package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/model/req"
	"hotel-booking-api/model/res"
	"hotel-booking-api/repository"
	"hotel-booking-api/security"
	"hotel-booking-api/services"
	"hotel-booking-api/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HotelController struct {
	HotelRepo repository.HotelRepo
	RoomRepo  repository.RoomRepo
}

// HandleGetRoomTypeByHotel godoc
// @Summary Get room type by hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id/rooms [get]
func (hotelController *HotelController) HandleGetRoomTypeByHotel(c echo.Context) error {
	var listRoomType []model.RoomType
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) ||
		security.CheckRole(claims, model.STAFF, false) ||
		security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"hotel", "created_at", "updated_at", "", "activated", "description", "max_adult", "max_children", "bed_nums", "bathroom_nums", "photos", "hotel_id", "-", "room_type_facility", "room_nights", "rate_plans", "room_type_views",
	}, &model.RoomType{})
	dataQueryModel.UserId = claims.UserId
	dataQueryModel.DataId = c.Param("id")
	listRoomType, err := hotelController.HotelRepo.GetRoomTypeFilter(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listRoomType)
	}
	return response.Ok(c, "Lấy danh sách phòng thành công", struct {
		Data   []model.RoomType `json:"data"`
		Paging res.PagingModel  `json:"paging"`
	}{
		Data: listRoomType,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
}

// HandleSearchHotel godoc
// @Summary Search Hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateProfile true "hotel"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/search [get]
func (hotelController *HotelController) HandleSearchHotel(c echo.Context) error {
	isCityParamValid := c.QueryParam("city") == ""
	if isCityParamValid {
		return response.BadRequest(c, "Invalid city data", nil)
	}
	isFromDayParamValid := c.QueryParam("from") == ""
	if isFromDayParamValid {
		return response.BadRequest(c, "Invalid from date data", nil)
	} else {
		_, err := time.Parse("2006-01-02", c.QueryParam("from"))
		if err != nil {
			return response.BadRequest(c, "Invalid from format date data", err.Error())
		}
	}
	isToDayParamValid := c.QueryParam("to") == ""
	if isToDayParamValid {
		return response.BadRequest(c, "Invalid to date data", nil)
	} else {
		_, err := time.Parse("2006-01-02", c.QueryParam("to"))
		if err != nil {
			return response.BadRequest(c, "Invalid to format date data", nil)
		}
	}
	listHotel, err := hotelController.HotelRepo.GetListHotelSearch(c)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listHotel)
	}
	return response.Ok(c, "Lấy danh sách khách sạn thành công", listHotel)
}

// HandleGetHotelById godoc
// @Summary Get hotel by Id
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id [get]
func (hotelController *HotelController) HandleGetHotelById(c echo.Context) error {

	hotel, err := hotelController.HotelRepo.GetHotelById(c)
	if err != nil {
		return response.InternalServerError(c, "Lấy room type thất bại", nil)
	}
	roomType := model.RoomType{
		HotelId: c.Param("id"),
	}
	roomTypeList, err := hotelController.RoomRepo.GetListRoomTypeDetail(roomType)
	for i := 0; i < len(roomTypeList); i++ {
		roomTypeItemFacility, _ := hotelController.RoomRepo.GetRoomTypeFacility(roomTypeList[i].ID)
		roomTypeList[i].RoomTypeFacility = &roomTypeItemFacility
		roomTypeItemViews, _ := hotelController.RoomRepo.GetRoomTypeViews(roomTypeList[i].ID)
		roomTypeList[i].RoomTypeViews = &roomTypeItemViews
		roomNights, _ := hotelController.RoomRepo.GetRoomNightsByRoomType(c, roomTypeList[i].ID)
		roomTypeList[i].RoomNights = roomNights
		ratePlans, _ := hotelController.RoomRepo.GetListRatePlans(c, roomTypeList[i].ID)
		roomTypeList[i].RatePlans = ratePlans
		for j := 0; j < len(roomTypeList[i].RatePlans); j++ {
			ratePackages, _ := hotelController.RoomRepo.GetListRatePackages(c, roomTypeList[i].RatePlans[j].ID)
			roomTypeList[i].RatePlans[j].RatePackages = ratePackages
		}
	}

	if err != nil {
		return response.InternalServerError(c, "Lấy room type thất bại", nil)
	}
	hotel.RoomTypes = roomTypeList
	listRoomTypeJson := []res.RoomType{}
	for i := 0; i < len(hotel.RoomTypes); i++ {
		listRoomNightsJson := []res.RoomNight{}
		for x := 0; x < len(hotel.RoomTypes[i].RoomNights); x++ {
			count, err := hotelController.RoomRepo.CountLockRoomByAvailabilityDay(hotel.RoomTypes[i].ID, hotel.RoomTypes[i].RoomNights[x].ID)
			if err != nil {
				count = 0
			}
			tempRoomNightModel := res.RoomNight{
				ID:             hotel.RoomTypes[i].RoomNights[x].ID,
				Inventory:      hotel.RoomTypes[i].RoomNights[x].Inventory,
				Remain:         hotel.RoomTypes[i].RoomNights[x].Remain - count,
				AvailabilityAt: hotel.RoomTypes[i].RoomNights[x].AvailabilityAt.String(),
				UpdatedAt:      hotel.RoomTypes[i].RoomNights[x].UpdatedAt.Unix(),
			}
			listRoomNightsJson = append(listRoomNightsJson, tempRoomNightModel)
		}
		listRatePlanJson := []res.RatePlan{}
		for j := 0; j < len(hotel.RoomTypes[i].RatePlans); j++ {
			listRatePackagesJson := []res.RatePackage{}
			for k := 0; k < len(hotel.RoomTypes[i].RatePlans[j].RatePackages); k++ {
				tempRatePackageModel := res.RatePackage{
					AvailableAt: hotel.RoomTypes[i].RatePlans[j].RatePackages[k].AvailabilityAt.String(),
					Currency:    hotel.RoomTypes[i].RatePlans[j].RatePackages[k].Currency,
					Price:       int(hotel.RoomTypes[i].RatePlans[j].RatePackages[k].Price),
					UpdatedAt:   hotel.RoomTypes[i].RatePlans[j].RatePackages[k].UpdatedAt.Unix(),
				}
				listRatePackagesJson = append(listRatePackagesJson, tempRatePackageModel)
			}
			tempType, _ := strconv.Atoi(hotel.RoomTypes[i].RatePlans[j].Type)
			tempRatePlanModel := res.RatePlan{
				ID:            hotel.RoomTypes[i].RatePlans[j].ID,
				Name:          hotel.RoomTypes[i].RatePlans[j].Name,
				Type:          tempType,
				Status:        hotel.RoomTypes[i].RatePlans[j].Status,
				FreeBreakfast: hotel.RoomTypes[i].RatePlans[j].FreeBreakfast,
				FreeLunch:     hotel.RoomTypes[i].RatePlans[j].FreeLunch,
				FreeDinner:    hotel.RoomTypes[i].RatePlans[j].FreeDinner,
				RatePackages:  listRatePackagesJson,
			}
			listRatePlanJson = append(listRatePlanJson, tempRatePlanModel)
		}
		tempRoomTypeModel := res.RoomType{
			Description: hotel.RoomTypes[i].Description,
			Facilities: map[string]bool{
				"air_conditional": hotel.RoomTypes[i].RoomTypeFacility.AirConditional,
				"tivi":            hotel.RoomTypes[i].RoomTypeFacility.Tivi,
				"kitchen":         hotel.RoomTypes[i].RoomTypeFacility.Kitchen,
				"private_pool":    hotel.RoomTypes[i].RoomTypeFacility.PrivatePool,
				"iron":            hotel.RoomTypes[i].RoomTypeFacility.Iron,
				"sofa":            hotel.RoomTypes[i].RoomTypeFacility.Sofa,
				"desk":            hotel.RoomTypes[i].RoomTypeFacility.Desk,
				"soundproof":      hotel.RoomTypes[i].RoomTypeFacility.Soundproof,
				"towels":          hotel.RoomTypes[i].RoomTypeFacility.Towels,
				"toiletries":      hotel.RoomTypes[i].RoomTypeFacility.Toiletries,
				"fruit":           hotel.RoomTypes[i].RoomTypeFacility.Fruit,
				"shower":          hotel.RoomTypes[i].RoomTypeFacility.Shower,
				"slippers":        hotel.RoomTypes[i].RoomTypeFacility.Slippers,
				"hairdry":         hotel.RoomTypes[i].RoomTypeFacility.Hairdry,
				"bbq":             hotel.RoomTypes[i].RoomTypeFacility.Bbq,
				"wine":            hotel.RoomTypes[i].RoomTypeFacility.Wine,
				"fryer":           hotel.RoomTypes[i].RoomTypeFacility.Fryer,
				"kitchen_tool":    hotel.RoomTypes[i].RoomTypeFacility.KitchenTool,
			}, //json.Marshal(hotel.RoomTypes[i].RoomTypeFacility),
			ID:           hotel.RoomTypes[i].ID,
			MaxAdult:     hotel.RoomTypes[i].MaxAdult,
			MaxChildren:  hotel.RoomTypes[i].MaxChildren,
			BedNums:      hotel.RoomTypes[i].BedNums,
			BathroomNums: hotel.RoomTypes[i].BathroomNums,
			Name:         hotel.RoomTypes[i].Name,
			Photos:       strings.Split(hotel.RoomTypes[i].Photos, ";"),
			RoomNights:   listRoomNightsJson,
			RatePlans:    listRatePlanJson,
			Views: res.Views{
				Beach:          hotel.RoomTypes[i].RoomTypeViews.Beach,
				City:           hotel.RoomTypes[i].RoomTypeViews.City,
				Lake:           hotel.RoomTypes[i].RoomTypeViews.Lake,
				Mountain:       hotel.RoomTypes[i].RoomTypeViews.Mountain,
				PrivateBalcony: hotel.RoomTypes[i].RoomTypeViews.PrivateBalcony,
				Garden:         hotel.RoomTypes[i].RoomTypeViews.Garden,
				River:          hotel.RoomTypes[i].RoomTypeViews.River,
				Bay:            hotel.RoomTypes[i].RoomTypeViews.Bay,
				Sea:            hotel.RoomTypes[i].RoomTypeViews.Sea,
				Ocean:          hotel.RoomTypes[i].RoomTypeViews.Ocean,
			},
		}
		listRoomTypeJson = append(listRoomTypeJson, tempRoomTypeModel)
	}
	return response.Ok(c, "Lấy thông tin thành công", res.HotelDetailModel{
		Activated:    hotel.Activate,
		CityCode:     strconv.Itoa(hotel.ProvinceCode),
		CountryCode:  strconv.Itoa(84),
		DistrictCode: strconv.Itoa(hotel.DistrictCode),
		RawAddress:   hotel.RawAddress,
		Facilities: res.Facilities{
			Beach:         hotel.HotelFacility.Beach,
			Pool:          hotel.HotelFacility.Pool,
			Bar:           hotel.HotelFacility.Bar,
			NoSmokingRoom: hotel.HotelFacility.NoSmokingRoom,
			Fitness:       hotel.HotelFacility.Fitness,
			Spa:           hotel.HotelFacility.Spa,
			Bath:          hotel.HotelFacility.Bath,
			Wifi:          hotel.HotelFacility.Wifi,
			Breakfast:     hotel.HotelFacility.Breakfast,
			Casio:         hotel.HotelFacility.Casino,
			Parking:       hotel.HotelFacility.Parking,
		},
		ID:          hotel.ID,
		CreatedAt:   hotel.CreatedAt.Unix(),
		Name:        hotel.Name,
		Overview:    hotel.Overview,
		HotelPhotos: strings.Split(hotel.HotelPhotos, ";"),
		Rating:      int(hotel.Rating),
		RoomTypes:   listRoomTypeJson,
		UpdatedAt:   hotel.UpdatedAt.Unix(),
	})
}

// HandleGetHotelPartner godoc
// @Summary Get hotel Controller
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels [get]
func (hotelController *HotelController) HandleGetHotelPartner(c echo.Context) error {
	var listHotel []model.Hotel
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) ||
		security.CheckRole(claims, model.STAFF, false) ||
		security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"hotelier", "created_at", "updated_at", "", "overview", "rating", "commission_rate", "status", "activate", "province_code", "district_code", "ward_core", "raw_address", "hotel_photos", "bank_account", "bank_beneficiary", "bank_name", "business_licence", "hotelier_id", "price_hotel", "discount_price", "discount_hotel", "hotel_type", "hotel_facility",
	}, &model.Hotel{})
	dataQueryModel.UserId = claims.UserId
	listHotel, err := hotelController.HotelRepo.GetHotelFilter(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listHotel)
	}
	return response.Ok(c, "Lấy danh sách khách sạn thành công", struct {
		Data   []model.Hotel   `json:"data"`
		Paging res.PagingModel `json:"paging"`
	}{
		Data: listHotel,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
}

// HandleGetHotelSearchMobile godoc
// @Summary Get hotel mobile
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/search [get]
func (hotelController *HotelController) HandleGetHotelSearchMobile(c echo.Context) error {
	var listHotel []model.Hotel
	listHotel, err := hotelController.HotelRepo.GetHotelMobile()
	if err != nil {
		return response.InternalServerError(c, err.Error(), listHotel)
	}
	return c.JSON(http.StatusOK, listHotel)
}

// HandleCreateHotel godoc
// @Summary Create hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestCreateHotel true "hotel"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /hotels [post]
func (hotelController *HotelController) HandleCreateHotel(c echo.Context) error {
	reqCreateHotel := req.RequestCreateHotel{}
	//binding
	if err := c.Bind(&reqCreateHotel); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	hotelId, err := utils.GetNewId()
	if err != nil {
		return response.Forbidden(c, "Đăng ký thất bại", nil)
	}
	reqCreateHotel.ID = hotelId
	result, err := hotelController.HotelRepo.SaveHotel(reqCreateHotel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Đăng ký khách sạn thành công", result)
}

// HandleUpdateHotelPhoto godoc
// @Summary Update hotel photos
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id/photos [post]
func (hotelController *HotelController) HandleUpdateHotelPhoto(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	form, err := c.MultipartForm()
	if err != nil {
		logger.Error("Error create multipart form", zap.Error(err))
	}
	var oldUrls []string
	for i := 0; i < len(form.Value["text"]); i++ {
		oldUrls = append(oldUrls, form.Value["text"][i])
		logger.Error(form.Value["text"][i])
	}

	urls := services.UploadMultipleFiles(c)
	//if len(urls) == 0 {
	//	logger.Error("Error upload avatar to cloudinary failed", zap.Error(nil))
	//	return response.InternalServerError(c, "Cập nhật hình ảnh thất bại", nil)
	//}
	urls = append(urls, oldUrls...)
	//find customer id by userid(account id)
	hotel := model.Hotel{
		ID:          c.Param("id"),
		HotelPhotos: strings.Join(urls, ";"),
	}
	hotel, err = hotelController.HotelRepo.UpdateHotelPhotos(hotel)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật avatar thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", hotel)
}

// HandleUpdateHotelBusinessLicense godoc
// @Summary Update hotel business license photos
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id/business-license [post]
func (hotelController *HotelController) HandleUpdateHotelBusinessLicense(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	form, err := c.MultipartForm()
	if err != nil {
		logger.Error("Error create multipart form", zap.Error(err))
	}
	var oldUrls []string
	for i := 0; i < len(form.Value["text"]); i++ {
		oldUrls = append(oldUrls, form.Value["text"][i])
		logger.Error(form.Value["text"][i])
	}
	urls := services.UploadMultipleFiles(c)
	//if len(urls) == 0 {
	//	logger.Error("Error upload avatar to cloudinary failed", zap.Error(nil))
	//	return response.InternalServerError(c, "Cập nhật hình ảnh thất bại", nil)
	//}
	urls = append(urls, oldUrls...)
	//find customer id by userid(account id)
	hotel := model.Hotel{
		ID:              c.Param("id"),
		BusinessLicence: strings.Join(urls, ";"),
	}
	hotel, err = hotelController.HotelRepo.UpdateHotelBusinessLicensePhotos(hotel)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật giấy phép kinh doanh thất bại", nil)
	}
	return response.Ok(c, "Cập nhật giấy phép kinh doanh thành công", hotel)
}

// HandleDeleteHotelBusinessLicense godoc
// @Summary Delete hotel photos
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id/photos [delete]
func (hotelController *HotelController) HandleDeleteHotelBusinessLicense(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	urls := services.UploadMultipleFiles(c)
	//if len(urls) == 0 {
	//	logger.Error("Error upload avatar to cloudinary failed", zap.Error(nil))
	//	return response.InternalServerError(c, "Cập nhật hình ảnh thất bại", nil)
	//}
	//find customer id by userid(account id)
	hotel := model.Hotel{
		ID:          c.Param("id"),
		HotelPhotos: strings.Join(urls, ";"),
	}
	hotel, err := hotelController.HotelRepo.UpdateHotelBusinessLicensePhotos(hotel)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật avatar thất bại", nil)
	}
	return response.Ok(c, "Cập nhật thành công", hotel)
}

// HandleSendRequestPaymentHotel godoc
// @Summary Send request payment hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id/payout [post]
func (hotelController *HotelController) HandleSendRequestPaymentHotel(c echo.Context) error {
	reqCreatePayout := req.RequestCreatePayout{}
	//binding
	if err := c.Bind(&reqCreatePayout); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, true) || security.CheckRole(claims, model.MANAGER, true)) {
		logger.Error("Error role access", zap.Error(nil))
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	payoutRequestId, err := utils.GetNewId()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	listPaymentId := utils.DecodeJSONArray(reqCreatePayout.Payments, true)

	payoutRequest := model.PayoutRequest{
		ID:           payoutRequestId,
		HotelId:      c.Param("id"),
		PettionerId:  claims.UserId,
		TotalRequest: len(listPaymentId),
		OpenAt:       time.Now(),
		Resolve:      false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	payoutRequestResult, err := hotelController.HotelRepo.CreateRequestPayout(payoutRequest, listPaymentId)

	if err != nil {
		logger.Error("Error uuid data", zap.Error(err))
		return response.InternalServerError(c, "Tạo yêu cầu thanh toán thất bại", nil)
	}
	return response.Ok(c, "Tạo yêu cầu thanh toán thành công", payoutRequestResult)
}

// HandleUpdateHotel godoc
// @Summary update hotel info
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateHotel true "hotel"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /hotel/:id [patch]
func (hotelController *HotelController) HandleUpdateHotel(c echo.Context) error {
	reqUpdateHotel := req.RequestUpdateHotel{}
	if err := c.Bind(&reqUpdateHotel); err != nil {
		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	hotel, err := hotelController.HotelRepo.UpdateHotel(reqUpdateHotel, c.Param("id"))
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật thông tin khách sạn thành công", hotel)
}

// HandleGetPayoutRequestByHotel godoc
// @Summary Get payout request by hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /hotels/:id [get]
func (hotelController *HotelController) HandleGetPayoutRequestByHotel(c echo.Context) error {
	var listPayoutRequest []model.PayoutRequest
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) ||
		security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"pettioner", "hotel", "payer", "", "created_at", "updated_at", "-", "open_at", "close_at",
	}, &model.Hotel{})
	dataQueryModel.UserId = claims.UserId
	listPayoutRequest, err := hotelController.HotelRepo.GetPayoutRequestByHotel(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listPayoutRequest)
	}
	return response.Ok(c, "Lấy danh sách yêu cầu thanh toán thành công", struct {
		Data   []model.PayoutRequest `json:"data"`
		Paging res.PagingModel       `json:"paging"`
	}{
		Data: listPayoutRequest,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
}

// HandleGetReviewByHotel godoc
// @Summary Get reviews by hotel
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /reviews/:id [get]
func (hotelController *HotelController) HandleGetReviewByHotel(c echo.Context) error {
	var listReview []model.Review
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"hotel_id", "hotel", "user_id", "", "created_at", "updated_at", "-", "", "close_at",
	}, &model.Review{})
	dataQueryModel.DataId = c.Param("id")
	listReview, err := hotelController.HotelRepo.GetReviewsByHotel(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listReview)
	}
	return response.Ok(c, "Lấy danh sách yêu cầu thanh toán thành công", struct {
		Data   []model.Review  `json:"data"`
		Paging res.PagingModel `json:"paging"`
	}{
		Data: listReview,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
}

func (hotelController *HotelController) HandleGetVouchersByHotel(c echo.Context) error {
	var listVoucher []model.Voucher
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"hotel_id", "hotel", "id", "", "created_at", "updated_at", "-", "excepts", "is_deleted",
	}, &model.Voucher{})
	dataQueryModel.DataId = c.Param("id")
	listVoucher, err := hotelController.HotelRepo.GetVoucherByHotelFilter(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listVoucher)
	}
	return response.Ok(c, "Lấy danh sách yêu cầu voucher thành công", struct {
		Data   []model.Voucher `json:"data"`
		Paging res.PagingModel `json:"paging"`
	}{
		Data: listVoucher,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
}
