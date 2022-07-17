package jwt

import (
	"invitations-mechanism/config"
	"invitations-mechanism/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(username, email string) (result model.ResultClaims, err error) {
	claims := model.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.GetServiceName() + config.GetServiceVersion(),
			ExpiresAt: time.Now().Add(time.Duration(config.GetJWTExpired() * int(time.Minute))).Unix(),
		},
		Username: username,
		Email:    email,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString(config.GetJWTSecret())
	if err != nil {
		return
	}

	result = model.ResultClaims{
		Token:     signedToken,
		ExpiredAt: time.Unix(claims.ExpiresAt, 0),
	}

	return
}
