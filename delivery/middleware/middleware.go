package middleware

import (
	"bytes"
	"encoding/json"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("xid", uuid.New().String())

		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		var data interface{}
		_ = json.Unmarshal(bodyBytes, &data)

		model := model.Logger{
			Xid:    c.Request.Header.Get("xid"),
			Method: c.Request.Method,
			Url:    c.Request.RequestURI,
			Header: c.Request.Header,
			Body:   data,
		}

		c.Set("logger", model)
		logger.LogInfoWithRequest(model, "")
	}
}
