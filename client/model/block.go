package model

import (
	"encoding/base64"

	"lukechampine.com/blake3"
)

/*
块是最基本的存储单位。每个块属于一个文件，拥有不固定的尺寸，表示文件中的一部分内容。
*/
type Block struct {
	hash    string
	belong  []*File
	content string
}

func NewBlock(data []byte) *Block {
	hash := blake3.Sum256(data)
	base := base64.StdEncoding.EncodeToString(hash[:])
	block := &Block{
		hash:    base,
		content: "",
	}
	block.content = string(data)
	return block
}

func (b *Block) GetContent() string {
	return b.content
}

func (b *Block) Equal(other *Block) bool {
	return b.hash == other.hash
}
