package model

import (
	"bytes"
	"crypto/sha256"
	"hash/adler32"
	"kloud/db"
)

/*
块是最基本的存储单位。每个块属于一个文件，拥有不固定的尺寸，表示文件中的一部分内容。
*/
type Block struct {
	weak   []byte // 弱哈希
	strong []byte // 强哈希
	belong []*File
}

func NewBlock(data []byte) *Block {
	block := &Block{}
	block.weak = adler32.New().Sum(data)
	block.strong = block.GetStrongHash()
	db.Set("content", block.strong, data)

	return block
}

func (b *Block) GetContent() []byte {
	data, err := db.Get("content", b.strong)
	if err != nil {
		return nil
	}
	return data
}

func (b *Block) GetWeakHash() []byte {
	return b.weak
}

func (b *Block) GetStrongHash() []byte {
	hash := sha256.Sum256(b.GetContent())
	return hash[:]
}

func (b *Block) Equal(other *Block) bool {
	return bytes.Equal(b.GetStrongHash(), other.GetStrongHash())
}
