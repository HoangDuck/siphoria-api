package response

import (
	"github.com/labstack/echo/v4"
	"hotel-booking-api/model/res"
	"net/http"
)

func Ok(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	})
}

func NotFound(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusNotFound, res.Response{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Data:       data,
	})
}

func BadRequest(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusBadRequest, res.Response{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Data:       data,
	})
}

func InternalServerError(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusInternalServerError, res.Response{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Data:       data,
	})
}

func Unauthorized(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusUnauthorized, res.Response{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
		Data:       data,
	})
}

func Conflict(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusConflict, res.Response{
		StatusCode: http.StatusConflict,
		Message:    message,
		Data:       data,
	})
}

func Forbidden(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusForbidden, res.Response{
		StatusCode: http.StatusForbidden,
		Message:    message,
		Data:       data,
	})
}

func UnprocessableEntity(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusUnprocessableEntity, res.Response{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
		Data:       data,
	})
}

func RequestTimedOut(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusRequestTimeout, res.Response{
		StatusCode: http.StatusRequestTimeout,
		Message:    message,
		Data:       data,
	})
}

func NotImplemented(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusNotImplemented, res.Response{
		StatusCode: http.StatusNotImplemented,
		Message:    message,
		Data:       data,
	})
}

func BadGateWay(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusBadGateway, res.Response{
		StatusCode: http.StatusBadGateway,
		Message:    message,
		Data:       data,
	})
}
