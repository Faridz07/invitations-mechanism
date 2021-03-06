package helper

import (
	"invitations-mechanism/infrastructure/constant"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOK(c *gin.Context, msg string) {
	body := model.Response{
		Status:  constant.OK,
		Message: msg,
	}

	c.Set("response", body)
	c.JSON(http.StatusOK, body)
	logger.LogInfoWithResponse(c, msg)
	return
}

func ResponseOKWithSingleData(c *gin.Context, msg string, data interface{}) {
	body := model.ResponseWithSingleData{
		Status:  constant.OK,
		Message: msg,
		Data:    data,
	}

	c.Set("response", body)
	c.JSON(http.StatusOK, body)
	logger.LogInfoWithResponse(c, msg)
	return
}

func ResponseOKWithArrayData(c *gin.Context, msg string, data []interface{}) {
	body := model.ResponseWithArrayData{
		Status:  constant.OK,
		Message: msg,
		Data:    data,
	}

	c.Set("response", body)
	c.JSON(http.StatusOK, body)
	logger.LogInfoWithResponse(c, msg)
	return
}

func ResponseFailed(c *gin.Context, code int, msg string, err error) {
	body := model.Response{
		Status:  constant.FAILED,
		Message: msg,
	}

	c.Set("response", body)
	c.JSON(code, body)

	if err != nil {
		logger.LogErrorWithRequest(c, err.Error())
		return
	}

	logger.LogErrorWithRequest(c, msg)
	return
}

func ResponseError(c *gin.Context, msg string, err error) {
	body := model.Response{
		Status:  constant.ERROR,
		Message: msg,
	}

	c.Set("response", body)
	c.JSON(http.StatusInternalServerError, body)

	if err != nil {
		logger.LogErrorWithRequest(c, err.Error())
		return
	}

	logger.LogErrorWithRequest(c, msg)
	return
}

func ResponseErrorWithCode(c *gin.Context, code int, msg string, err error) {
	body := model.Response{
		Status:  constant.ERROR,
		Message: msg,
	}

	c.Set("response", body)
	c.JSON(code, body)

	if err != nil {
		logger.LogErrorWithRequest(c, err.Error())
		return
	}

	logger.LogErrorWithRequest(c, msg)
	return
}
