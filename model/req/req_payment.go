package req

type RequestPaymentList struct {
	CustomerName  string  `json:"customer_name"`
	Amount        float32 `json:"amount" validate:"required,min=0"`
	PaymentMethod string  `json:"payment_method"`
	PaymentTime   string  `json:"payment_time"`
	StatusPayment string  `json:"status_payment"`
}

type RequestGetHistoryPaymentList struct {
	CustomerID string `json:"customer_id"`
}

type RequestGetBillPayment struct {
	PaymentID string `json:"payment_id"`
	BookingID string `json:"booking_id"`
}

type RequestCreatePayment struct {
	CustomerID string `json:"customer_id"`
	BookingID  string `json:"booking_id"`
	Amount     int    `json:"amount"`
	//PaymentMethod string  `json:"payment_method"`
	Description string `json:"description"`
}

type RequestUpdatePayment struct {
	PaymentID      string  `json:"payment_id"`
	FineAmount     float32 `json:"fine_amount"`
	Amount         int     `json:"amount"`
	DueTimePayment string  `json:"due_time_payment"`
	PaymentMethod  string  `json:"payment_method"`
	PaymentTime    string  `json:"payment_time"`
	StatusPayment  string  `json:"status_payment"`
}

type RequestCancelPayment struct {
	PaymentID string `json:"payment_id"`
}

type RequestAddPaymentMethod struct {
	MethodName  string `json:"method_name"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
}

type RequestDeletePayment struct {
	PaymentID string `json:"payment_id"`
}

type RequestReceiveResultPayment struct {
	PaymentID string `json:"payment_id"`
}

type RequestCreatePaymentBill struct {
	BookingID string `json:"booking_id"`
	PaymentID string `json:"payment_id"`
}

type RequestStatisticRevenueByDay struct {
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
	Mode      int    `json:"mode"`
}

type RequestAddMoneyTopUp struct {
	Amount float32 `json:"amount"`
	Method string  `json:"method"`
}
