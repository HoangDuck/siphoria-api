package controller

//
//import (
//	"github.com/golang-jwt/jwt"
//	"github.com/labstack/echo/v4"
//	"go.uber.org/zap"
//	"hotel-booking-api/logger"
//	"hotel-booking-api/model"
//	"hotel-booking-api/model/model_func"
//	"hotel-booking-api/model/req"
//	"hotel-booking-api/repository"
//	"hotel-booking-api/security"
//	"time"
//)
//
//type StaffController struct {
//	StaffRepo repository.StaffRepo
//	UserRepo  repository.UserRepo
//}
//
//// HandleSignIn godoc
//// @Summary Sign In Account
//// @Tags staff-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestSignInStaff true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 401 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /staff/sign-in [post]
//func (staffController *StaffController) HandleSignIn(c echo.Context) error {
//	reqSignIn := req.RequestSignInStaff{}
//	if err := c.Bind(&reqSignIn); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	if err := c.Validate(reqSignIn); err != nil {
//		logger.Error("Error validate data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	account, err := staffController.StaffRepo.CheckLogin(reqSignIn.Email)
//	if err != nil {
//		return response.Unauthorized(c, err.Error(), nil)
//	}
//	isTheSamePass := security.ComparePasswords(account.Password, []byte(reqSignIn.Password))
//	if !isTheSamePass {
//		return response.Unauthorized(c, "Đăng nhập thất bại", nil)
//	}
//	if account.Role != model.STAFF.String() {
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
//// @Summary Change password Account
//// @Tags staff-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestUpdatePasswordStaff true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /staff/change-pwd [post]
//func (staffController *StaffController) HandleChangePassword(c echo.Context) error {
//	reqUpdatePassword := req.RequestUpdatePasswordStaff{}
//	if err := c.Bind(&reqUpdatePassword); err != nil {
//		return err
//	}
//	err := c.Validate(reqUpdatePassword)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	account := model.StaffAccount{}
//	account, err = staffController.StaffRepo.GetAccountById(claims.UserId)
//	isTheSamePass := security.ComparePasswords(account.Password, []byte(reqUpdatePassword.OldPassword))
//	if !isTheSamePass {
//		return response.BadRequest(c, "Mật khẩu không khớp", nil)
//	}
//	hash := security.HashAndSalt([]byte(reqUpdatePassword.NewPassword))
//	isSuccess, _ := staffController.StaffRepo.UpdatePassword( /*c.Request().Context(), */ account.ID, hash)
//	if !isSuccess {
//		return response.InternalServerError(c, "Cập nhật mật khẩu thất bại", nil)
//	}
//
//	return response.Ok(c, "Cập nhật thành công", nil)
//}
//
//// HandleGetAccountInfo godoc
//// @Summary Get Account info
//// @Tags staff-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetAccountInfoStaff true "staffaccount"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /staff/account-info [post]
//func (staffController *StaffController) HandleGetAccountInfo(c echo.Context) error {
//	reqGetAccountInfo := req.RequestGetAccountInfoStaff{}
//	//binding
//	if err := c.Bind(&reqGetAccountInfo); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	if !(claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	account, err := staffController.StaffRepo.GetAccountById(claims.UserId)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy thông tin tài khoản thành công", account)
//}
//
//// HandleUpdateStaffProfile godoc
//// @Summary Update staff profile
//// @Tags staff-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestUpdateStaffProfile true "staff"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 403 {object} res.Response
//// @Failure 409 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /staff/update-profile [post]
//func (staffController *StaffController) HandleUpdateStaffProfile(c echo.Context) error {
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
//	if !(claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	//validate existed email
//	_, err := staffController.StaffRepo.CheckEmail(reqRegister.Email)
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
//	staff, err = staffController.StaffRepo.UpdatePersonalInfo(staff)
//	if err != nil {
//		return response.Conflict(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Cập nhật thành công", staff)
//}
//
//// HandleUpdateAvatarStaffProfile godoc
//// @Summary Update staff avatar
//// @Tags staff-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestChangeAvatarStaff true "staff"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /staff/change-avatar [post]
//func (staffController *StaffController) HandleUpdateAvatarStaffProfile(c echo.Context) error {
//	reqRegister := req.RequestChangeAvatarStaff{}
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
//	if !(claims.Role == model.STAFF.String()) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	//Init staff
//	staff := model.StaffAccount{
//		ID:     reqRegister.AccountID,
//		Avatar: reqRegister.AvatarUrl,
//	}
//
//	//Save account
//	staff, err := staffController.StaffRepo.UpdateAvatarInfo(staff.ID, staff.Avatar)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Cập nhật thành công", staff)
//}
