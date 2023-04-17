package custom_error

import "errors"

var (
	PaymentConflict     = errors.New("Payment đã tồn tại")
	PaymentNotFound     = errors.New("Payment không tồn tại")
	PaymentNotUpdated   = errors.New("Cập nhật thông tin Payment thất bại")
	PaymentDeleteFailed = errors.New("Xóa Payment thất bại")
	PaymentNotSaved     = errors.New("Lưu Payment thất bại")
	PaymentsEmpty       = errors.New("Trống")
)
