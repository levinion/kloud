package main

import (
	"kloud/server/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	r.GET("/file/:user/*path", controller.GetFetchFile)
	r.POST("/file/:user/*path", controller.PostPostFile)
	return r
}
