package middleware

import (
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("xid", uuid.New().String())
		model := model.Logger{
			Xid:    c.Request.Header.Get("xid"),
			Method: c.Request.Method,
			Url:    c.Request.RequestURI,
			Header: c.Request.Header,
			Body:   c.Request.Body,
		}

		c.Set("logger", model)
		logger.LogInfoWithRequest(model, "")
	}
}
