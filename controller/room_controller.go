package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
	"hotel-booking-api/utils"
	_ "math/rand"
)

type RoomController struct {
	RoomRepo repository.RoomRepo
}

//// HandleSaveRoomBusyStatusCategory godoc
//// @Summary Save room busy status category
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddRoomBusyStatusCategory true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/add-room-status-cate [post]
//func (roomReceiver *RoomController) HandleSaveRoomBusyStatusCategory(c echo.Context) error {
//	reqAddRoomBusyStatusCategory := req.RequestAddRoomBusyStatusCategory{}
//	//binding
//	if err := c.Bind(&reqAddRoomBusyStatusCategory); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	roomBusyStatusCategoryId, err := uuid.NewUUID()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	roomBusyStatusCategory := model.RoomBusyStatusCategory{
//		ID:          roomBusyStatusCategoryId.String(),
//		StatusCode:  reqAddRoomBusyStatusCategory.StatusCode,
//		StatusName:  reqAddRoomBusyStatusCategory.StatusName,
//		Description: reqAddRoomBusyStatusCategory.Description,
//	}
//	result, err := roomReceiver.RoomRepo.SaveRoomBusyStatusCategory(roomBusyStatusCategory)
//	if err != nil || !result {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lưu thành công", nil)
//}

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
func (roomReceiver *RoomController) HandleSaveRoomType(c echo.Context) error {
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
		return response.Forbidden(c, "Đăng ký thất bại", nil)
	}
	reqAddRoomType.ID = roomTypeId
	result, err := roomReceiver.RoomRepo.SaveRoomType(reqAddRoomType)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Lưu thành công", result)
}

