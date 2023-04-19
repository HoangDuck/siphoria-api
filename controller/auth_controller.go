package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
	"hotel-booking-api/security"
	"hotel-booking-api/services"
	"os"
)

type AuthController struct {
	AccountRepo repository.AccountRepo
}

//// HandleGetCustomerAccountInfo godoc
//// @Summary Get Customer Account Info
//// @Tags auth-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetAccountInfo true "account"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /auth/account-info [post]
//func (authReceiver *AuthController) HandleGetCustomerAccountInfo(c echo.Context) error {
//	reqGetAccountInfo := req.RequestGetAccountInfo{}
//	//binding
//	if err := c.Bind(&reqGetAccountInfo); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	claims := security.GetClaimsJWT(c)
//	if security.CheckRole(claims, model.ADMIN) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	account, err := authReceiver.AccountRepo.GetAccountById(claims.UserId)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy thông tin thành công", account)
//}
//
//// TestSendEmail godoc
//// @Summary Test send email
//// @Tags auth-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Router /auth/send-email [get]
//func (authReceiver *AuthController) TestSendEmail(c echo.Context) error {
//	testEmailService := services.GetEmailServiceInstance()
//	customerPageUrl, err := authReceiver.AccountRepo.GetCustomerPageUrl()
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	err = testEmailService.SendEmailService(
//		[]string{"19110375@student.hcmute.edu.vn"},
//		testEmailService.CreateBodyTemplate("Reset mật khẩu", "Đặt lại mật khẩu", customerPageUrl+"ABC"),
//		"Test email")
//	if err != nil {
//		return response.BadRequest(c, "Gửi email thất bại", nil)
//	}
//	return response.Ok(c, "Gửi email thành công", nil)
//}

// HandleRegister godoc
// @Summary Create new account
// @Tags auth-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestSignUp true "account"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 403 {object} res.Response
// @Failure 409 {object} res.Response
// @Router /auth/signup [post]
func (authReceiver *AuthController) HandleRegister(c echo.Context) error {
	reqRegister := req.RequestSignUp{}
	//binding
	if err := c.Bind(&reqRegister); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, "Thông tin không hợp lệ", nil)
	}
	//Validate
	if err := c.Validate(reqRegister); err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, "Thông tin không hợp lệ", nil)
	}
	//validate existed email
	_, err := authReceiver.AccountRepo.CheckEmailExisted(reqRegister.Email)
	if err != nil {
		logger.Error("Error existed email", zap.Error(err))
		return response.Conflict(c, "Email đã tồn tại", nil)
	}
	//Generate UUID
	accountId, err := uuid.NewUUID()
	if err != nil {
		logger.Error("Error uuid data", zap.Error(err))
		return response.Forbidden(c, "Đăng ký thất bại", nil)
	}
	//create password
	hash := security.HashAndSalt([]byte(reqRegister.Password))
	// Generate Verification Code

	//Init account
	account := model.User{
		ID:        accountId.String(),
		Email:     reqRegister.Email,
		Password:  hash,
		FirstName: reqRegister.FirstName,
		LastName:  reqRegister.LastName,
		FullName:  reqRegister.FirstName + reqRegister.LastName,
		Role:      model.CUSTOMER.String(),
		Status:    1,
	}
	//Save account
	accountResult, err := authReceiver.AccountRepo.SaveAccount(account)

	if err != nil {
		logger.Error("Error uuid data", zap.Error(err))
		return response.InternalServerError(c, "Đăng ký thất bại", nil)
	}
	_, err = security.GenToken(&accountResult)
	if err != nil {
		logger.Error("err gen token", zap.Error(err))
		return response.InternalServerError(c, "Đăng ký thất bại", nil)
	}
	_, _, err = security.GenRefToken(&accountResult)
	if err != nil {
		logger.Error("err gen token data", zap.Error(err))
		return response.InternalServerError(c, "Đăng ký thất bại", nil)
	}
	sendResetPasswordEmailService := services.GetEmailServiceInstance()
	customerPageUrl, err := authReceiver.AccountRepo.GetCustomerActivatePageUrl()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	err = sendResetPasswordEmailService.SendEmailService(
		[]string{account.Email},
		sendResetPasswordEmailService.CreateBodyTemplate("Nhấn nút bên dưới để thực hiện kích hoạt",
			"Xác nhận tài khoản",
			customerPageUrl+accountId.String()),
		"Siphoria activate account")
	if err != nil {
		return response.InternalServerError(c, "Gửi email thất bại", nil)
	}
	return response.Ok(c, "Đăng ký thành công", accountResult.Token)
}

