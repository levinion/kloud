package controller

import (
	"kloud/server/db"
	"kloud/server/model"

	"github.com/gin-gonic/gin"
)

func GetFetchOnlineFile(c *gin.Context) {
	user := c.Param("user")
	path := c.Param("path")
	file, _ := getDBFile(user, path)
	c.JSON(200, file.Hashs)
}

func getDBFile(user, path string) (*model.File, bool) {
	rawFile, err := db.Get("files", []byte(user+"-"+path))
	var file *model.File
	if err != nil {
		// if not found file
		file = createFile(user, path)
		return file, false
	} else {
		file = new(model.File).UnMarshal(rawFile)
		return file, true
	}
}
