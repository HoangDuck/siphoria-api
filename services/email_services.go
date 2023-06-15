package services

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"hotel-booking-api/model"
)

type EmailService struct {
	Sender         string
	SenderPassword string
}

var emailService *EmailService

func GetEmailServiceInstance() *EmailService {
	if emailService == nil {
		emailService = new(EmailService)
		emailService.SetUpConfig(*ConfigInfo)
	}
	return emailService
}

func (service *EmailService) CreateBodyTemplate(content string, buttonName string, url string) string {
	return "<div style=\"width: 100%;align-items: center;\">" +
		"<br><img src=\"https://res.cloudinary.com/dqqdwydoo/image/upload/v1686791558/353946686_220191677493448_7398544158072599447_n_w9v97l.png\" " +
		"alt=\"Siphoria Logo\" style=\"max-width: 100px;user-select: none;\" />" +
		"<br><p> " + content + " </p>" +
		"<br><button style=\"padding: 15px 20px;border-radius: 8px;background-color: #22c55e;color: white;font-weight: 600;border: none;cursor: pointer;\" " +
		"><a style=\"text-decoration: none; color: white;\" href=\"" + url + "\">" + buttonName + "</a> </button>" +
		"<br><p>Nếu không phải là bạn yêu cầu, xin vui lòng bỏ qua email</p>" +
		"<br><div style=\"text-align: center;\"> Cảm ơn, " +
		"<div>Đội ngũ Siphoria</div> </div>    " +
		"</div>"
}

func (service *EmailService) SetUpConfig(cfg model.Config) {
	service.Sender = cfg.Email.SenderEmail
	service.SenderPassword = cfg.Email.SenderPass
	if service.Sender != "" && service.SenderPassword != "" {
		fmt.Println("Email services connected")
	} else {
		fmt.Println("Email services connect failed")
	}
}

func (service EmailService) SendEmailService(Receiver []string, Body string, Subject string) error {
	from := service.Sender
	password := service.SenderPassword
	toList := Receiver
	host := "smtp.gmail.com"
	port := 587
	msg := Body
	//body := []byte(msg)
	mailSend := gomail.NewMessage()
	mailSend.SetHeader("From", from)
	mailSend.SetHeader("To", toList[0])
	mailSend.SetHeader("Subject", Subject)
	mailSend.SetBody("text/html", msg)
	d := gomail.NewDialer(host, port, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Now send E-Mail
	if err := d.DialAndSend(mailSend); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return nil
}
