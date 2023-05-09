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
	"hotel-booking-api/services"
)

type NotificationController struct {
	NotificationRepo repository.NotificationRepo
}

// PushNotificationMessageAPIAdmin godoc
// @Summary Admin push notifications
// @Tags admin-service
// @Accept  json
// @Produce  json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/push-notification [post]
func (notificationReceiver *NotificationController) PushNotificationMessageAPIAdmin(c echo.Context) error {
	services.GetClientFirebase()
	reqMessageNotification := req.RequestPushNotificationAdmin{}
	if err := c.Bind(&reqMessageNotification); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	//Validate
	if err := c.Validate(reqMessageNotification); err != nil {
		logger.Error("Error validate data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	registrationToken := reqMessageNotification.FCMKey
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	if !(security.CheckRole(claims, model.ADMIN, false)) {
		return response.BadRequest(c, "Bạn không có quyền thực hiện chức năng này", nil)
	}

	responseMessage := services.PushNotificationMessage(registrationToken, map[string]string{
		"title": reqMessageNotification.Title,
		"body":  reqMessageNotification.Description,
	})

	// Response is a message ID string.
	return response.Ok(c, "Thông báo", responseMessage)
}
