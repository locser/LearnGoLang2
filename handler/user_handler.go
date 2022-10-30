package handler

import (
	"LearnGoLang2/log"
	"LearnGoLang2/model"
	req2 "LearnGoLang2/model/req"
	"LearnGoLang2/repository"
	security "LearnGoLang2/sercurity"
	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req2.ReqSignUp{}
	// khi nguời dùng đẩy dữ liu lên backend thì Context có trách nhiệm
	// laasy dữ liệu ra.
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validator2 := validator.New()
	if err := validator2.Struct(req); err != nil {
		log.Error(err.Error())

		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)

	if err != nil {
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	} else {
		//Hide password
		user.Password = ""
		return c.JSON(http.StatusOK, model.Response{
			StatusCode: http.StatusOK,
			Message:    "Thành công- Lưu User",
			Data:       user,
		})
	}
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Ryan",
		"email": "ryanLucifeed@gmail.com",
	})
}
