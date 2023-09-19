package model

import (
	"bytes"
	"kloud/server/db"

	"lukechampine.com/blake3"
)

/*
块是最基本的存储单位。每个块属于一个文件，拥有不固定的尺寸，表示文件中的一部分内容。
*/
type Block struct {
	hash    []byte
	belong  []*File
	content []byte
}

func NewBlock(data []byte) *Block {
	hash := blake3.Sum256(data)
	block := &Block{
		hash:    hash[:],
		content: nil,
	}
	db.Set("content", block.hash, data)
	return block
}

func (b *Block) GetContent() []byte {
	data, err := db.Get("content", b.hash)
	if err != nil {
		return nil
	}
	return data
}

func (b *Block) Equal(other *Block) bool {
	return bytes.Equal(b.hash, other.hash)
}
