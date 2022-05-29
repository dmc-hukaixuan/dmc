package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// starand response data structure
type Response struct {
	Code   int         `json:"code"`
	Result int         `json:"result"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

// const
const (
	ERROR   = -1
	SUCCESS = 1
)

// pubile return result with data
func Result(result int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Result: result,
		Data:   data,
		Msg:    msg,
	})
}

// a unified prompt that the user has performed a centain operation on the page successfully
// 用户在页面做了某一个操作成功的统一提示
func Success(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "operation successful", c)
}

// Returns with message as parameter
func SuccessWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

// Returns with data as parameter
func SuccessWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "", c)
}

// Returns with data and message as parameter
func SuccessWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// fail return
func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

// fail return with some message
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

// fail return with data and message
func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
