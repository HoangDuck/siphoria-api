package res

type DataPaymentRes struct {
	Amount       int    `json:"amount"`
	Message      string `json:"message"`
	OrderID      string `json:"orderId"`
	PartnerCode  string `json:"partnerCode"`
	PayURL       string `json:"payUrl"`
	RequestID    string `json:"requestId"`
	ResponseTime int64  `json:"responseTime"`
	ResultCode   int    `json:"resultCode"`
}
