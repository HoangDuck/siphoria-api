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
	"hotel-booking-api/security"
	"hotel-booking-api/utils"
	"time"
)

type VoucherController struct {
	VoucherRepo repository.VoucherRepo
}

// HandleSaveVoucher godoc
// @Summary Save voucher
// @Tags voucher-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestAddVoucher true "voucher"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /vouchers [post]
func (voucherController *VoucherController) HandleSaveVoucher(c echo.Context) error {
	reqAddVoucher := req.RequestAddVoucher{}
	//binding
	if err := c.Bind(&reqAddVoucher); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) || security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	voucherId, err := utils.GetNewId()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	dateBeginAt, err := time.Parse("2006-01-02", reqAddVoucher.BeginAt)
	dateEndAt, err := time.Parse("2006-01-02", reqAddVoucher.EndAt)
	if err != nil {
	}
	voucher := model.Voucher{
		ID:        voucherId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		HotelId:   reqAddVoucher.HotelID,
		Name:      reqAddVoucher.Name,
		Discount:  reqAddVoucher.Discount,
		Activated: reqAddVoucher.Activated,
		Code:      reqAddVoucher.Code,
		BeginAt:   dateBeginAt,
		EndAt:     dateEndAt,
		IsDeleted: false,
	}
	result, err := voucherController.VoucherRepo.SaveVoucher(voucher)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}

	listVoucherExcept, err := voucherController.VoucherRepo.SaveBatchVoucher(reqAddVoucher.ExceptRoom, result.ID)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	result.Excepts = listVoucherExcept
	return response.Ok(c, "Lưu thành công", result)
}

// HandleUpdateVoucher godoc
// @Summary Update voucher
// @Tags voucher-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestUpdateVoucher true "voucher"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /vouchers/: [patch]
func (voucherController *VoucherController) HandleUpdateVoucher(c echo.Context) error {
	reqUpdateVoucher := req.RequestUpdateVoucher{}
	//binding
	if err := c.Bind(&reqUpdateVoucher); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) || security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	voucherId, err := utils.GetNewId()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	dateBeginAt, err := time.Parse("2006-01-02", reqUpdateVoucher.BeginAt)
	dateEndAt, err := time.Parse("2006-01-02", reqUpdateVoucher.EndAt)
	if err != nil {
	}
	voucher := model.Voucher{
		ID:        voucherId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		HotelId:   c.Param("id"),
		Name:      reqUpdateVoucher.Name,
		Discount:  reqUpdateVoucher.Discount,
		Activated: reqUpdateVoucher.Activated,
		Code:      reqUpdateVoucher.Code,
		BeginAt:   dateBeginAt,
		EndAt:     dateEndAt,
	}
	result, err := voucherController.VoucherRepo.UpdateVoucher(voucher)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}

	listVoucherExcept, err := voucherController.VoucherRepo.SaveBatchVoucher(reqUpdateVoucher.ExceptRoom, result.ID)
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	result.Excepts = listVoucherExcept
	return response.Ok(c, "Lưu thành công", result)
}

// HandleDeleteVoucher godoc
// @Summary Delete voucher
// @Tags voucher-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /vouchers/:id [delete]
func (voucherController *VoucherController) HandleDeleteVoucher(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.HOTELIER, false) || security.CheckRole(claims, model.MANAGER, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	voucherId := c.Param("id")
	voucher := model.Voucher{
		ID:        voucherId,
		IsDeleted: true,
	}
	result, err := voucherController.VoucherRepo.DeleteVoucher(voucher)
	if !result {
		return response.InternalServerError(c, err.Error(), nil)
	}
	return response.Ok(c, "Xoá thành công", nil)
}
