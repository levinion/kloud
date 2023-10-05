package controller

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetDownloadFile(c *gin.Context) {
	user := c.Param("user")
	path := c.Param("path")
	filename := filepath.Base(path)
	if file, ok := getDBFile(user, path); ok {
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Data(200, "application/octet-stream", file.GetContent())
		return
	}
	c.JSON(400, gin.H{
		"code": 400,
		"msg":  "file not found",
	})

}
