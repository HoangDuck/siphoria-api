package repo_impl

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hotel-booking-api/db"
	"hotel-booking-api/logger"
	"hotel-booking-api/model"
	"hotel-booking-api/repository"
	"time"
)

type VoucherRepoImpl struct {
	sql *db.Sql
}

func (voucherReceiver *VoucherRepoImpl) SaveBatchVoucher(listRoomTypeId []string, voucherId string) ([]model.VoucherExcept, error) {
	var listTempVoucherExcept []model.VoucherExcept
	listTempVoucherExcept = []model.VoucherExcept{}
	for index := 0; index < len(listRoomTypeId); index++ {
		listTempVoucherExcept = append(listTempVoucherExcept, model.VoucherExcept{
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			VoucherId:  voucherId,
			RoomTypeId: listRoomTypeId[index],
			IsDeleted:  false,
		})
	}
	logger.Info(listTempVoucherExcept[0].RoomTypeId)
	err := voucherReceiver.sql.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "voucher_id"}, {Name: "room_type_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at", "is_deleted"}),
	}).Create(&listTempVoucherExcept).Error
	if err != nil {
		return listTempVoucherExcept, err
	}
	return listTempVoucherExcept, nil
}

func (voucherReceiver *VoucherRepoImpl) DeleteVoucher(voucher model.Voucher) (bool, error) {
	err := voucherReceiver.sql.Db.Select("is_deleted").Model(&voucher).Updates(voucher)
	if err.Error != nil {
		logger.Error("Error update user failed ", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return false, err.Error
		}
		return false, err.Error
	}
	return true, nil
}

func (voucherReceiver *VoucherRepoImpl) UpdateVoucher(voucher model.Voucher) (model.Voucher, error) {
	err := voucherReceiver.sql.Db.Model(&voucher).Updates(voucher)
	if err.Error != nil {
		logger.Error("Error update user failed ", zap.Error(err.Error))
		if err.Error == gorm.ErrRecordNotFound {
			return voucher, err.Error
		}
		return voucher, err.Error
	}
	return voucher, nil
}

func (voucherReceiver *VoucherRepoImpl) SaveVoucher(voucher model.Voucher) (model.Voucher, error) {
	result := voucherReceiver.sql.Db.Create(&voucher)
	if result.Error != nil {
		return voucher, result.Error
	}
	return voucher, nil
}

func NewVoucherRepo(sql *db.Sql) repository.VoucherRepo {
	return &VoucherRepoImpl{
		sql: sql,
	}
}
