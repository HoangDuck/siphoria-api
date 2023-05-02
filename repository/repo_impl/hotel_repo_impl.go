package repo_impl

import (
	"gorm.io/gorm"
	"hotel-booking-api/custom_error"
	"hotel-booking-api/db"
	"hotel-booking-api/model"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
)

type HotelRepoImpl struct {
	sql *db.Sql
}

func (hotelReceiver *HotelRepoImpl) UpdateHotel(requestUpdateHotel req.RequestUpdateHotel, idHotel string) (model.Hotel, error) {
	hotel := model.Hotel{
		ID:              idHotel,
		Name:            requestUpdateHotel.Name,
		Overview:        requestUpdateHotel.Overview,
		Activate:        requestUpdateHotel.Activate,
		ProvinceCode:    requestUpdateHotel.Province,
		DistrictCode:    requestUpdateHotel.District,
		WardCode:        requestUpdateHotel.Ward,
		RawAddress:      requestUpdateHotel.RawAddress,
		BankAccount:     requestUpdateHotel.BankAccount,
		BankBeneficiary: requestUpdateHotel.BankBeneficiary,
		BankName:        requestUpdateHotel.BankName,
	}
	hotelType := model.HotelType{
		HotelId:   idHotel,
		Hotel:     requestUpdateHotel.Hotel,
		Apartment: requestUpdateHotel.Apartment,
		Resort:    requestUpdateHotel.Resort,
		Villa:     requestUpdateHotel.Villa,
		Camping:   requestUpdateHotel.Camping,
		Motel:     requestUpdateHotel.Motel,
		HomeStay:  requestUpdateHotel.Homestay,
	}
	hotelFacility := model.HotelFacility{
		HotelId:       idHotel,
		Beach:         requestUpdateHotel.Beach,
		Pool:          requestUpdateHotel.Pool,
		Bar:           requestUpdateHotel.Bar,
		NoSmokingRoom: requestUpdateHotel.NoSmokingRoom,
		Fitness:       requestUpdateHotel.Fitness,
		Spa:           requestUpdateHotel.Spa,
		Bath:          requestUpdateHotel.Bath,
		Wifi:          requestUpdateHotel.Wifi,
		Breakfast:     requestUpdateHotel.Breakfast,
		Casino:        requestUpdateHotel.Casio,
		Parking:       requestUpdateHotel.Parking,
	}
	err := hotelReceiver.sql.Db.Model(&hotel).Updates(hotel)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}

		return hotel, custom_error.HotelNotUpdated
	}
	err = hotelReceiver.sql.Db.Model(&hotelType).Updates(hotelType)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}

		return hotel, custom_error.HotelNotUpdated
	}
	err = hotelReceiver.sql.Db.Model(&hotelFacility).Updates(hotelFacility)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return hotel, err.Error
		}

		return hotel, custom_error.HotelNotUpdated
	}
	return hotel, nil
}

func (hotelReceiver *HotelRepoImpl) SaveHotel(requestAddHotel req.RequestCreateHotel) (model.Hotel, error) {
	err := hotelReceiver.sql.Db.Exec("call sp_addhotel(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);",
		requestAddHotel.ID,
		requestAddHotel.HotelierID,
		requestAddHotel.Name,
		requestAddHotel.Overview,
		requestAddHotel.Activate,
		requestAddHotel.Photos,
		requestAddHotel.RawAddress,
		requestAddHotel.BankAccount,
		requestAddHotel.BankName,
		requestAddHotel.BankBeneficiary,
		requestAddHotel.Hotel,
		requestAddHotel.Apartment,
		requestAddHotel.Resort,
		requestAddHotel.Villa,
		requestAddHotel.Camping,
		requestAddHotel.Motel,
		requestAddHotel.Homestay,
		requestAddHotel.Beach,
		requestAddHotel.Spa,
		requestAddHotel.Pool,
		requestAddHotel.Bar,
		requestAddHotel.NoSmokingRoom,
		requestAddHotel.Fitness,
		requestAddHotel.Bath,
		requestAddHotel.Wifi,
		requestAddHotel.Breakfast,
		requestAddHotel.Casino,
		requestAddHotel.Parking,
		requestAddHotel.District,
		requestAddHotel.Province,
		requestAddHotel.Ward,
		requestAddHotel.BusinessLicense)
	var hotel model.Hotel
	if err.Error != nil {
		return hotel, err.Error
	} else {
		hotel = model.Hotel{
			ID:              requestAddHotel.ID,
			Name:            requestAddHotel.Name,
			Overview:        requestAddHotel.Overview,
			Activate:        requestAddHotel.Activate,
			ProvinceCode:    requestAddHotel.Province,
			DistrictCode:    requestAddHotel.District,
			WardCode:        requestAddHotel.Ward,
			RawAddress:      requestAddHotel.RawAddress,
			HotelPhotos:     requestAddHotel.Photos,
			BankAccount:     requestAddHotel.BankAccount,
			BankBeneficiary: requestAddHotel.BankBeneficiary,
			BankName:        requestAddHotel.BankName,
			BusinessLicence: requestAddHotel.BusinessLicense,
			HotelierId:      requestAddHotel.HotelierID,
		}
	}
	return hotel, nil
}

func (hotelReceiver *HotelRepoImpl) CreateRequestPayout(payoutRequest model.PayoutRequest, paymentIds string) (model.PayoutRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (hotelReceiver *HotelRepoImpl) UpdateHotelBusinessLicensePhotos(hotel model.Hotel) (model.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (hotelReceiver *HotelRepoImpl) UpdateHotelPhotos(hotel model.Hotel) (model.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func NewHotelRepo(sql *db.Sql) repository.HotelRepo {
	return &HotelRepoImpl{
		sql: sql,
	}
}
