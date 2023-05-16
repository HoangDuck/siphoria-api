package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/sony/sonyflake"
	"log"
	"net/http"
	"strconv"
)

type MomoService struct {
}

type Payload struct {
	PartnerCode  string `json:"partnerCode"`
	AccessKey    string `json:"accessKey"`
	RequestID    string `json:"requestId"`
	Amount       int    `json:"amount"`
	OrderID      string `json:"orderId"`
	OrderInfo    string `json:"orderInfo"`
	PartnerName  string `json:"partnerName"`
	StoreId      string `json:"storeId"`
	OrderGroupId string `json:"orderGroupId"`
	Lang         string `json:"lang"`
	AutoCapture  bool   `json:"autoCapture"`
	RedirectUrl  string `json:"redirectUrl"`
	IpnUrl       string `json:"ipnUrl"`
	ExtraData    string `json:"extraData"`
	RequestType  string `json:"requestType"`
	Signature    string `json:"signature"`
}

type PosHash struct {
	PartnerCode  string `json:"partnerCode"`
	PartnerRefID string `json:"partnerRefId"`
	Amount       int    `json:"amount"`
	PaymentCode  string `json:"paymentCode"`
}

type PosPayload struct {
	PartnerCode  string `json:"partnerCode"`
	PartnerRefID string `json:"partnerRefId"`
	Hash         string `json:"hash"`
	Version      int    `json:"version"`
}

var momoService *MomoService

func GetMomoServiceInstance() *MomoService {
	if momoService == nil {
		momoService = new(MomoService)
	}
	return momoService
}

func (service MomoService) PaymentService(condition map[string]interface{}) map[string]interface{} {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	//randome orderID and requestID
	b, _ := flake.NextID()

	var endpoint = "https://test-payment.momo.vn/v2/gateway/api/create"
	var accessKey = "yFRGoK0eLSrthX4Y"
	var secretKey = "tZNafmaHgldR8XfZA9wiYCFIkaXbzxbu"
	var orderInfo = fmt.Sprint(condition["booking-description"])
	var partnerCode = "MOMOQDD420220927"
	var redirectUrl = fmt.Sprint(condition["redirect-url"])
	var ipnUrl = fmt.Sprint(condition["ipn-url"]) + "/api/payment/result-momo"
	intVar, _ := strconv.Atoi(fmt.Sprint(condition["amount"]))
	var amount = intVar
	var orderId = fmt.Sprint(condition["booking-info"]) + "_" + fmt.Sprint(condition["payment_id"])
	var requestId = strconv.FormatUint(b, 16)
	var extraData = ""
	var partnerName = "Siphoria"
	var storeId = "Test Store"
	var orderGroupId = ""
	var lang = "vi"
	var requestType = "captureWallet"

	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("accessKey=")
	rawSignature.WriteString(accessKey)
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(fmt.Sprint(condition["amount"]))
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)
	rawSignature.WriteString("&ipnUrl=")
	rawSignature.WriteString(ipnUrl)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(orderId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(partnerCode)
	rawSignature.WriteString("&redirectUrl=")
	rawSignature.WriteString(redirectUrl)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(requestId)
	rawSignature.WriteString("&requestType=")
	rawSignature.WriteString(requestType)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmacSignature := hmac.New(sha256.New, []byte(secretKey))

	// Write Data to it
	hmacSignature.Write(rawSignature.Bytes())
	fmt.Println("Raw signature: " + rawSignature.String())

	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmacSignature.Sum(nil))

	var payload = Payload{
		PartnerCode: partnerCode,
		//AccessKey:    accessKey,
		RequestID:    requestId,
		Amount:       amount,
		RequestType:  requestType,
		RedirectUrl:  redirectUrl,
		IpnUrl:       ipnUrl,
		OrderID:      orderId,
		StoreId:      storeId,
		PartnerName:  partnerName,
		OrderGroupId: orderGroupId,
		AutoCapture:  true,
		Lang:         lang,
		OrderInfo:    orderInfo,
		ExtraData:    extraData,
		Signature:    signature,
	}

	var jsonPayload []byte
	var err error
	jsonPayload, err = json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Request to Momo: ", payload)

	///send HTTP to momo endpoint
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}

	///result
	var result map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Response from Momo: ", result)
	return result
}
