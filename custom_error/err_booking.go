package custom_error

import "errors"

var (
	BookingConflict     = errors.New("Booking đã tồn tại")
	BookingNotFound     = errors.New("Booking không tồn tại")
	BookingNotUpdated   = errors.New("Cập nhật thông tin Booking thất bại")
	BookingDeleteFailed = errors.New("Xóa Booking thất bại")
	BookingNotSaved     = errors.New("Lưu Booking thất bại")
	BookingsEmpty       = errors.New("Trống")
)
