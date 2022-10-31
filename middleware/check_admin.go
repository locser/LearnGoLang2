package middleware

import (
	"LearnGoLang2/model"
	"LearnGoLang2/model/req"
	"github.com/labstack/echo"
	"net/http"
)

func ISAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// handle logic
			req := req.ReqSignIn{}
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}

			if req.Email != "admin@gmail.com" {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "Đăng nhập bằng quyền Admin để sử dụng!",
					Data:       nil,
				})
			}

			return next(c)
		}
	}
}
