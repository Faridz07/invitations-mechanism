package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"invitations-mechanism/config"
	"invitations-mechanism/delivery/helper"
	"invitations-mechanism/infrastructure/constant"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set(constant.XID, uuid.New().String())

		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		var data interface{}
		_ = json.Unmarshal(bodyBytes, &data)

		model := model.Logger{
			Xid:    c.Request.Header.Get(constant.XID),
			Method: c.Request.Method,
			Url:    c.Request.RequestURI,
			Header: c.Request.Header,
			Body:   data,
		}

		c.Set("logger", model)
		logger.LogInfoWithRequest(model, "")
	}
}

func AdminValidations() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get(constant.AUTHORIZATION)
		if !strings.Contains(strings.ToLower(authorizationHeader), strings.ToLower(constant.BEARER)) {
			helper.ResponseErrorWithCode(c, http.StatusUnauthorized, constant.UNAUTHORIZED, nil)
			c.Abort()
			return
		}

		tokenString := strings.Replace(authorizationHeader, constant.BEARER, "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New(constant.ErrInvalidSigningMethod)
			} else if method != jwt.SigningMethodHS256 {
				return nil, errors.New(constant.ErrInvalidSigningMethod)
			}

			return config.GetJWTSecret(), nil
		})

		if err != nil {
			helper.ResponseErrorWithCode(c, http.StatusUnauthorized, constant.UNAUTHORIZED, err)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			helper.ResponseErrorWithCode(c, http.StatusUnauthorized, constant.UNAUTHORIZED, errors.New(constant.ErrInvalidToken))
			c.Abort()
			return
		}

		if claims[constant.ROLE] == "" || claims[constant.ROLE] != constant.ADMIN {
			helper.ResponseErrorWithCode(c, http.StatusUnauthorized, constant.UNAUTHORIZED, errors.New(constant.ErrRoleName))
			c.Abort()
			return
		}

		c.Set(constant.JWTClaims, claims)
	}
}
