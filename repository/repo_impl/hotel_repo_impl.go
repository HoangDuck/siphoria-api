package repo_impl

import (
	"hotel-booking-api/db"
	"hotel-booking-api/repository"
)

type HotelRepoImpl struct {
	sql *db.Sql
}

func NewHotelRepo(sql *db.Sql) repository.HotelRepo {
	return &HotelRepoImpl{
		sql: sql,
	}
}
