package repository

import "hotel-booking-api/model"

type VoucherRepo interface {
	SaveVoucher(voucher model.Voucher) (model.Voucher, error)
	UpdateVoucher(voucher model.Voucher) (model.Voucher, error)
}
