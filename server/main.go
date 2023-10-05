package main

import (
	"kloud/cache"
	"kloud/db"
	"kloud/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()
	cache.UseNoCache()
	db.UseBadgerDB()
	r := gin.Default()
	Route(r)
	r.Run(":8080")
}
