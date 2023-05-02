package custom_error

import "errors"

var (
	RateplanConflict     = errors.New("Rateplan đã tồn tại")
	RateplanNotFound     = errors.New("Rateplan không tồn tại")
	RateplanNotUpdated   = errors.New("Cập nhật thông tin rateplan thất bại")
	RateplanDeleteFailed = errors.New("Xóa rateplan thất bại")
)
