package custom_error

import "errors"

var (
	HotelConflict   = errors.New("Khách sạn đã tồn tại")
	HotelNotFound   = errors.New("Khách sạn không tồn tại")
	HotelNotUpdated = errors.New("cập nhật thông tin khách sạn thất bại")
	//UploadFail         = errors.New("That bai")
	//CreateUserError    = errors.New("Tạo profile bị sự cố")
)
