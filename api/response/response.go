package response

import (
	"github.com/kongmsr/oneid-core/modelo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, modelo.Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	Result(modelo.SUCCESS, map[string]interface{}{}, "Success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(modelo.SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(modelo.SUCCESS, data, "Success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(modelo.SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(modelo.ERROR, map[string]interface{}{}, "Fail", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(modelo.ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(modelo.ERROR, data, message, c)
}

func FailWithError(data error) *modelo.Response {
	return &modelo.Response{
		Code: modelo.SUCCESS,
		Data: data.Error(),
		Msg:  "Failed",
	}
}
