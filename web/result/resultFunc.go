package result

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

type GetDataFunc func(c *gin.Context) (data interface{}, err error)

func (getData GetDataFunc) ToGinHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		data, err := getData(context)
		resultEntity := BuildResult(data, err)
		context.JSON(http.StatusOK, resultEntity.ToGinH())
	}
}
