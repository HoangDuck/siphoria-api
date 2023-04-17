package custom_error

import "errors"

var (
	RoomConflict     = errors.New("Phòng đã tồn tại")
	RoomNotFound     = errors.New("Phòng không tồn tại")
	RoomNotUpdated   = errors.New("Cập nhật thông tin phòng thất bại")
	RoomDeleteFailed = errors.New("Xóa phòng thất bại")
)
