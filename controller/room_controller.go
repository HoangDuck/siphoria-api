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
	"io"
	_ "math/rand"
	"net/http"
	"strings"
)

type RoomController struct {
	RoomRepo repository.RoomRepo
}

// HandleGetHotelSearchById godoc
// @Summary Get search hotel by Id
// @Tags hotel-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /rooms/search/:id [get]
func (roomController *RoomController) HandleGetHotelSearchById(c echo.Context) error {
	roomType := model.RoomType{
		HotelId: c.Param("id"),
	}
	roomTypeList, err := roomController.RoomRepo.GetListRoomTypeDetail(roomType)
	for i := 0; i < len(roomTypeList); i++ {
		roomTypeItemFacility, _ := roomController.RoomRepo.GetRoomTypeFacility(roomTypeList[i].ID)
		roomTypeList[i].RoomTypeFacility = &roomTypeItemFacility
		roomTypeItemViews, _ := roomController.RoomRepo.GetRoomTypeViews(roomTypeList[i].ID)
		roomTypeList[i].RoomTypeViews = &roomTypeItemViews
		roomNights, _ := roomController.RoomRepo.GetRoomNightsByRoomType(c, roomTypeList[i].ID)
		roomTypeList[i].RoomNights = roomNights
		ratePlans, _ := roomController.RoomRepo.GetListRatePlans(c, roomTypeList[i].ID)
		roomTypeList[i].RatePlans = ratePlans
		for j := 0; j < len(roomTypeList[i].RatePlans); j++ {
			ratePackages, _ := roomController.RoomRepo.GetListRatePackages(c, roomTypeList[i].RatePlans[j].ID)
			roomTypeList[i].RatePlans[j].RatePackages = ratePackages
		}
	}
	if err != nil {
		return response.InternalServerError(c, "Lấy room type thất bại", nil)
	}
	return c.JSON(http.StatusOK, roomTypeList) //response.Ok(c, "Cập nhật thành công", roomTypeList)
}

// HandleGetRoomTypeDetail godoc
// @Summary Get room type detail
// @Tags room-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms/:id [get]
func (roomController *RoomController) HandleGetRoomTypeDetail(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) || security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	roomType := model.RoomType{
		ID: c.Param("id"),
	}
	roomType, err := roomController.RoomRepo.GetRoomTypeDetail(roomType)
	if err != nil {
		return response.InternalServerError(c, "Lấy room type thất bại", nil)
	}
	return response.Ok(c, "Lấy room type thành công", roomType)
}

// HandleSaveRoomType godoc
// @Summary Save room type
// @Tags room-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestCreateRoomType true "room"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms [post]
func (roomController *RoomController) HandleSaveRoomType(c echo.Context) error {
	reqAddRoomType := req.RequestCreateRoomType{}
	//binding
	if err := c.Bind(&reqAddRoomType); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	roomTypeId, err := utils.GetNewId()
	if err != nil {
		return response.InternalServerError(c, "Đăng ký thất bại", nil)
	}
	reqAddRoomType.ID = roomTypeId
	result, err := roomController.RoomRepo.SaveRoomType(reqAddRoomType)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Lưu room type thành công", result)
}

// HandleUpdateRoomNight godoc
// @Summary Update room nights
// @Tags room-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateRoomNight true "room"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms/roomnights [post]
func (roomController *RoomController) HandleUpdateRoomNight(c echo.Context) error {
	reqUpdateRoomNight := req.RequestUpdateRoomNight{}
	//binding
	if err := c.Bind(&reqUpdateRoomNight); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	result, err := roomController.RoomRepo.UpdateRoomNight(reqUpdateRoomNight)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật room nights thành công", result)
}

// HandleUpdateRatePackages godoc
// @Summary Update ratepackages
// @Tags room-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateRatePackage true "room"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms/ratepackages [post]
func (roomController *RoomController) HandleUpdateRatePackages(c echo.Context) error {
	reqUpdateRatePackages := req.RequestUpdateRatePackage{}
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	reqUpdateRatePackages, err = req.UnmarshalRequestUpdateRatePackage(body)
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	result, err := roomController.RoomRepo.UpdateRatePackages(reqUpdateRatePackages)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật rate packages thành công", result)
}

// HandleUpdateRoomType godoc
// @Summary update room type
// @Tags room-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateRoomType true "RoomType"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms/:id [patch]
func (roomController *RoomController) HandleUpdateRoomType(c echo.Context) error {
	reqUpdateRoomType := req.RequestUpdateRoomType{}
	if err := c.Bind(&reqUpdateRoomType); err != nil {
		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) || security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	roomType, err := roomController.RoomRepo.UpdateRoomType(reqUpdateRoomType, c.Param("id"))
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Cập nhật thông tin room type thành công", roomType)
}

