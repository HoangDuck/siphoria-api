package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/model/model_func"
	"hotel-booking-api/model/req"
	"hotel-booking-api/model/res"
	"hotel-booking-api/repository"
	"net/http"
)

type RatePlanController struct {
	RatePlanRepo repository.RatePlanRepo
}

// HandleGetListRatePlan godoc
// @Summary Get rateplan list
// @Tags rateplan-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /ratePlan/rateplans [get]
func (ratePlanController *RatePlanController) HandleGetListRatePlan(c echo.Context) error {
	var listRatePlan []model.RatePlan
	listRatePlan, err := ratePlanController.RatePlanRepo.GetListRatePlan()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       listRatePlan,
		})
	}
	return c.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy danh sách gói ưu đãi, tiện ích thành công",
		Data:       listRatePlan,
	})
}

// HandleGetRatePlanInfo godoc
// @Summary Rateplan info
// @Tags rateplan-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestGetRatePlan true "rateplan"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /ratePlan/rateplan-info [post]
func (ratePlanController *RatePlanController) HandleGetRatePlanInfo(c echo.Context) error {
	reqGetRatePlanInfo := req.RequestGetRatePlan{}
	//binding
	if err := c.Bind(&reqGetRatePlanInfo); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	ratePlanModel := model.RatePlan{
		ID: reqGetRatePlanInfo.RatePlanID,
	}
	ratePlan, err := ratePlanController.RatePlanRepo.GetRatePlanInfo(ratePlanModel)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return c.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy thông tin gói ưu đãi, tiện ích thành công",
		Data:       ratePlan,
	})
}

// HandleUpdateRatePlan godoc
// @Summary Update rateplan
// @Tags rateplan-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateRatePlan true "rateplan"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /ratePlan/update-rateplan [post]
func (ratePlanController *RatePlanController) HandleUpdateRatePlan(c echo.Context) error {
	reqUpdateRatePlan := req.RequestUpdateRatePlan{}
	if err := c.Bind(&reqUpdateRatePlan); err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	err := c.Validate(reqUpdateRatePlan)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	ratePlanModel := model.RatePlan{
		ID: reqUpdateRatePlan.RatePlanID,
		//RoomTypeID:  reqUpdateRatePlan.RoomTypeID,
		//Description: reqUpdateRatePlan.Description,
		//Price:       reqUpdateRatePlan.Price,
	}
	ratePlan, err := ratePlanController.RatePlanRepo.UpdateRatePlanInfo(ratePlanModel)
	if err != nil {
		return response.UnprocessableEntity(c, err.Error(), nil)
	}
	return c.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "Cập nhật thành công",
		Data:       ratePlan,
	})
}

// HandleDeleteRatePlan godoc
// @Summary Delete rateplan
// @Tags rateplan-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestDeleteRatePlan true "rateplan"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /ratePlan/delete-rateplan [post]
func (ratePlanController *RatePlanController) HandleDeleteRatePlan(c echo.Context) error {
	reqDeleteRatePlan := req.RequestDeleteRatePlan{}
	//binding
	if err := c.Bind(&reqDeleteRatePlan); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	condition := map[string]interface{}{
		"ID": reqDeleteRatePlan.RatePlanID,
	}
	result, err := ratePlanController.RatePlanRepo.DeleteRatePlanInfo(condition)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	if result != true {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return c.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "Xóa gói ưu đãi, tiện ích thành công",
		Data:       nil,
	})
}

// HandleSaveRatePlan godoc
// @Summary Save rateplan
// @Tags rateplan-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestAddRatePlan true "rateplan"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /ratePlan/save-rateplan [post]
func (ratePlanController *RatePlanController) HandleSaveRatePlan(c echo.Context) error {
	reqAddRatePlan := req.RequestAddRatePlan{}
	//binding
	if err := c.Bind(&reqAddRatePlan); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.STAFF.String() || claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	roomType := model.RoomType{
		//TypeRoomCode: reqAddRatePlan.RoomTypeCode,
	}
	_, err := ratePlanController.RatePlanRepo.GetRoomTypeInfo(roomType)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	ratePlanId, err := uuid.NewUUID()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	ratePlan := model.RatePlan{
		ID: ratePlanId.String(),
		//RoomTypeID:  roomTypeModel.ID,
		//Description: reqAddRatePlan.Description,
		//Price:       reqAddRatePlan.Price,
	}
	result, err := ratePlanController.RatePlanRepo.SaveRatePlan(ratePlan)
	if err != nil || !result {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Lưu thành công", nil)
}
