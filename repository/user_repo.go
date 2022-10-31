package repository

import (
	"LearnGoLang2/model"
	req2 "LearnGoLang2/model/req"
	"context"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req2.ReqSignIn) (model.User, error)
	SelectUserById(context context.Context, userId string) (model.User, error)
	UpdateUser(context context.Context, user model.User) (model.User, error)
}