// HandleUpdateRoomPhotos godoc
// @Summary Update room photos
// @Tags room-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /rooms/:id/photos [patch]
func (roomController *RoomController) HandleUpdateRoomPhotos(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) || security.CheckRole(claims, model.MANAGER, false)) {
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
	room := model.RoomType{
		ID:     c.Param("id"),
		Photos: strings.Join(urls, ";"),
	}
	room, err = roomController.RoomRepo.UpdateRoomPhotos(room)
	if err != nil {
		logger.Error("Error save database", zap.Error(err))
		return response.InternalServerError(c, "Cập nhật hình ảnh thất bại", nil)
	}
	return response.Ok(c, "Cập nhật hình ảnh thành công", room)
}

// HandleGetRatePlanByRoomType godoc
// @Summary Get rateplans list by room type
// @Tags -service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /room/:id/rateplans [get]
func (roomController *RoomController) HandleGetRatePlanByRoomType(c echo.Context) error {
	var listRatePlan []model.RatePlan
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) ||
		security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	dataQueryModel := utils.GetQueryDataModel(c, []string{
		"-", "created_at", "updated_at", "", "rate_packages", "room_type", "prices", "activate", "free_breakfast", "free_lunch", "free_dinner", "room_type_id",
	}, &model.RatePlan{})
	dataQueryModel.UserId = claims.UserId
	dataQueryModel.DataId = c.Param("id")
	listRatePlan, err := roomController.RoomRepo.GetRatePlanByRoomTypeFilter(&dataQueryModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), listRatePlan)
	}

	return response.Ok(c, "Lấy danh sách rate plan thành công", struct {
		Data   []model.RatePlan `json:"data"`
		Paging res.PagingModel  `json:"paging"`
	}{
		Data: listRatePlan,
		Paging: res.PagingModel{
			TotalItems: dataQueryModel.TotalRows,
			TotalPages: dataQueryModel.TotalPages,
			Page:       dataQueryModel.PageViewIndex,
			Offset:     dataQueryModel.Limit,
		},
	})
}

// HandleGetRoomInventories godoc
// @Summary Get room inventories
// @Tags -service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /rooms/:id/inventories [get]
func (roomController *RoomController) HandleGetRoomInventories(c echo.Context) error {
	roomType := model.RoomType{
		ID: c.Param("id"),
	}
	roomNights, err := roomController.RoomRepo.GetRoomNightsByRoomType(c, roomType.ID)
	roomType.RoomNights = roomNights
	ratePlans, err := roomController.RoomRepo.GetListRatePlans(c, roomType.ID)
	roomType.RatePlans = ratePlans
	for j := 0; j < len(roomType.RatePlans); j++ {
		ratePackages, _ := roomController.RoomRepo.GetListRatePackages(c, roomType.RatePlans[j].ID)
		roomType.RatePlans[j].RatePackages = ratePackages
	}
	if err != nil {
		return response.InternalServerError(c, "Lấy room type thất bại", nil)
	}

	var listRoomNightsResponse []res.RoomNightResponse
	for i, _ := range roomType.RoomNights {
		listRoomNightsResponse = append(listRoomNightsResponse, res.RoomNightResponse{
			ID:             roomType.RoomNights[i].ID,
			AvailabilityAt: roomType.RoomNights[i].AvailabilityAt,
			Quantity:       roomType.RoomNights[i].Remain,
		})
	}
	var listRatePlansResponse []res.RatePlanResponse
	for j, _ := range roomType.RatePlans {
		var ratePlanItem = res.RatePlanResponse{
			RateplanID: roomType.RatePlans[j].ID,
		}
		for k, _ := range roomType.RatePlans[j].RatePackages {
			ratePlanItem.Prices = append(ratePlanItem.Prices, res.Price{
				ID:             roomType.RatePlans[j].RatePackages[k].ID,
				AvailabilityAt: roomType.RatePlans[j].RatePackages[k].AvailabilityAt,
				Price:          roomType.RatePlans[j].RatePackages[k].Price,
			})
		}
		listRatePlansResponse = append(listRatePlansResponse, ratePlanItem)
	}

	return response.Ok(c, "Lấy inventory thành công", struct {
		RoomNights []res.RoomNightResponse `json:"roomnight"`
		Rateplans  []res.RatePlanResponse  `json:"rateplans"`
	}{
		RoomNights: listRoomNightsResponse,
		Rateplans:  listRatePlansResponse,
	})
}
