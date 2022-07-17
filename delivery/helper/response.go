package helper

import (
	"invitations-mechanism/constant"
	"invitations-mechanism/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, model.Response{
		Status:  constant.OK,
		Message: msg,
	})
}

func ResponseOKWithSingleData(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, model.ResponseWithSingleData{
		Status:  constant.OK,
		Message: msg,
		Data:    data,
	})
}

func ResponseOKWithArrayData(c *gin.Context, msg string, data []interface{}) {
	c.JSON(http.StatusOK, model.ResponseWithArrayData{
		Status:  constant.OK,
		Message: msg,
		Data:    data,
	})
}

func ResponseFailed(c *gin.Context, code int, msg string) {
	c.JSON(code, model.Response{
		Status:  constant.FAILED,
		Message: msg,
	})
}

func ResponseError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, model.Response{
		Status:  constant.ERROR,
		Message: msg,
	})
}

func ResponseErrorWithCode(c *gin.Context, code int, msg string) {
	c.JSON(code, model.Response{
		Status:  constant.ERROR,
		Message: msg,
	})
}
