package services

type PaypalService struct {
}

var paypalService *PaypalService

func GetPaypalService() *PaypalService {
	if paypalService == nil {
		paypalService = new(PaypalService)
	}
	return paypalService
}
