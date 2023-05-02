package controller

import (
	"github.com/labstack/echo/v4"
	response "hotel-booking-api/model/model_func"
)

type MapModel struct {
	Value string `json:"data"`
}

func GetEmbeddedMap(c echo.Context) error {
	
	return response.Ok(c, "success", MapModel{
		Value: `<iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3919.2228246450686!2d106.6499702!3d10.7942387!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x3175294af9f6f929%3A0x2f07b2a4ec91d9c8!2zOTcgVHLGsOG7nW5nIENoaW5oLCBQaMaw4budbmcgMTIsIFTDom4gQsOsbmgsIFRow6BuaCBwaOG7kSBI4buTIENow60gTWluaA!5e0!3m2!1svi!2s!4v1682957889463!5m2!1svi!2s" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>`,
	})
}
