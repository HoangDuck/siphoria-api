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
	"strings"
	"time"
)

type HotelController struct {
	HotelRepo repository.HotelRepo
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
// @Router /hotels/search [post]
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
			return response.BadRequest(c, "Invalid from format date data", nil)
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

	return response.Ok(c, "Cập nhật thành công", nil)
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
