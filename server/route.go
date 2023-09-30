package main

import (
	"kloud/server/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	r.GET("/file/:user/*path", controller.GetFetchOnlineFile)
	r.POST("/file/:user/*path", controller.PostPostOnlineFile)
	r.GET("/download/:user/*path", controller.GetDownloadFile)
	return r
}
