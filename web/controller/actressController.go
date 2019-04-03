package controller

import (
	"github.com/gin-gonic/gin"
	"DawnLightGin/web/result"

	"gin_gorm_restful/util"
	"DawnLightGin/model/actressDao"
	"strconv"
	"fmt"
)

var (
	getActress    result.GetDataFunc
	listActress   result.GetDataFunc
	addActress    result.GetDataFunc
	updateActress result.GetDataFunc
	deleteActress result.GetDataFunc
)

func init() {
	InitActressRouterHandler()
}
func InitActressRouterHandler() {
	getActress = func(c *gin.Context) (data interface{}, err error) {
		// 从URL中获取用户ID
		id, err := util.Str2Uint(c.Param("id"))
		if err != nil {
			return nil, err
		}

		actress, err := actressDao.GetActress(id)
		return actress, err
	}

	listActress = func(c *gin.Context) (data interface{}, err error) {
		//start := c.DefaultQuery("start", "0")
		start, err := strconv.Atoi(c.DefaultQuery("start", "0"))
		pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "100"))
		fmt.Println(fmt.Sprintf("请求数据:%d，页面大小：%d", start, pageSize))
		actresses, err := actressDao.ListActress(start, pageSize)
		return actresses, err
	}

	addActress = func(c *gin.Context) (data interface{}, err error) {
		actress, err := GetActressByContext(c)

		if err != nil {
			return nil, err
		}

		err = actressDao.AddActress(&actress)
		return actress, err
	}

	updateActress = func(c *gin.Context) (data interface{}, err error) {
		actress, err := GetActressByContext(c)

		if err != nil {
			return nil, err
		}

		// 从URL中获取用户ID
		id, err := util.Str2Uint(c.Param("id"))
		if err != nil {
			return nil, err
		}

		actress.ID = id

		err = actressDao.UpdateActress(&actress)
		return actress, err
	}

	deleteActress = func(c *gin.Context) (data interface{}, err error) {
		// 从URL中获取用户ID
		id, err := util.Str2Uint(c.Param("id"))
		if err != nil {
			return nil, err
		}

		err = actressDao.DeleteActressByID(id)
		return nil, err
	}
}
func SetActressRouterGroup(router *gin.Engine) {
	actressGroup := router.Group("/actress")
	actressGroup.GET("/:id", getActress.ToGinHandler()).
		GET("/", listActress.ToGinHandler()).
		POST("/", addActress.ToGinHandler()).
		PUT("/:id", updateActress.ToGinHandler()).
		DELETE("/:id", deleteActress.ToGinHandler())
	//movieGroup := router.Group("/")
	//movieGroup.GET("/:id", getActress.ToGinHandler()).
	//	GET("/", listActress.ToGinHandler()).
	//	POST("/", addActress.ToGinHandler()).
	//	PUT("/:id", updateActress.ToGinHandler()).
	//	DELETE("/:id", deleteActress.ToGinHandler())
}
func GetActressByContext(c *gin.Context) (actress actressDao.Actress, err error) {
	actress = actressDao.Actress{}

	// 将JSON转为User对象
	err = c.ShouldBindJSON(&actress)

	return actress, err
}
