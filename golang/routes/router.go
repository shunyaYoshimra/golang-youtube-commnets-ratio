package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimura/YouTube_Comments/controllers"
)

func NewRouter(r *gin.Engine) {
	videoController := controllers.NewVideoController()
	r.POST("/video", videoController.Create)
	r.GET("/videos", videoController.Index)
}
