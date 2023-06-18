package services

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/robfig/cron.v2"
	"hotel-booking-api/repository"
)

type SchedulerGoCronService struct {
	Echo     *echo.Echo
	Cron     *cron.Cron
	RoomRepo repository.RoomRepo
}

func (service *SchedulerGoCronService) InitSchedulerGoCronLockRoom() {
	_, _ = service.Cron.AddFunc("@every 0h5m0s", service.RoomRepo.UpdateLockRoom)
	service.Cron.Start()
}
