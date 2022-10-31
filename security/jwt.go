package security

import (
	"LearnGoLang2/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SECRET_KEY = "learngolanglalalafdfds"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		fmt.Println("Loi tao token", err.Error())
		return "Tao Token Loi!", err
	}

	return result, nil

}
