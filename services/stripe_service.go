package services

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v74/webhook"
	"hotel-booking-api/logger"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type StripeService struct {
}

var stripeService *StripeService

func GetStripeServiceInstance() *StripeService {
	if stripeService == nil {
		stripeService = new(StripeService)
	}
	return stripeService
}

func (service StripeService) CreatePaymentStripe(total int64) string {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(total),
		Currency: stripe.String(string(stripe.CurrencyVND)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("pi.New: %v", err)
		return ""
	}
	return pi.ClientSecret
}

func (service StripeService) CheckWebHookEvent(c echo.Context) string {
	endpointSecret := "whsec_07036e7edd6b8492664fab2ff157fc6e281f15cf489cabf87851c980461adb5f"
	signatureHeader := c.Request().Header.Get("Stripe-Signature")
	const MaxBodyBytes = int64(65536)
	requestBody := http.MaxBytesReader(c.Response(), c.Request().Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(requestBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)

		return ""
	}

	event, err := webhook.ConstructEvent(payload, signatureHeader, endpointSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Webhook signature verification failed. %v\n", err)
		return ""
	}
	logger.Infof("event type", event.Type)
	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)

			return "payment_intent.succeeded"
		}
		log.Printf("Successful payment for %d.", paymentIntent.Amount)
		// Then define and call a func to handle the successful payment intent.
		// handlePaymentIntentSucceeded(paymentIntent)
	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)

			return "payment_method.attached"
		}
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handlePaymentMethodAttached(paymentMethod)
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}
	return ""
}