//func (authReceiver *AuthController) HandleAuthenticateWithGoogle(c echo.Context) error {
//	oauthGoogleServiceInstance := services.GetOauth2ServiceInstance()
//	oauthGoogleServiceInstance.GoogleAuthenticationService(c.Response(), c.Request())
//	return c.String(200, "Redirect URL")
//}
//
//func (authReceiver *AuthController) HandleAuthenticateWithGoogleCallBack(c echo.Context) error {
//	oauthGoogleServiceInstance := services.GetOauth2ServiceInstance()
//	dataContent := oauthGoogleServiceInstance.AuthenticationCallBack(c.Response(), c.Request())
//	//_, err := authReceiver.AccountRepo.CheckEmail(dataContent["email"].(string))
//	//
//	//if err != nil {
//	//	dataContent["isExisted"] = "true"
//	//	return c.JSON(http.StatusOK, res.Response{
//	//		StatusCode: http.StatusOK,
//	//		Message:    "Get data success",
//	//		Data:       dataContent,
//	//	})
//	//}
//	return c.JSON(http.StatusOK, res.Response{
//		StatusCode: http.StatusOK,
//		Message:    "Get data success",
//		Data:       dataContent,
//	})
//}

// HandleSignIn godoc
// @Summary Sign In Account
// @Tags auth-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestSignIn true "account"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/login/general [post]
func (authReceiver *AuthController) HandleSignIn(c echo.Context) error {
	reqSignIn := req.RequestSignIn{}
	if err := c.Bind(&reqSignIn); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	if err := c.Validate(reqSignIn); err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	account, err := authReceiver.AccountRepo.CheckLogin(reqSignIn)
	if err != nil {
		return response.Unauthorized(c, err.Error(), nil)
	}
	if account.Status == 0 {
		return response.Unauthorized(c, "Tài khoản bị khóa", nil)
	}
	isTheSamePass := security.ComparePasswords(account.Password, []byte(reqSignIn.Password))
	if !isTheSamePass {
		return response.Unauthorized(c, "Đăng nhập thất bại", nil)
	}
	//generate token
	_, err = security.GenToken(&account)
	if err != nil {
		logger.Error("err gen token", zap.Error(err))
		return response.InternalServerError(c, "Đăng nhập thất bại", nil)
	}
	_, _, err = security.GenRefToken(&account)
	if err != nil {
		logger.Error("err gen token data", zap.Error(err))
		return response.InternalServerError(c, "Đăng nhập thất bại", nil)
	}
	return response.Ok(c, "Đăng nhập thành công", account.Token)
}

//// CheckEmailExisted godoc
//// @Summary Check email is existed
//// @Tags auth-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestCheckEmail true "user"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Router /auth/check-email [post]
//func (authReceiver *AuthController) CheckEmailExisted(c echo.Context) error {
//	reqCheckEmail := req.RequestCheckEmail{}
//	err := c.Bind(&reqCheckEmail)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, res.Response{
//			StatusCode: http.StatusBadRequest,
//			Message:    "Yêu cầu bị từ chối",
//		})
//	}
//	err = c.Validate(reqCheckEmail)
//	if err != nil {
//		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
//	}
//	isExisted, errorCheck := authReceiver.AccountRepo.CheckEmailExisted( /*c.Request().Context(), */ reqCheckEmail.Email)
//	if errorCheck != nil {
//		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
//	}
//	if isExisted.ID != "-1" {
//		return c.JSON(http.StatusOK, res.Response{
//			StatusCode: http.StatusOK,
//			Message:    "Email đã tồn tại trong hệ thống",
//		})
//	} else {
//		return c.JSON(http.StatusOK, res.Response{
//			StatusCode: http.StatusOK,
//			Message:    "Email chưa tồn tại",
//		})
//	}
//}

