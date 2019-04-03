package result

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         // 状态码
	Message string      // 结果信息
	Data    interface{} // 返回内容实体
}

func (result *Result) ToGinH() (h gin.H) {
	h = gin.H{}
	h["code"] = result.Code
	h["message"] = result.Message
	h["data"] = result.Data
	return h
}

func BuildResult(data interface{}, err error) Result {
	if err != nil {
		return Result{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return Result{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    data,
	}

}