//// HandleGetListRoomType godoc
//// @Summary Get list room type
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/roomtypes [get]
//func (roomReceiver *RoomController) HandleGetListRoomType(c echo.Context) error {
//	var listRoomType []model.RoomType
//	listRoomType, err := roomReceiver.RoomRepo.GetRoomListType()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoomType,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách loại phòng thành công",
//		Data:       listRoomType,
//	})
//}
//
//// HandleGetRoomTypeInfo godoc
//// @Summary Get room type info
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetRoomTypeInfo true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/room-type-info [post]
//func (roomReceiver *RoomController) HandleGetRoomTypeInfo(c echo.Context) error {
//	reqGetRoomTypeInfo := req.RequestGetRoomTypeInfo{}
//	//binding
//	if err := c.Bind(&reqGetRoomTypeInfo); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	roomTypeModel := model.RoomType{
//		//TypeRoomCode: reqGetRoomTypeInfo.TypeRoomCode,
//	}
//	roomType, err := roomReceiver.RoomRepo.GetRoomTypeInfo(roomTypeModel)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy thông tin loại phòng thành công",
//		Data:       roomType,
//	})
//}
//
//// HandleGetRoomList godoc
//// @Summary Get room list
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/rooms [get]
//func (roomReceiver *RoomController) HandleGetRoomList(c echo.Context) error {
//	var listRoom []model.Room
//	condition := map[string]interface{}{
//		"isGetAll": "true",
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	listRoom, err := roomReceiver.RoomRepo.GetRoomListByCondition(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoom,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listRoom,
//	})
//}
//
//// HandleDeleteRoomByCode godoc
//// @Summary Delete room by room_code
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestDeleteRoom true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/delete-room [post]
//func (roomReceiver *RoomController) HandleDeleteRoomByCode(c echo.Context) error {
//	reqDeleteRoomByCode := req.RequestDeleteRoom{}
//	//binding
//	if err := c.Bind(&reqDeleteRoomByCode); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	result, err := roomReceiver.RoomRepo.DeleteRoomByID(reqDeleteRoomByCode)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	if result != true {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Xóa phòng thành công",
//		Data:       nil,
//	})
//}
//
//// HandleGetRoomInfoByCode godoc
//// @Summary Get room info by room_code
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetRoomInfoByID true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/room-info [post]
//func (roomReceiver *RoomController) HandleGetRoomInfoByCode(c echo.Context) error {
//	reqGetRoomInfo := req.RequestGetRoomInfoByID{}
//	//binding
//	if err := c.Bind(&reqGetRoomInfo); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	room, err := roomReceiver.RoomRepo.GetRoomInfo(reqGetRoomInfo)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy thông tin phòng thành công",
//		Data:       room,
//	})
//}
//
//// HandleSaveRoom godoc
//// @Summary Save room
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddRoom true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/save-room [post]
//func (roomReceiver *RoomController) HandleSaveRoom(c echo.Context) error {
//	reqAddRoom := req.RequestAddRoom{}
//	//binding
//	if err := c.Bind(&reqAddRoom); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	roomTypeQuery := model.RoomType{
//		//TypeRoomCode: reqAddRoom.TypeRoomCode,
//	}
//	roomType, err := roomReceiver.RoomRepo.GetRoomTypeInfo(roomTypeQuery)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//
//	roomId, err := uuid.NewUUID()
//	minNumber := 10
//	maxNumber := 99
//	room := model.Room{
//		ID:         roomId.String(),
//		RoomCode:   fmt.Sprintf("P.%d%d", reqAddRoom.Floor, rand.Intn(maxNumber-minNumber)+minNumber),
//		RoomTypeID: roomType.ID,
//		Floor:      reqAddRoom.Floor,
//	}
//	resultSave, err := roomReceiver.RoomRepo.SaveRoom(room)
//	if err != nil || !resultSave {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lưu thành công", nil)
//}
//
//// HandleUpdateRoom godoc
//// @Summary Update room
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestUpdateRoomInfo true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/update-room [post]
//func (roomReceiver *RoomController) HandleUpdateRoom(c echo.Context) error {
//	reqUpdateRoomInfo := req.RequestUpdateRoomInfo{}
//	if err := c.Bind(&reqUpdateRoomInfo); err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	err := c.Validate(reqUpdateRoomInfo)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	roomTypeModel := model.RoomType{
//		//TypeRoomCode: reqUpdateRoomInfo.TypeRoomCode,
//	}
//	roomType, err := roomReceiver.RoomRepo.GetRoomTypeInfo(roomTypeModel)
//	room := model.Room{
//		RoomCode:   reqUpdateRoomInfo.RoomCode,
//		RoomTypeID: roomType.ID,
//	}
//	room, err = roomReceiver.RoomRepo.UpdateRoomByRoomCode(room)
//	if err != nil {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Cập nhật thông tin phòng thành công",
//		Data:       room,
//	})
//}
//
//// HandleCheckRoomCountAvailable godoc
//// @Summary Check room is available
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestCheckRoomAvailable true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 422 {object} res.Response
//// @Router /room/checkRoomCountValid [post]
//func (roomReceiver *RoomController) HandleCheckRoomCountAvailable(c echo.Context) error {
//	reqCheckRoomAvailable := req.RequestCheckRoomAvailable{}
//	if err := c.Bind(&reqCheckRoomAvailable); err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	err := c.Validate(reqCheckRoomAvailable)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	condition := map[string]interface{}{
//		"type_room_code": strings.ToLower(reqCheckRoomAvailable.TypeRoomCode),
//		"start_time":     reqCheckRoomAvailable.TimeStart,
//		"end_time":       reqCheckRoomAvailable.TimeEnd,
//	}
//
//	result, err := roomReceiver.RoomRepo.CheckRoomAvailable(condition)
//	if err != nil {
//		return c.JSON(http.StatusUnprocessableEntity, res.Response{
//			StatusCode: http.StatusUnprocessableEntity,
//			Message:    err.Error(),
//			Data:       0,
//		})
//	}
//	if len(result) == 0 {
//		return c.JSON(http.StatusOK, res.Response{
//			StatusCode: http.StatusOK,
//			Message:    "Không khả dụng",
//			Data:       0,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Khả dụng",
//		Data:       len(result),
//	})
//}
//
//// HandleCheckRoomAvailable godoc
//// @Summary Check room is available
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestCheckRoomAvailable true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 422 {object} res.Response
//// @Router /room/check-room-valid [post]
//func (roomReceiver *RoomController) HandleCheckRoomAvailable(c echo.Context) error {
//	reqCheckRoomAvailable := req.RequestCheckRoomAvailable{}
//	if err := c.Bind(&reqCheckRoomAvailable); err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	err := c.Validate(reqCheckRoomAvailable)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	condition := map[string]interface{}{
//		"type_room_code": reqCheckRoomAvailable.TypeRoomCode,
//		"start_time":     reqCheckRoomAvailable.TimeStart,
//		"end_time":       reqCheckRoomAvailable.TimeEnd,
//	}
//
//	result, err := roomReceiver.RoomRepo.CheckRoomAvailable(condition)
//	if err != nil {
//		return c.JSON(http.StatusUnprocessableEntity, res.Response{
//			StatusCode: http.StatusUnprocessableEntity,
//			Message:    err.Error(),
//			Data:       result,
//		})
//	}
//	if len(result) == 0 {
//		return c.JSON(http.StatusOK, res.Response{
//			StatusCode: http.StatusOK,
//			Message:    "Không khả dụng",
//			Data:       result,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Khả dụng",
//		Data:       result,
//	})
//}
//
//// HandleGetRoomListAvailable godoc
//// @Summary Get room list available
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetRoomListAvailable true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 422 {object} res.Response
//// @Router /room/rooms-valid-roomcode [post]
//func (roomReceiver *RoomController) HandleGetRoomListAvailable(c echo.Context) error {
//	reqGetRoomListAvailable := req.RequestGetRoomListAvailable{}
//	//binding
//	if err := c.Bind(&reqGetRoomListAvailable); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	var listRoom []model.Room
//	condition := map[string]interface{}{
//		"isGetAll":        "false",
//		"room_type_code":  strings.ToLower(reqGetRoomListAvailable.RoomTypeCode),
//		"number_adult":    reqGetRoomListAvailable.NumberAdult,
//		"number_children": reqGetRoomListAvailable.NumberChildren,
//		"bed_number":      reqGetRoomListAvailable.NumberBed,
//		"number_toilet":   reqGetRoomListAvailable.NumberToilet,
//		"time_start":      reqGetRoomListAvailable.TimeStart,
//		"time_end":        reqGetRoomListAvailable.TimeEnd,
//	}
//
//	listRoom, err := roomReceiver.RoomRepo.GetRoomListByCondition(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoom,
//		})
//	}
//	tempStringIDs := ""
//	for tempIndex := 0; tempIndex < len(listRoom); tempIndex++ {
//		if tempIndex == len(listRoom)-1 {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'"
//		} else {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'" + ","
//		}
//	}
//	condition["room_ids"] = tempStringIDs
//	listRoom, err = roomReceiver.RoomRepo.GetRoomListAvailable(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoom,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listRoom,
//	})
//}
//
//// HandleSearchGetRoomListAvailable godoc
//// @Summary Get room list available
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestSearchGetRoomListAvailable true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 422 {object} res.Response
//// @Router /room/rooms-valid [post]
//func (roomReceiver *RoomController) HandleSearchGetRoomListAvailable(c echo.Context) error {
//	reqGetRoomListAvailable := req.RequestSearchGetRoomListAvailable{}
//	//binding
//	if err := c.Bind(&reqGetRoomListAvailable); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	var listRoom []model.Room
//	condition := map[string]interface{}{
//		"isGetAll":        "false",
//		"room_type_code":  "all",
//		"number_adult":    reqGetRoomListAvailable.NumberAdult,
//		"number_children": reqGetRoomListAvailable.NumberChildren,
//		"bed_number":      reqGetRoomListAvailable.NumberBed,
//		"number_toilet":   reqGetRoomListAvailable.NumberToilet,
//		"time_start":      reqGetRoomListAvailable.TimeStart,
//		"time_end":        reqGetRoomListAvailable.TimeEnd,
//	}
//
//	listRoom, err := roomReceiver.RoomRepo.GetRoomListByCondition(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoom,
//		})
//	}
//	tempStringIDs := ""
//	for tempIndex := 0; tempIndex < len(listRoom); tempIndex++ {
//		if tempIndex == len(listRoom)-1 {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'"
//		} else {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'" + ","
//		}
//	}
//	condition["room_ids"] = tempStringIDs
//	listRoom, err = roomReceiver.RoomRepo.GetRoomListAvailable(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoom,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listRoom,
//	})
//}
//
//func (roomReceiver *RoomController) HandleGetRoomListCountAvailable(c echo.Context) error {
//	reqGetRoomListAvailable := req.RequestSearchGetRoomListAvailable{}
//	//binding
//	if err := c.Bind(&reqGetRoomListAvailable); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	var listRoom []model.Room
//	condition := map[string]interface{}{
//		"isGetAll":        "false",
//		"room_type_code":  "all",
//		"number_adult":    reqGetRoomListAvailable.NumberAdult,
//		"number_children": reqGetRoomListAvailable.NumberChildren,
//		"bed_number":      reqGetRoomListAvailable.NumberBed,
//		"number_toilet":   reqGetRoomListAvailable.NumberToilet,
//		"start_time":      reqGetRoomListAvailable.TimeStart,
//		"end_time":        reqGetRoomListAvailable.TimeEnd,
//	}
//
//	listRoom, err := roomReceiver.RoomRepo.GetRoomListByCondition(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoom,
//		})
//	}
//	tempStringIDs := ""
//	for tempIndex := 0; tempIndex < len(listRoom); tempIndex++ {
//		if tempIndex == len(listRoom)-1 {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'"
//		} else {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'" + ","
//		}
//	}
//	condition["room_ids"] = tempStringIDs
//	listRoomCountGroupByType, err := roomReceiver.RoomRepo.CountAvailableRoomAndGroupByType(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoomCountGroupByType,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listRoomCountGroupByType,
//	})
//}
//
//// HandleGetRoomListFilterSearch godoc
//// @Summary Get room list filter search
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetRoomListFilterSearch true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 422 {object} res.Response
//// @Router /room/rooms-by-filter [post]
//func (roomReceiver *RoomController) HandleGetRoomListFilterSearch(c echo.Context) error {
//	reqGetRoomList := req.RequestGetRoomListFilterSearch{}
//	//binding
//	if err := c.Bind(&reqGetRoomList); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	var listRoom []model.Room
//	condition := map[string]interface{}{
//		"isGetAll":        "false",
//		"room_type_code":  reqGetRoomList.RoomTypeCode,
//		"number_adult":    reqGetRoomList.NumberAdult,
//		"number_children": reqGetRoomList.NumberChildren,
//		"bed_number":      reqGetRoomList.NumberBed,
//		"number_toilet":   reqGetRoomList.NumberToilet,
//	}
//
//	listRoom, err := roomReceiver.RoomRepo.GetRoomListByCondition(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoom,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listRoom,
//	})
//}
//
//func (roomReceiver *RoomController) HandleGetRoomByCapacityAndTimeCheck(c echo.Context) error {
//	reqGetRoomListAvailable := req.RequestSearchRoomListByCapacityAndTimeCheck{}
//	//binding
//	if err := c.Bind(&reqGetRoomListAvailable); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	var listRoom []model.Room
//	condition := map[string]interface{}{
//		"isGetAll":        "false",
//		"room_type_code":  strings.ToLower("all"),
//		"number_adult":    reqGetRoomListAvailable.NumberAdult,
//		"number_children": reqGetRoomListAvailable.NumberChildren,
//		"start_time":      reqGetRoomListAvailable.TimeStart,
//		"end_time":        reqGetRoomListAvailable.TimeEnd,
//	}
//
//	listRoom, err := roomReceiver.RoomRepo.GetRoomListByCapacityAndTimeCheck(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoom,
//		})
//	}
//	tempStringIDs := ""
//	for tempIndex := 0; tempIndex < len(listRoom); tempIndex++ {
//		if tempIndex == len(listRoom)-1 {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'"
//		} else {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'" + ","
//		}
//	}
//	condition["room_ids"] = tempStringIDs
//	listRoomCountGroupByType, err := roomReceiver.RoomRepo.CountAvailableRoomAndGroupByType(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoomCountGroupByType,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listRoomCountGroupByType,
//	})
//}
//
//func (roomReceiver *RoomController) HandleSearchRoomAvailableByCapacityAndTimeCheckNumberRoom(c echo.Context) error {
//	reqGetRoomListAvailable := req.RequestSearchRoomAvailable{}
//	//binding
//	if err := c.Bind(&reqGetRoomListAvailable); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	//check number is valid
//	var listRoom []model.Room
//	condition := map[string]interface{}{
//		"isGetAll":        "false",
//		"room_type_code":  strings.ToLower("all"),
//		"number_adult":    reqGetRoomListAvailable.NumberAdult,
//		"number_children": reqGetRoomListAvailable.NumberChildren,
//		"start_time":      reqGetRoomListAvailable.TimeStart,
//		"end_time":        reqGetRoomListAvailable.TimeEnd,
//		"number_room":     reqGetRoomListAvailable.NumberRoom,
//	}
//
//	listRoom, err := roomReceiver.RoomRepo.GetRoomListByCapacityAndTimeCheck(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoom,
//		})
//	}
//	tempStringIDs := ""
//	for tempIndex := 0; tempIndex < len(listRoom); tempIndex++ {
//		if tempIndex == len(listRoom)-1 {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'"
//		} else {
//			tempStringIDs += "'" + listRoom[tempIndex].ID + "'" + ","
//		}
//	}
//	condition["room_ids"] = tempStringIDs
//	listRoomCountGroupByType, err := roomReceiver.RoomRepo.GetNumberRoomAvailableByCapacityAndTimeCheck(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoomCountGroupByType,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listRoomCountGroupByType,
//	})
//}
//
//// HandleSearchRoomAvailableByCapacityAndTimeCheckNumberRoomV2 godoc
//// @Summary Get room list filter search by capacity and time check
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestSearchRoomAvailable true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/search-num-room-valid [post]
//func (roomReceiver *RoomController) HandleSearchRoomAvailableByCapacityAndTimeCheckNumberRoomV2(c echo.Context) error {
//	reqGetRoomListAvailable := req.RequestSearchRoomAvailable{}
//	//binding
//	if err := c.Bind(&reqGetRoomListAvailable); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	condition := map[string]interface{}{
//		"number_adult":    reqGetRoomListAvailable.NumberAdult,
//		"number_children": reqGetRoomListAvailable.NumberChildren,
//		"start_time":      reqGetRoomListAvailable.TimeStart,
//		"end_time":        reqGetRoomListAvailable.TimeEnd,
//		"number_room":     reqGetRoomListAvailable.NumberRoom,
//	}
//	listRoomCountGroupByType, err := roomReceiver.RoomRepo.GetNumberRoomAvailableByCapacityAndTimeCheckV2(condition)
//
//	var listGroupByRoomCode []res.RoomTypeGroupRatePlan
//	//group by list here
//	for index, element := range listRoomCountGroupByType {
//		if len(listGroupByRoomCode) == 0 || index == 0 || listRoomCountGroupByType[index].TypeRoomCode != listRoomCountGroupByType[index-1].TypeRoomCode {
//			tempRoomRes := res.RoomTypeGroupRatePlan{
//				TypeRoomCode:     element.TypeRoomCode,
//				TypeRoomName:     element.TypeRoomName,
//				Remains:          element.Remains,
//				CostType:         element.CostType,
//				NumberAdult:      element.NumberAdult,
//				NumberChildren:   element.NumberChildren,
//				RoomImages:       element.RoomImages,
//				NumberBed:        element.NumberBed,
//				ShortDescription: element.ShortDescription,
//			}
//			var tempListRatePlan []res.RatePlanReduceModel
//			for _, elementRoom := range listRoomCountGroupByType {
//				if elementRoom.TypeRoomCode == element.TypeRoomCode {
//					tempRatePlan := res.RatePlanReduceModel{
//						RatePlanId:    elementRoom.RatePlanId,
//						RatePlanPrice: elementRoom.RatePlanPrice,
//						Description:   elementRoom.Description,
//					}
//					tempListRatePlan = append(tempListRatePlan, tempRatePlan)
//				}
//			}
//			tempRoomRes.ListRatePlan = tempListRatePlan
//			listGroupByRoomCode = append(listGroupByRoomCode, tempRoomRes)
//		}
//	}
//
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listGroupByRoomCode,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listGroupByRoomCode,
//	})
//}
//
//// HandleGetRoomListAtReception godoc
//// @Summary Get room list filter search time check at reception
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestSearchRoomAvailableAtReception true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/get-room-code [post]
//func (roomReceiver *RoomController) HandleGetRoomListAtReception(c echo.Context) error {
//	reqGetRoomListAvailable := req.RequestSearchRoomAvailableAtReception{}
//	//binding
//	if err := c.Bind(&reqGetRoomListAvailable); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	condition := map[string]interface{}{
//		"room_type_code": reqGetRoomListAvailable.RoomTypeCode,
//		"start_time":     reqGetRoomListAvailable.TimeStart,
//		"end_time":       reqGetRoomListAvailable.TimeEnd,
//	}
//	listRoomCountGroupByType, err := roomReceiver.RoomRepo.GetListRoomAtReception(condition)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    err.Error(),
//			Data:       listRoomCountGroupByType,
//		})
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listRoomCountGroupByType,
//	})
//}
//
//// HandleAddRoomStatusDetails godoc
//// @Summary Save room status time check at reception
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddRoomBusyStatusDetail true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/add-status-detail [post]
//func (roomReceiver *RoomController) HandleAddRoomStatusDetails(c echo.Context) error {
//	reqAddRoomBusyStatusDetail := req.RequestAddRoomBusyStatusDetail{}
//	//binding
//	if err := c.Bind(&reqAddRoomBusyStatusDetail); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	roomBusyStatusDetailId, err := uuid.NewUUID()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	timeCheckIn, err := time.Parse("2006-01-02", reqAddRoomBusyStatusDetail.FromTime)
//	if err != nil {
//		return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//	}
//	timeCheckOut, err := time.Parse("2006-01-02", reqAddRoomBusyStatusDetail.ToTime)
//	if err != nil {
//		return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//	}
//	roomBusyStatusCategory := model.RoomBusyStatusDetail{
//		ID:                       roomBusyStatusDetailId.String(),
//		RoomBusyStatusCategoryID: reqAddRoomBusyStatusDetail.StatusID,
//		BookingID:                reqAddRoomBusyStatusDetail.BookingID,
//		RoomID:                   reqAddRoomBusyStatusDetail.RoomID,
//		FromTime:                 timeCheckIn,
//		ToTime:                   timeCheckOut,
//	}
//	result, err := roomReceiver.RoomRepo.SaveRoomBusyStatusDetail(roomBusyStatusCategory)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lưu thành công", result)
//}
//
//func (roomReceiver *RoomController) HandleUpdateStatusRoomDetail(c echo.Context) error {
//	reqUpdateRoomInfo := req.RequestUpdateRoomBusyStatusDetail{}
//	if err := c.Bind(&reqUpdateRoomInfo); err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	err := c.Validate(reqUpdateRoomInfo)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	roomBusyStatusModel := model.RoomBusyStatusDetail{
//		ID:                       reqUpdateRoomInfo.ID,
//		RoomID:                   reqUpdateRoomInfo.RoomID,
//		RoomBusyStatusCategoryID: reqUpdateRoomInfo.StatusID,
//		FromTime:                 reqUpdateRoomInfo.FromTime,
//		ToTime:                   reqUpdateRoomInfo.ToTime,
//	}
//	roomBusyStatusModel, err = roomReceiver.RoomRepo.UpdateRoomBusyStatusDetail(roomBusyStatusModel)
//	if err != nil {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Cập nhật thông tin phòng thành công",
//		Data:       roomBusyStatusModel,
//	})
//}
//
//// HandleCheckOutBooking godoc
//// @Summary Check out room
//// @Tags booking-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestCheckOutBooking true "booking"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/check-out [post]
//func (roomReceiver *RoomController) HandleCheckOutBooking(c echo.Context) error {
//	reqCheckOutBooking := req.RequestCheckOutBooking{}
//	//binding
//	if err := c.Bind(&reqCheckOutBooking); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	roomBusyStatusDetailId, err := uuid.NewUUID()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	timeCheckIn, err := time.Parse("2006-01-02", reqCheckOutBooking.FromTime)
//	if err != nil {
//		return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//	}
//	timeCheckOut, err := time.Parse("2006-01-02", reqCheckOutBooking.ToTime)
//	if err != nil {
//		return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//	}
//	roomBusyStatusCategory := model.RoomBusyStatusDetail{
//		ID:                       roomBusyStatusDetailId.String(),
//		RoomBusyStatusCategoryID: "3",
//		BookingID:                reqCheckOutBooking.BookingID,
//		RoomID:                   reqCheckOutBooking.RoomID,
//		FromTime:                 timeCheckIn,
//		ToTime:                   timeCheckOut,
//	}
//	result, err := roomReceiver.RoomRepo.SaveRoomBusyStatusDetail(roomBusyStatusCategory)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Check out thành công",
//		Data:       result,
//	})
//}
//
//// HandleGetRoomStatusInfo godoc
//// @Summary Get room status info
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetRoomStatusInfo true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/get-room-status [post]
//func (roomReceiver *RoomController) HandleGetRoomStatusInfo(c echo.Context) error {
//	reqGetRoomStatusInfo := req.RequestGetRoomStatusInfo{}
//	//binding
//	if err := c.Bind(&reqGetRoomStatusInfo); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	condition := map[string]interface{}{
//		"booking_id": reqGetRoomStatusInfo.BookingID,
//	}
//	roomResult, err := roomReceiver.RoomRepo.GetStatusRoomCustomerInfo(condition)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy thông tin thành công",
//		Data:       roomResult,
//	})
//}
//
//// HandleGetRoomListCheckOut godoc
//// @Summary Get room list checkout
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetRoomStatusInfo true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/get-rooms-checkout [get]
//func (roomReceiver *RoomController) HandleGetRoomListCheckOut(c echo.Context) error {
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	timeCheckOut := time.Now()
//	fmt.Println(timeCheckOut.String())
//	fmt.Println(timeCheckOut.Format("2006-01-02 15:04:05"))
//	//timeCheckOut, _ := time.Parse("2006-01-02", time.Now().Format("yyyy-MM-dd"))
//	condition := map[string]interface{}{
//		"time": timeCheckOut.Format("2006-01-02") + " 00:00:00",
//	}
//	var roomResult []query.RoomStayedQuery
//	roomResult, err := roomReceiver.RoomRepo.GetListRoomCheckOut(condition)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       roomResult,
//	})
//}
//
//// HandleGetRoomRatePlanByTypeRoomCode godoc
//// @Summary Get room rateplan list by room type code
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetRoomRatePlan true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/get-room-rateplan [post]
//func (roomReceiver *RoomController) HandleGetRoomRatePlanByTypeRoomCode(c echo.Context) error {
//	reqGetRoomRatePlan := req.RequestGetRoomRatePlan{}
//	//binding
//	if err := c.Bind(&reqGetRoomRatePlan); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	responseResult := res.ResponseListRoomRatePlanByRoomTypeCode{
//		TypeRoomCode: reqGetRoomRatePlan.RoomTypeCode,
//	}
//	roomTypeQuery := model.RoomType{
//		//TypeRoomCode: reqGetRoomRatePlan.RoomTypeCode,
//	}
//	roomTypeResult, err := roomReceiver.RoomRepo.GetRoomTypeInfo(roomTypeQuery)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	var roomResult []model.Room
//	roomResult, err = roomReceiver.RoomRepo.GetRoomListByRoomTypeCode(roomTypeResult.ID)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	var listGroupByRoomFloor []res.ListRoomByFloorItem
//	for index, element := range roomResult {
//		if len(listGroupByRoomFloor) == 0 || index == 0 || roomResult[index].Floor != roomResult[index-1].Floor {
//			tempRoomFloorRes := res.ListRoomByFloorItem{
//				Floor: element.Floor,
//			}
//			var tempListRoom []model.Room
//			for _, elementRoom := range roomResult {
//				if elementRoom.Floor == element.Floor {
//					tempListRoom = append(tempListRoom, elementRoom)
//				}
//			}
//			tempRoomFloorRes.ListRoom = tempListRoom
//			listGroupByRoomFloor = append(listGroupByRoomFloor, tempRoomFloorRes)
//		}
//	}
//
//	ratePlanResult, err := roomReceiver.RoomRepo.GetListRatePlanByRoomTypeCode(roomTypeResult.ID)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	responseResult.ListRatePlan = ratePlanResult
//	responseResult.ListRoom = listGroupByRoomFloor
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       responseResult,
//	})
//}
//
//// HandleGetRoomCheckIn godoc
//// @Summary Get room check in
//// @Tags room-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetRoomAvailable true "room"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /room/get-room-check-in [post]
//func (roomReceiver *RoomController) HandleGetRoomCheckIn(c echo.Context) error {
//	reqGetRoomAvailable := req.RequestGetRoomAvailable{}
//	//binding
//	if err := c.Bind(&reqGetRoomAvailable); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	roomResult, err := roomReceiver.RoomRepo.GetListRoomCheckIn(reqGetRoomAvailable.RoomTypeID)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    "Lấy danh sách phòng thất bại",
//		})
//	}
//	var listGroupByRoomFloor []res.ListRoomByFloorItem
//	for index, element := range roomResult {
//		if len(listGroupByRoomFloor) == 0 || index == 0 || roomResult[index].Floor != roomResult[index-1].Floor {
//			tempRoomFloorRes := res.ListRoomByFloorItem{
//				Floor: element.Floor,
//			}
//			var tempListRoom []model.Room
//			for _, elementRoom := range roomResult {
//				if elementRoom.Floor == element.Floor {
//					tempListRoom = append(tempListRoom, elementRoom)
//				}
//			}
//			tempRoomFloorRes.ListRoom = tempListRoom
//			listGroupByRoomFloor = append(listGroupByRoomFloor, tempRoomFloorRes)
//		}
//	}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Lấy danh sách phòng thành công",
//		Data:       listGroupByRoomFloor,
//	})
//}