// HandleSendEmailResetPassword godoc
// @Summary Send Email To Reset Password
// @Tags auth-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestResetPassword true "account"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/forgot [post]
func (authReceiver *AuthController) HandleSendEmailResetPassword(c echo.Context) error {
	requestResetPassword := req.RequestResetPassword{}
	if err := c.Bind(&requestResetPassword); err != nil {
		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
	}

	if err := c.Validate(requestResetPassword); err != nil {
		return response.BadRequest(c, "Email không khả dụng", nil)
	}
	isExistedEmail, _ := authReceiver.AccountRepo.CheckEmailExisted(requestResetPassword.Email)
	if !isExistedEmail {
		return response.InternalServerError(c, "Email không tồn tại", nil)
	}
	//generate token
	token, err := security.GenTokenResetPassword(requestResetPassword.Email)
	if err != nil {
		return response.InternalServerError(c, "Tạo token thất bại", nil)
	}
	sendResetPasswordEmailService := services.GetEmailServiceInstance()
	customerPageUrl, err := authReceiver.AccountRepo.GetCustomerPageUrl()
	if err != nil {
		return response.InternalServerError(c, err.Error(), nil)
	}
	err = sendResetPasswordEmailService.SendEmailService(
		[]string{requestResetPassword.Email},
		sendResetPasswordEmailService.CreateBodyTemplate("Nhấn nút bên dưới để thực hiện đặt lại mật khẩu",
			"Đặt lại mật khẩu",
			customerPageUrl+token),
		"Siphoria reset password")
	if err != nil {
		return response.InternalServerError(c, "Gửi email thất bại", nil)
	}
	return response.Ok(c, "Gửi email thành công", nil)
}

// HandleResetPassword godoc
// @Summary Handle Reset Password
// @Tags auth-service
// @Accept  json
// @Produce  json
// @Param data body req.RequestNewPasswordReset true "account"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/reset [post]
func (authReceiver *AuthController) HandleResetPassword(c echo.Context) error {
	reqNewPassword := req.RequestNewPasswordReset{}
	if err := c.Bind(&reqNewPassword); err != nil {
		return err
	}
	err := c.Validate(reqNewPassword)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}
	hash := security.HashAndSalt([]byte(reqNewPassword.Password))
	claims := jwt.MapClaims{}
	tokenResult, err := jwt.ParseWithClaims(reqNewPassword.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	var email string
	for key, val := range claims {
		if key == "email" {
			email = val.(string)
		}
	}
	if tokenResult.Valid {
		isSuccess, err := authReceiver.AccountRepo.ResetPassword(email, hash)
		if isSuccess == false && err != nil {
			return response.InternalServerError(c, "Đặt lại mật khẩu thất bại", nil)
		}
		return response.Ok(c, "Đặt lại mật khẩu thành công", nil)
	}
	return response.InternalServerError(c, "Đặt lại mật khẩu thất bại", nil)
}

//// HandleChangePassword godoc
//// @Summary Handle Change Password
//// @Tags auth-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestUpdatePassword true "account"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /auth/change-pwd [post]
//func (authReceiver *AuthController) HandleChangePassword(c echo.Context) error {
//	reqUpdatePassword := req.RequestUpdatePassword{}
//	if err := c.Bind(&reqUpdatePassword); err != nil {
//		return err
//	}
//	err := c.Validate(reqUpdatePassword)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	claims := security.GetClaimsJWT(c)
//	if !security.CheckRole(claims, model.CUSTOMER) {
//		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
//	}
//	account, err := authReceiver.AccountRepo.GetAccountById(claims.UserId)
//	isTheSamePass := security.ComparePasswords(account.Password, []byte(reqUpdatePassword.OldPassword))
//	if !isTheSamePass {
//		return response.BadRequest(c, "Mật khẩu không khớp", nil)
//	}
//	hash := security.HashAndSalt([]byte(reqUpdatePassword.NewPassword))
//	isSuccess, _ := authReceiver.AccountRepo.UpdatePassword(account.ID, hash)
//	if !isSuccess {
//		return response.InternalServerError(c, "Cập nhật mật khẩu thất bại", nil)
//	}
//	return response.Ok(c, "Cập nhật mật khẩu thành công", nil)
//}

// HandleActivateAccount godoc
// @Summary Handle Activate Account
// @Tags auth-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 422 {object} res.Response
// @Router /auth/verifyemail/:code [get]
func (authReceiver *AuthController) HandleActivateAccount(c echo.Context) error {
	emailCode := c.Param("code")
	account := model.User{
		ID: emailCode,
	}
	account, err := authReceiver.AccountRepo.ActivateAccount(account)
	if err != nil {
		logger.Error("err active account data", zap.Error(err))
		return response.UnprocessableEntity(c, err.Error(), nil)
	}
	return response.Ok(c, "Mở khóa thành công", account)
}

