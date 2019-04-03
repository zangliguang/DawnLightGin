package controller

import (
	"DawnLightGin/web/result"
	"github.com/gin-gonic/gin"

	"DawnLightGin/model/mosaicMovieDao"
	"DawnLightGin/model/noMosaicMovieDao"
	"fmt"
	"strconv"
)

var (
	listMovie result.GetDataFunc
	getMovie  result.GetDataFunc
	//addMovie    result.GetDataFunc
	//updateMovie result.GetDataFunc
	//deleteMovie result.GetDataFunc
)

func init() {
	InitMovieRouterHandler()
}
func InitMovieRouterHandler() {

	listMovie = func(c *gin.Context) (data interface{}, err error) {
		//start := c.DefaultQuery("start", "0")
		start, err := strconv.Atoi(c.DefaultQuery("start", "0"))
		pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
		mosaic, err := strconv.Atoi(c.DefaultQuery("mosaic", "1"))
		fmt.Println(fmt.Sprintf("请求数据:%d，页面大小：%d,马赛克：%d", start, pageSize, mosaic))
		if mosaic == 1 {
			return mosaicMovieDao.ListMovie(start, pageSize)
		} else {
			return noMosaicMovieDao.ListMovie(start, pageSize)
		}

	}

}
func SetMovieRouterGroup(router *gin.Engine) {
	MovieGroup := router.Group("/movie")
	MovieGroup.GET("/:id", getMovie.ToGinHandler()).
		GET("/", listMovie.ToGinHandler())
}
