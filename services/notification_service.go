package services

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/api/option"
	"hotel-booking-api/logger"
	response "hotel-booking-api/model/model_func"
	"hotel-booking-api/model/req"
	"log"
)

type NotificationService struct {
	Echo *echo.Echo
}

var client *messaging.Client

func GetClientFirebase() {
	if client == nil {
		InitFirebase()
	}
}

func InitFirebase() {
	opt := option.WithCredentialsFile("../fir-golang-test-firebase-adminsdk-32eya-d01adefe0e.json")
	config := &firebase.Config{ProjectID: "fir-golang-test"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err = app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
}

func PushNotificationMessage(FcmKey string, Message map[string]string) string {
	GetClientFirebase()

	registrationToken := FcmKey

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data:  Message,
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	responseMessage, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	log.Fatal(responseMessage)
	return responseMessage
}

func (notification *NotificationService) PushNotificationMessageAPITest(c echo.Context) error {
	GetClientFirebase()
	reqMessageNotification := req.RequestMessageNotification{}
	if err := c.Bind(&reqMessageNotification); err != nil {
		logger.Error("Error binding data", zap.Error(err))
		return response.BadRequest(c, err.Error(), nil)
	}
	registrationToken := reqMessageNotification.FCMKey

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"title": "Test message",
			"body":  "Hello to Final Project Khang & Duc",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	responseMessage, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	return response.Ok(c, responseMessage, message.Data)
}
