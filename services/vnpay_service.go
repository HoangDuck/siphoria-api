package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
	"time"
)

type VNPayService struct {
}

var vnpayService *VNPayService

func GetVNPayServiceInstance() *VNPayService {
	if vnpayService == nil {
		vnpayService = new(VNPayService)
	}
	return vnpayService
}

func (service VNPayService) VNPayPaymentService(condition map[string]interface{}) string {
	intVar, _ := strconv.Atoi(fmt.Sprint(condition["amount"]))
	var vnp_Version = "2.1.0"
	var vnp_Command = "pay"
	var vnp_TmnCode = ConfigInfo.VNPay.VNPTmnCode
	var vnp_Locale = "vn"
	var vnp_CurrCode = "VND"
	var vnp_TxnRef = fmt.Sprint(condition["booking-info"]) + "_" + fmt.Sprint(condition["payment_id"])
	var vnp_OrderInfo = fmt.Sprint(condition["booking-description"])
	var vnp_OrderType = "other"
	var vnp_Amount = intVar * 100
	var vnp_ReturnUrl = fmt.Sprint(condition["ipn-url"])
	now := time.Now()
	var vnp_CreateDate = now.Format("20060102150405")
	var vnp_IpAddr = "115.73.215.9"
	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("vnp_Amount=")
	rawSignature.WriteString(fmt.Sprint(vnp_Amount))
	rawSignature.WriteString("&vnp_Command=")
	rawSignature.WriteString(vnp_Command)
	rawSignature.WriteString("&vnp_CreateDate=")
	rawSignature.WriteString(vnp_CreateDate)
	rawSignature.WriteString("&vnp_CurrCode=")
	rawSignature.WriteString(vnp_CurrCode)
	rawSignature.WriteString("&vnp_IpAddr=")
	rawSignature.WriteString(vnp_IpAddr)
	rawSignature.WriteString("&vnp_Locale=")
	rawSignature.WriteString(vnp_Locale)
	rawSignature.WriteString("&vnp_OrderInfo=")
	rawSignature.WriteString(vnp_OrderInfo)
	rawSignature.WriteString("&vnp_OrderType=")
	rawSignature.WriteString(vnp_OrderType)
	rawSignature.WriteString("&vnp_ReturnUrl=")
	rawSignature.WriteString(vnp_ReturnUrl)
	rawSignature.WriteString("&vnp_TmnCode=")
	rawSignature.WriteString(vnp_TmnCode)
	rawSignature.WriteString("&vnp_TxnRef=")
	rawSignature.WriteString(vnp_TxnRef)
	rawSignature.WriteString("&vnp_Version=")
	rawSignature.WriteString(vnp_Version)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmacSignature := hmac.New(sha512.New, []byte(ConfigInfo.VNPay.VNPHashSecret))

	// Write Data to it
	hmacSignature.Write(rawSignature.Bytes())
	fmt.Println("Raw signature: " + rawSignature.String())

	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmacSignature.Sum(nil))

	rawSignature.WriteString("&vnp_SecureHash=")
	rawSignature.WriteString(signature)

	///send HTTP to vnpay endpoint
	return rawSignature.String()
}

func (service VNPayService) CheckSignatureResultVNPay(c echo.Context) bool {
	dataFromVNPay := c.QueryParams()
	secureHash := dataFromVNPay.Get("vnp_SecureHash")
	dataFromVNPay.Del("vnp_SecureHash")
	var rawSignature bytes.Buffer
	rawSignature.WriteString(strings.Replace(dataFromVNPay.Encode(), "\u0026", "&", -1))
	hmacSignature := hmac.New(sha512.New, []byte(ConfigInfo.VNPay.VNPHashSecret))

	// Write Data to it
	hmacSignature.Write(rawSignature.Bytes())
	fmt.Println("Raw signature: " + rawSignature.String())
	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmacSignature.Sum(nil))
	if signature == secureHash {
		return true
	}
	return false
}
