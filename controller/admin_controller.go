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
)

type AdminController struct {
	AdminRepo repository.AdminRepo
}

// HandleCreateAccount godoc
// @Summary create account by admin
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestCreateStaffAccount true "staffaccount"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/create-staff-account [post]
func (adminController *AdminController) HandleCreateAccount(c echo.Context) error {
	reqRegister := req.RequestCreateAccountByAdmin{}
	//binding
	if err := c.Bind(&reqRegister); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	//Validate
	if err := c.Validate(reqRegister); err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(claims.Role == model.ADMIN.String()) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}
	//validate existed email
	//_, err := adminController.AdminRepo.CheckEmail(reqRegister.Email)
	//if err != nil {
	//	return response.Conflict(c, "Email đã tồn tại", nil)
	//}
	////Generate UUID
	//accountId, err := uuid.NewUUID()
	//if err != nil {
	//	logger.Error("Error uuid data", zap.Error(err))
	//	return response.Forbidden(c, err.Error(), nil)
	//}
	////create password
	//hash := security.HashAndSalt([]byte(reqRegister.Password))
	////Init account
	//account := model.User{
	//	ID:              accountId.String(),
	//	Email:           reqRegister.Email,
	//	Password:        hash,
	//	StaffID:         reqRegister.StaffID,
	//	CreatedAt:       time.Now(),
	//	UpdatedAt:       time.Now(),
	//	Role:            reqRegister.Role,
	//	Status: 1,
	//}
	//
	////Save account
	//account, err = adminController.AdminRepo.SaveStaffAccount(account)
	//if err != nil {
	//	return response.Conflict(c, err.Error(), nil)
	//}
	return response.Ok(c, "Đăng ký thành công", nil)
}

