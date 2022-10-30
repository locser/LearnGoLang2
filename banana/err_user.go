package banana

import "errors"

var (
	UserConflict = errors.New("Lỗi: người dùng đã tồn tại")
	SignUpFail   = errors.New("Đăng kí thất bại")
)
