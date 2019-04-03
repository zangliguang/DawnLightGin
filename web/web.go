package web

import (
	"DawnLightGin/web/controller"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}
func Run() {
	controller.SetActressRouterGroup(router)
	controller.SetMovieRouterGroup(router)
	router.Run(":8000")
}
