package repo_impl

import (
	"hotel-booking-api/db"
	"hotel-booking-api/model"
	"hotel-booking-api/model/req"
	"hotel-booking-api/repository"
)

type HotelRepoImpl struct {
	sql *db.Sql
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
