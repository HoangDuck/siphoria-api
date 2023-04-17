package custom_error

import "errors"

var (
	UserConflict   = errors.New("người dùng đã tồn tại")
	UserNotFound   = errors.New("người dùng không tồn tại")
	UserNotUpdated = errors.New("cập nhật thông tin người dùng thất bại")
	// EmailAlreadyExists SignUpFail         = errors.New("Đăng ký thất bại")
	EmailAlreadyExists = errors.New("email đã tồn tại")
	//UploadFail         = errors.New("That bai")
	//CreateUserError    = errors.New("Tạo profile bị sự cố")
)