//// HandleSaveBookingStatus godoc
//// @Summary Save booking status
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddStatusBooking true "status booking"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/add-status-booking [post]
//func (adminController *AdminController) HandleSaveBookingStatus(c echo.Context) error {
//	reqAddStatusBooking := req.RequestAddStatusBooking{}
//	//binding
//	if err := c.Bind(&reqAddStatusBooking); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		logger.Error("Role is not available")
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	statusBooking := model.StatusBooking{
//		StatusCode:  reqAddStatusBooking.StatusCode,
//		StatusName:  reqAddStatusBooking.StatusName,
//		Description: reqAddStatusBooking.Description,
//	}
//	result, err := adminController.AdminRepo.SaveStatusBooking(statusBooking)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lưu thành công", result)
//}
//
//// HandleSaveWorkStatus godoc
//// @Summary Save work status
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddStatusWork true "status work"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/add-status-work [post]
//func (adminController *AdminController) HandleSaveWorkStatus(c echo.Context) error {
//	reqAddStatusWork := req.RequestAddStatusWork{}
//	//binding
//	if err := c.Bind(&reqAddStatusWork); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		logger.Error("Role is not available")
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	statusWork := model.StatusWork{
//		StatusCode:  reqAddStatusWork.StatusCode,
//		StatusName:  reqAddStatusWork.StatusName,
//		Description: reqAddStatusWork.Description,
//	}
//	result, err := adminController.AdminRepo.SaveStatusWork(statusWork)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lưu thành công", result)
//}
//
//// HandleSaveAccountStatus godoc
//// @Summary Save account status
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddStatusAccount true "status account"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/add-status-account [post]
//func (adminController *AdminController) HandleSaveAccountStatus(c echo.Context) error {
//	reqAddStatusAccount := req.RequestAddStatusAccount{}
//	//binding
//	if err := c.Bind(&reqAddStatusAccount); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	statusAccount := model.StatusAccount{
//		StatusCode:  reqAddStatusAccount.StatusCode,
//		StatusName:  reqAddStatusAccount.StatusName,
//		Description: reqAddStatusAccount.Description,
//	}
//	result, err := adminController.AdminRepo.SaveStatusAccount(statusAccount)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lưu thành công", result)
//}
//
//// HandleSignIn godoc
//// @Summary Sign In Account
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestSignInStaff true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 401 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/sign-in [post]
//func (adminController *AdminController) HandleSignIn(c echo.Context) error {
//	reqSignIn := req.RequestSignInStaff{}
//	if err := c.Bind(&reqSignIn); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	if err := c.Validate(reqSignIn); err != nil {
//		logger.Error("Error validate data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	account, err := adminController.AdminRepo.CheckLogin(reqSignIn.Email)
//	if err != nil {
//		return response.Unauthorized(c, err.Error(), nil)
//	}
//	isTheSamePass := security.ComparePasswords(account.Password, []byte(reqSignIn.Password))
//	if !isTheSamePass {
//		return response.Unauthorized(c, "Đăng nhập thất bại", "Sai mật khẩu")
//	}
//	if account.Role != model.ADMIN.String() {
//		return response.Unauthorized(c, "Đăng nhập thất bại", nil)
//	}
//	//generate token
//	token, err := security.GenStaffToken(account)
//	if err != nil {
//		logger.Error("err gen token", zap.Error(err))
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	account.Token = &model.Token{}
//	account.Token.AccessToken = token
//	tokenRefresh, timeDuration, err := security.GenStaffRefToken(account)
//	if err != nil {
//		logger.Error("err gen token data", zap.Error(err))
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	account.Token.RefreshToken = tokenRefresh
//	account.Token.ExpiredTime = timeDuration
//	return response.Ok(c, "Đăng nhập thành công", account)
//}
//
//// HandleChangePassword godoc
//// @Summary Change password
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestUpdatePassword true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/sign-in [post]
//func (adminController *AdminController) HandleChangePassword(c echo.Context) error {
//	reqUpdatePassword := req.RequestUpdatePassword{}
//	if err := c.Bind(&reqUpdatePassword); err != nil {
//		return err
//	}
//	err := c.Validate(reqUpdatePassword)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	account := model.StaffAccount{}
//	account, err = adminController.AdminRepo.GetAccountById(claims.UserId)
//	isTheSamePass := security.ComparePasswords(account.Password, []byte(reqUpdatePassword.OldPassword))
//	if !isTheSamePass {
//		return response.BadRequest(c, "Mật khẩu không khớp", nil)
//	}
//	hash := security.HashAndSalt([]byte(reqUpdatePassword.NewPassword))
//	isSuccess, _ := adminController.AdminRepo.UpdatePassword( /*c.Request().Context(), */ account.ID, hash)
//	if !isSuccess {
//		return response.InternalServerError(c, "Cập nhật mật khẩu thất bại", nil)
//	}
//	return response.Ok(c, "Cập nhật mật khẩu thành công", nil)
//}
//
//// HandleGetAccountInfo godoc
//// @Summary Get account info
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetAccountInfoStaff true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/account-info [post]
//func (adminController *AdminController) HandleGetAccountInfo(c echo.Context) error {
//	reqGetAccountInfo := req.RequestGetAccountInfoStaff{}
//	//binding
//	if err := c.Bind(&reqGetAccountInfo); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	account, err := adminController.AdminRepo.GetAccountById(claims.UserId)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy thông tin thành công", account)
//}
//
//// HandleGetStaffProfileInfo godoc
//// @Summary Get staff profile info
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetStaffProfileByAccountId true "staff"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/staff-info [post]
//func (adminController *AdminController) HandleGetStaffProfileInfo(c echo.Context) error {
//	reqGetProfileInfo := req.RequestGetStaffProfileByAccountId{}
//	//binding
//	if err := c.Bind(&reqGetProfileInfo); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	staff := model.Staff{
//		ID: reqGetProfileInfo.StaffId,
//	}
//
//	customerResult, err := adminController.AdminRepo.GetStaffProfile(staff.ID)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy thông tin thành công", customerResult)
//}
//
//// HandleCreateStaffAccount godoc
//// @Summary create staff account
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestCreateStaffAccount true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/create-staff-account [post]
//func (adminController *AdminController) HandleCreateStaffAccount(c echo.Context) error {
//	reqRegister := req.RequestCreateStaffAccount{}
//	//binding
//	if err := c.Bind(&reqRegister); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	//Validate
//	if err := c.Validate(reqRegister); err != nil {
//		logger.Error("Error validate data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	//validate existed email
//	_, err := adminController.AdminRepo.CheckEmail(reqRegister.Email)
//	if err != nil {
//		return response.Conflict(c, "Email đã tồn tại", nil)
//	}
//	//Generate UUID
//	accountId, err := uuid.NewUUID()
//	if err != nil {
//		logger.Error("Error uuid data", zap.Error(err))
//		return response.Forbidden(c, err.Error(), nil)
//	}
//	//create password
//	hash := security.HashAndSalt([]byte(reqRegister.Password))
//	//Init account
//	account := model.StaffAccount{
//		ID:              accountId.String(),
//		Email:           reqRegister.Email,
//		Password:        hash,
//		StaffID:         reqRegister.StaffID,
//		CreatedAt:       time.Now(),
//		UpdatedAt:       time.Now(),
//		Role:            reqRegister.Role,
//		StatusAccountID: 1,
//	}
//
//	//Save account
//	account, err = adminController.AdminRepo.SaveStaffAccount(account)
//	if err != nil {
//		return response.Conflict(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Đăng ký thành công", account)
//}
//
//// HandleUpdateStaffAccount godoc
//// @Summary update staff account
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestUpdateStaffAccount true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/update-staff-account [post]
//func (adminController *AdminController) HandleUpdateStaffAccount(c echo.Context) error {
//	reqChangeAccount := req.RequestUpdateStaffAccount{}
//	if err := c.Bind(&reqChangeAccount); err != nil {
//		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
//	}
//	err := c.Validate(reqChangeAccount)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	account := model.StaffAccount{
//		ID:     reqChangeAccount.ID,
//		Email:  reqChangeAccount.Email,
//		Avatar: reqChangeAccount.Avatar,
//	}
//	account, err = adminController.AdminRepo.UpdateStaffAccount(account)
//	if err != nil {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Cập nhật cài đặt thành công", account)
//}
//
//// HandleActivateAccount godoc
//// @Summary activate staff account
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestActivateAccountStaff true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/activate-staff-acc [post]
//func (adminController *AdminController) HandleActivateAccount(c echo.Context) error {
//	reqActivateAccount := req.RequestActivateAccountStaff{}
//	if err := c.Bind(&reqActivateAccount); err != nil {
//		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
//	}
//	err := c.Validate(reqActivateAccount)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if reqActivateAccount.AccountID == "" {
//		reqActivateAccount.AccountID = claims.UserId
//	}
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//
//	account := model.StaffAccount{
//		ID: reqActivateAccount.AccountID,
//	}
//
//	account, err = adminController.AdminRepo.ActivateStaffAccount(account)
//	if err != nil {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Mở khóa thành công", account)
//}
//
//// HandleDeactivateAccount godoc
//// @Summary deactivate staff account
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestDeactivateAccountStaff true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 422 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/deactivate-staff-acc [post]
//func (adminController *AdminController) HandleDeactivateAccount(c echo.Context) error {
//	reqActivateAccount := req.RequestDeactivateAccountStaff{}
//	if err := c.Bind(&reqActivateAccount); err != nil {
//		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
//	}
//	err := c.Validate(reqActivateAccount)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if reqActivateAccount.AccountID == "" {
//		reqActivateAccount.AccountID = claims.UserId
//	}
//	account := model.StaffAccount{
//		ID:     reqActivateAccount.AccountID,
//		Avatar: "https://placeimg.com/192/192/people",
//	}
//	account, err = adminController.AdminRepo.DeactivateStaffAccount( /*c.Request().Context(), */ account)
//	if err != nil {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Khóa thành công", account)
//}
//
//// HandleSaveStaffProfile godoc
//// @Summary Create staff profile
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddStaffProfile true "staff"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 403 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/create-staff-profile [post]
//func (adminController *AdminController) HandleSaveStaffProfile(c echo.Context) error {
//	reqRegister := req.RequestAddStaffProfile{}
//	//binding
//	if err := c.Bind(&reqRegister); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	//Validate
//	if err := c.Validate(reqRegister); err != nil {
//		logger.Error("Error validate data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	//validate existed email
//	_, err := adminController.AdminRepo.CheckEmail(reqRegister.Email)
//	if err != nil {
//		return response.Conflict(c, "Email đã tồn tại", nil)
//	}
//	//Generate UUID
//	staffId, err := uuid.NewUUID()
//	if err != nil {
//		logger.Error("Error uuid data", zap.Error(err))
//		return response.Forbidden(c, err.Error(), nil)
//	}
//	tempTimeStringConvert, _ := time.Parse("2006-01-02", reqRegister.DateOfBirth)
//	//Init staff
//	staff := model.Staff{
//		ID:               staffId.String(),
//		Email:            reqRegister.Email,
//		StaffID:          staffId.String(),
//		CreatedAt:        time.Now(),
//		UpdatedAt:        time.Now(),
//		Position:         reqRegister.Position,
//		StatusWorkID:     1,
//		HomeTown:         reqRegister.HomeTown,
//		Ethnic:           reqRegister.Ethnic,
//		IdentifierNumber: reqRegister.IdentifierNumber,
//		FirstName:        reqRegister.FirstName,
//		LastName:         reqRegister.LastName,
//		Phone:            reqRegister.Phone,
//		Gender:           reqRegister.Gender,
//		DateOfBirth:      tempTimeStringConvert,
//		Address:          reqRegister.Address,
//	}
//
//	//Save account
//	staff, err = adminController.AdminRepo.SaveStaffProfile(staff)
//	if err != nil {
//		return response.Conflict(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Đăng ký thành công", staff)
//}
//
//// HandleUpdateStaffProfile godoc
//// @Summary Update staff profile
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestUpdateStaffProfile true "staff"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 403 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/update-staff-profile [post]
//func (adminController *AdminController) HandleUpdateStaffProfile(c echo.Context) error {
//	reqRegister := req.RequestUpdateStaffProfile{}
//	//binding
//	if err := c.Bind(&reqRegister); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	//Validate
//	if err := c.Validate(reqRegister); err != nil {
//		logger.Error("Error validate data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	//validate existed email
//	_, err := adminController.AdminRepo.CheckEmail(reqRegister.Email)
//	if err != nil {
//		return response.Conflict(c, "Email đã tồn tại", nil)
//	}
//	if err != nil {
//		logger.Error("Error uuid data", zap.Error(err))
//		return response.Forbidden(c, err.Error(), nil)
//	}
//	timeDob, err := time.Parse("2006-01-02", reqRegister.DateOfBirth)
//	if err != nil {
//		return response.BadRequest(c, "Định dạng ngày không hợp lệ", nil)
//	}
//	//Init staff
//	staff := model.Staff{
//		ID:               reqRegister.StaffID,
//		Email:            reqRegister.Email,
//		CreatedAt:        time.Now(),
//		UpdatedAt:        time.Now(),
//		Position:         reqRegister.Position,
//		StatusWorkID:     1,
//		HomeTown:         reqRegister.HomeTown,
//		Ethnic:           reqRegister.Ethnic,
//		IdentifierNumber: reqRegister.IdentifierNumber,
//		FirstName:        reqRegister.FirstName,
//		LastName:         reqRegister.LastName,
//		Phone:            reqRegister.Phone,
//		Gender:           reqRegister.Gender,
//		DateOfBirth:      timeDob,
//		Address:          reqRegister.Address,
//	}
//
//	//Save account
//	staff, err = adminController.AdminRepo.UpdateStaffProfile(staff)
//	if err != nil {
//		return response.Conflict(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Cập nhật thành công", staff)
//}
//
//// HandleChangeRoleName godoc
//// @Summary Change staff role
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestChangeRoleName true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 401 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/change-role [post]
//func (adminController *AdminController) HandleChangeRoleName(c echo.Context) error {
//	reqChangeRole := req.RequestChangeRoleName{}
//	//binding
//	if err := c.Bind(&reqChangeRole); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	//Validate
//	if err := c.Validate(reqChangeRole); err != nil {
//		logger.Error("Error validate data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	//validate existed email
//	_, err := adminController.AdminRepo.ChangeRoleAccount(reqChangeRole.AccountID, reqChangeRole.Role)
//	if err != nil {
//		logger.Error("Error change role data", zap.Error(err))
//		return response.Unauthorized(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Cập nhật thành công", nil)
//}
//
//// HandleResetPassword godoc
//// @Summary Handle reset password
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/reset-pwd [post]
//func (adminController *AdminController) HandleResetPassword(c echo.Context) error {
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	hash := security.HashAndSalt([]byte("Siphoria@2022"))
//	isSuccess, _ := adminController.AdminRepo.UpdatePassword(claims.UserId, hash)
//	if !isSuccess {
//		return response.InternalServerError(c, "Cập nhật mật khẩu thất bại", nil)
//	}
//	return response.Ok(c, "Cập nhật mật khẩu thành công", nil)
//}
//
//// HandleSavePaymentStatus godoc
//// @Summary Save payment status
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestAddStatusPayment true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/add-status-payment [post]
//func (adminController *AdminController) HandleSavePaymentStatus(c echo.Context) error {
//	reqAddPaymentStatus := req.RequestAddStatusPayment{}
//	//binding
//	if err := c.Bind(&reqAddPaymentStatus); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	paymentStatus := model.PaymentStatus{
//		StatusCode:  reqAddPaymentStatus.StatusCode,
//		StatusName:  reqAddPaymentStatus.StatusName,
//		Description: reqAddPaymentStatus.Description,
//	}
//	result, err := adminController.AdminRepo.SavePaymentStatus(paymentStatus)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//
//	return response.Ok(c, "Lưu thành công", result)
//}
//
//// HandleGetAllAccountStatus godoc
//// @Summary Get all account status list
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/account-statuses [get]
//func (adminController *AdminController) HandleGetAllAccountStatus(c echo.Context) error {
//	var listAccountStatus []model.StatusAccount
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	listAccountStatus, err := adminController.AdminRepo.GetAccountStatusList()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), listAccountStatus)
//	}
//	return response.Ok(c, "Lấy trạng thái tài khoản thành công", listAccountStatus)
//}
//
//// HandleGetAllWorkStatus godoc
//// @Summary Get all work status list
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/work-statuses [get]
//func (adminController *AdminController) HandleGetAllWorkStatus(c echo.Context) error {
//	var listWorkStatus []model.StatusWork
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	listWorkStatus, err := adminController.AdminRepo.GetWorkStatusList()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), listWorkStatus)
//	}
//	return response.Ok(c, "Lấy trạng thái làm việc thành công", listWorkStatus)
//}
//
//// HandleStatisticRevenueDay godoc
//// @Summary Get statistic revenue day
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestStatisticRevenueByDay true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/revenue-statistic-day [post]
//func (adminController *AdminController) HandleStatisticRevenueDay(c echo.Context) error {
//	reqStatisticRevenue := req.RequestStatisticRevenueByDay{}
//	//binding
//	if err := c.Bind(&reqStatisticRevenue); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	condition := map[string]interface{}{
//		"time_start": reqStatisticRevenue.TimeStart,
//		"time_end":   reqStatisticRevenue.TimeEnd,
//	}
//	result, err := adminController.AdminRepo.GetStatisticRevenueByDay(condition)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	var listGroupByRoomCodeRevenue []res.StatisticRevenueByDay
//	//group by list here
//	for index, element := range result {
//		if len(result) == 0 || index == 0 || (result[index].PaymentTime.Day() != result[index-1].PaymentTime.Day() || result[index].PaymentTime.Month() != result[index-1].PaymentTime.Month() || result[index].PaymentTime.Year() != result[index-1].PaymentTime.Year()) {
//			tempStatisticRes := res.StatisticRevenueByDay{
//				Day: element.PaymentTime.String(),
//			}
//			var tempListRoomType []res.RoomTypeRevenue
//			var totalRevenue float32
//			totalRevenue = 0
//			for _, elementRoom := range result {
//				if reqStatisticRevenue.Mode == 0 {
//					if elementRoom.PaymentTime.Day() == element.PaymentTime.Day() && elementRoom.PaymentTime.Month() == element.PaymentTime.Month() && elementRoom.PaymentTime.Year() == element.PaymentTime.Year() {
//						totalRevenue = totalRevenue + elementRoom.Sum
//					}
//				} else {
//					if elementRoom.PaymentTime.Day() == element.PaymentTime.Day() && elementRoom.PaymentTime.Month() == element.PaymentTime.Month() && elementRoom.PaymentTime.Year() == element.PaymentTime.Year() {
//						tempRoomTypeRevenue := res.RoomTypeRevenue{
//							RoomTypeCode: elementRoom.RoomTypeCode,
//							Sum:          elementRoom.Sum,
//						}
//						tempListRoomType = append(tempListRoomType, tempRoomTypeRevenue)
//					}
//				}
//
//			}
//			if reqStatisticRevenue.Mode == 0 {
//				tempRoomTypeRevenue := res.RoomTypeRevenue{
//					Sum: totalRevenue,
//				}
//				tempListRoomType = append(tempListRoomType, tempRoomTypeRevenue)
//			}
//			tempStatisticRes.ListRoomTypeRevenue = tempListRoomType
//			listGroupByRoomCodeRevenue = append(listGroupByRoomCodeRevenue, tempStatisticRes)
//		}
//	}
//	return response.Ok(c, "Lấy dữ liệu thành công", listGroupByRoomCodeRevenue)
//}
//
//// HandleStatisticRevenueTypeRoomCode godoc
//// @Summary Get statistic revenue type room code
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestStatisticRevenueByRoomType true "payment"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/revenue-statistic-type-code [post]
//func (adminController *AdminController) HandleStatisticRevenueTypeRoomCode(c echo.Context) error {
//	reqStatisticRevenue := req.RequestStatisticRevenueByRoomType{}
//	//binding
//	if err := c.Bind(&reqStatisticRevenue); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	condition := map[string]interface{}{
//		"time_start": reqStatisticRevenue.TimeStart,
//		"time_end":   reqStatisticRevenue.TimeEnd,
//	}
//	result, err := adminController.AdminRepo.GetStatisticRevenueByRoomTypeCode(condition)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy dữ liệu thành công", result)
//}
//
//// HandleGetAllCustomer godoc
//// @Summary Get all customer
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/get-all-customer [get]
//func (adminController *AdminController) HandleGetAllCustomer(c echo.Context) error {
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	result, err := adminController.AdminRepo.GetAllCustomer()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy dữ liệu thành công", result)
//}
//
//// HandleGetAllStaff godoc
//// @Summary Get all staff
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/get-all-staff [get]
//func (adminController *AdminController) HandleGetAllStaff(c echo.Context) error {
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	result, err := adminController.AdminRepo.GetAllStaff()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy dữ liệu thành công", result)
//}
//
//// HandleGetAllStaffAccount godoc
//// @Summary Get all staff account
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/get-all-staff-account [get]
//func (adminController *AdminController) HandleGetAllStaffAccount(c echo.Context) error {
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	result, err := adminController.AdminRepo.GetAllStaffAccount()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy dữ liệu thành công", result)
//}
//
//// HandleGetAllAccount godoc
//// @Summary Get all account
//// @Tags admin-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /admin/get-all-account [get]
//func (adminController *AdminController) HandleGetAllAccount(c echo.Context) error {
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.ADMIN.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	result, err := adminController.AdminRepo.GetAllCustomerAccount()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy dữ liệu thành công", result)
//}