//// HandleDeactivateAccount godoc
//// @Summary Handle Deactivate Account
//// @Tags auth-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 422 {object} res.Response
//// @Router /auth/deactive/by/account-id [get]
//func (authReceiver *AuthController) HandleDeactivateAccount(c echo.Context) error {
//	claims := security.GetClaimsJWT(c)
//	account := model.Account{
//		ID: claims.UserId,
//	}
//	account, err := authReceiver.AccountRepo.DeactivateAccount(account)
//	if err != nil {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Khóa thành công", account)
//}
//
//// HandleChangeAccountSettings godoc
//// @Summary Handle update settings Account
//// @Tags auth-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestChangeSettings true "account"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 422 {object} res.Response
//// @Router /auth/update-setting [post]
//func (authReceiver *AuthController) HandleChangeAccountSettings(c echo.Context) error {
//	reqChangeAccountSettings := req.RequestChangeSettings{}
//	if err := c.Bind(&reqChangeAccountSettings); err != nil {
//		return response.BadRequest(c, "Yêu cầu không hợp lệ", nil)
//	}
//	err := c.Validate(reqChangeAccountSettings)
//	if err != nil {
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	account := model.Account{
//		ID: claims.UserId,
//	}
//	account, err = authReceiver.AccountRepo.UpdateAccountSettings( /*c.Request().Context(), */ account)
//	if err != nil {
//		return response.UnprocessableEntity(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Cập nhật cài đặt thành công", account)
//}
//
//// HandleSignInGoogleToken godoc
//// @Summary Handle Sign in google with token
//// @Tags auth-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestSignInGoogle true "account"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /auth/sign-in-google [post]
//func (authReceiver *AuthController) HandleSignInGoogleToken(c echo.Context) error {
//	user := c.Get("user").(*jwt.Token)
//	claims := user.Claims.(model.JwtCustomClaims)
//	email := claims.Email
//	//oauthGoogleServiceInstance := services.GetOauth2ServiceInstance()
//	//dataContent := oauthGoogleServiceInstance.GetUserInfoWithToken(reqSignInGoogle.IDToken)
//	result, err := authReceiver.AccountRepo.CheckEmailExisted(email)
//
//	if result.ID == "-1" {
//		return c.JSON(http.StatusInternalServerError, res.Response{
//			StatusCode: http.StatusInternalServerError,
//			Message:    "Tài khoản không tồn tại",
//		})
//	}
//	//generate token
//	_, err = security.GenToken(&result)
//	if err != nil {
//		logger.Error("err gen token", zap.Error(err))
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	_, _, err = security.GenRefToken(&result)
//	if err != nil {
//		logger.Error("err gen token data", zap.Error(err))
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//
//	return response.Ok(c, "Đăng nhập thành công", result)
//
//}

// HandleGenerateNewAccessToken godoc
// @Summary Create access token by refresh token
// @Tags auth-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /auth/refresh-token [get]
func (authReceiver *AuthController) HandleGenerateNewAccessToken(c echo.Context) error {
	reqRefreshToken := req.RequestRefreshToken{}
	if err := c.Bind(&reqRefreshToken); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, "Thông tin không hợp lệ", nil)
	}
	claims := jwt.MapClaims{}
	tokenResult, err := jwt.ParseWithClaims(reqRefreshToken.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_REFRESH_KEY")), nil
	})
	var userId string
	for key, val := range claims {
		if key == "UserId" {
			userId = val.(string)
		}
	}
	if tokenResult.Valid {
		user, err := authReceiver.AccountRepo.GetAccountById(userId)

		if err != nil {
			return response.BadRequest(c, err.Error(), nil)
		}

		//generate token
		_, err = security.GenToken(&user)
		if err != nil {
			logger.Error("err gen token", zap.Error(err))
			return response.InternalServerError(c, err.Error(), nil)
		}
		_, _, err = security.GenRefToken(&user)
		if err != nil {
			logger.Error("err gen token data", zap.Error(err))
			return response.InternalServerError(c, err.Error(), nil)
		}
		token := model.Token{
			AccessToken:  user.Token.AccessToken,
			RefreshToken: user.Token.RefreshToken,
			ExpiredTime:  user.Token.ExpiredTime,
		}
		return response.Ok(c, "Tạo token thành công", token)
	}
	logger.Error("err gen token data", zap.Error(err))
	return response.InternalServerError(c, err.Error(), nil)
}

