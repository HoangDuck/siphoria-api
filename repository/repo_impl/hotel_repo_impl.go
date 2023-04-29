package repo_impl

import (
	"hotel-booking-api/db"
	"hotel-booking-api/model"
	"hotel-booking-api/repository"
)

type HotelRepoImpl struct {
	sql *db.Sql
}

func (h *HotelRepoImpl) UpdateHotelBusinessLicensePhotos(hotel model.Hotel) (model.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepoImpl) UpdateHotelPhotos(hotel model.Hotel) (model.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func NewHotelRepo(sql *db.Sql) repository.HotelRepo {
	return &HotelRepoImpl{
		sql: sql,
	}
}
