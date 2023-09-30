package controller

import (
	"kloud/server/db"
	"kloud/server/model"

	"github.com/gin-gonic/gin"
)

// 请求信息，分别是文件对应的哈希列表以及差异块
type postPayload struct {
	Hashs []string   `json:"hashs"`
	Diffs []diffItem `json:"diffs"`
}

// 差异块
type diffItem struct {
	Hash    string `json:"hash"`    // blake3 + base64
	Content string `json:"content"` // zstd default + base64
}

func PostPostOnlineFile(c *gin.Context) {
	user := c.Param("user")
	path := c.Param("path")
	var payload postPayload
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		panic(err)
	}
	updateDBBlocks(payload.Diffs)
	updateDBFile(payload.Hashs, user, path)
	c.JSON(200, nil)
}

func updateDBBlocks(diffs []diffItem) {
	for _, diff := range diffs {
		db.Set("blocks", []byte(diff.Hash), model.NewBlock(diff.Hash, []byte(diff.Content)).Marshal())
	}
}

func updateDBFile(hashs []string, user, path string) {
	rawFile, err := db.Get("files", []byte(user+"-"+path))
	var file *model.File
	if err != nil {
		// if not found file
		file = createFile(user, path)
	} else {
		file = new(model.File).UnMarshal(rawFile)
	}
	file.Update(hashs)
	db.Set("files", []byte(user+"-"+path), file.Marshal())
}

func createFile(user, path string) *model.File {
	return model.NewFile(user, path)
}
