package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"hotel-booking-api/logger"
	"io/ioutil"
	"os"
)

type LogsHandler struct {
}

func (log *LogsHandler) CheckLogs(c echo.Context) error {
	filePath := "../cmd/go.log"
	fmt.Println("loading log file from ", filePath)

	f, err := os.Open(filePath)
	if err != nil {
		logger.Error("Error open file", zap.Error(err))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	textFile, err := ioutil.ReadFile(filePath)
	return c.String(200, string(textFile))
}

func (log *LogsHandler) CheckLogsSystem(c echo.Context) error {
	filePath := "../go.log"
	fmt.Println("loading log file from ", filePath)

	f, err := os.Open(filePath)
	if err != nil {
		logger.Error("Error open file", zap.Error(err))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	textFile, err := ioutil.ReadFile(filePath)
	return c.String(200, string(textFile))
}
