package main

import (
	"kloud/server/cache"
	"kloud/server/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cache.UseNoCache()
	db.UseBadgerDB()
	r := gin.Default()
	Route(r)
}
