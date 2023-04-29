package repo_impl

import (
	"hotel-booking-api/db"
	"hotel-booking-api/repository"
)

type NotificationRepoImpl struct {
	sql *db.Sql
}

func NewNotificationRepo(sql *db.Sql) repository.NotificationRepo {
	return &NotificationRepoImpl{
		sql: sql,
	}
}