//// HandleAuthenticationGoogleWithInfo godoc
//// @Summary Handle Sign in google with info
//// @Tags auth-service
//// @Accept  json
//// @Produce  json
//// @Param data body req.RequestGetOauthInfo true "account"
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 403 {object} res.Response
//// @Failure 409 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /auth/sign-in-oauth-info [post]
//func (authReceiver *AuthController) HandleAuthenticationGoogleWithInfo(c echo.Context) error {
//	reqGoogleInfo := req.RequestGetOauthInfo{}
//	if err := c.Bind(&reqGoogleInfo); err != nil {
//		logger.Error("Error binding data", zap.Error(err))
//		return response.BadRequest(c, err.Error(), nil)
//	}
//	if !reqGoogleInfo.EmailVerified {
//		return c.JSON(http.StatusBadRequest, res.Response{
//			StatusCode: http.StatusBadRequest,
//			Message:    "Tài khoản không hợp lệ",
//		})
//	}
//	reqSignIn := req.RequestSignIn{
//		Email: reqGoogleInfo.Email,
//	}
//	account, err := authReceiver.AccountRepo.CheckLogin( /*c.Request().Context(), */ reqSignIn)
//	if err == custom_error.UserNotFound {
//		reqRegister := req.RequestSignUp{
//			Email:     reqGoogleInfo.Email,
//			FirstName: reqGoogleInfo.GivenName,
//			LastName:  reqGoogleInfo.FamilyName,
//			Password:  "123456",
//		}
//		//Generate UUID
//		customerId, err := uuid.NewUUID()
//		if err != nil {
//			logger.Error("Error uuid data", zap.Error(err))
//			return response.Forbidden(c, err.Error(), nil)
//		}
//		customer := model.Customer{
//			ID:         customerId.String(),
//			CustomerID: customerId.String(),
//			Email:      reqRegister.Email,
//			FirstName:  reqRegister.FirstName,
//			LastName:   reqRegister.LastName,
//			FullName:   reqRegister.FirstName + " " + reqRegister.LastName,
//		}
//		customer, err = authReceiver.AccountRepo.SaveProfileCustomer(customer)
//		if err != nil {
//			return c.JSON(http.StatusOK, res.Response{
//				StatusCode: http.StatusOK,
//				Message:    err.Error(),
//				Data:       nil,
//			})
//		}
//		//create password
//		hash := security.HashAndSalt([]byte(reqRegister.Password))
//		//Generate UUID
//		accountId, err := uuid.NewUUID()
//		if err != nil {
//			logger.Error("Error uuid data", zap.Error(err))
//			return response.Forbidden(c, err.Error(), nil)
//		}
//		//Init account
//		account := model.Account{
//			ID:              accountId.String(),
//			Email:           reqRegister.Email,
//			Password:        hash,
//			CustomerID:      customerId.String(),
//			CreatedAt:       time.Now(),
//			UpdatedAt:       time.Now(),
//			StatusAccountID: 1,
//		}
//
//		//Save account
//		//account, err = authReceiver.AccountRepo.SaveAccount(account)
//		if err != nil {
//			return response.Conflict(c, err.Error(), nil)
//		}
//		return c.JSON(http.StatusOK, res.Response{
//			StatusCode: http.StatusOK,
//			Message:    "Đăng ký thành công",
//			Data:       account,
//		})
//	}
//	if account.StatusAccountID == 0 {
//		return c.JSON(http.StatusUnauthorized, res.Response{
//			StatusCode: http.StatusUnauthorized,
//			Message:    "Tài khoản bị khóa",
//			Data:       nil,
//		})
//	}
//	//generate token
//	_, err = security.GenToken(&account)
//	if err != nil {
//		logger.Error("err gen token", zap.Error(err))
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	_, _, err = security.GenRefToken(&account)
//	if err != nil {
//		logger.Error("err gen token data", zap.Error(err))
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Đăng nhập thành công", account)
//}
//
//// HandleGetMe godoc
//// @Summary Get me
//// @Tags auth-service
//// @Accept  json
//// @Produce  json
//// @Success 200 {object} res.Response
//// @Failure 400 {object} res.Response
//// @Failure 500 {object} res.Response
//// @Router /auth/me [get]
//func (authReceiver *AuthController) HandleGetMe(c echo.Context) error {
//	claims := security.GetClaimsJWT(c)
//	account, err := authReceiver.AccountRepo.GetAccountById(claims.UserId)
//	if err != nil {
//		return response.InternalServerError(c, err.Error(), nil)
//	}
//	return response.Ok(c, "Lấy thông tin thành công", account)
//}
//
//func (authReceiver *AuthController) HandleSignInGoogle(c echo.Context) error {
//	return response.BadRequest(c, "Hello", nil)
//}
