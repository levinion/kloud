package model

import (
	"encoding/base64"
	"kloud/db"

	"github.com/bytedance/sonic"
	"github.com/valyala/gozstd"
)

/*
块是最基本的存储单位。每个块属于一个文件，拥有不固定的尺寸，表示文件中的一部分内容。
*/
type Block struct {
	Hash   string  `json:"hash"`
	Belong []*File `json:"belong"`
}

// 创建一个新的存储块；由于客户端已进行哈希计算，这里直接使用哈希
func NewBlock(hash string, content []byte) *Block {
	block := &Block{
		Hash: hash,
	}
	db.Set("content", []byte(hash), content)
	return block
}

func (b *Block) GetContent() []byte {
	// get zstd default + base64 string
	ori, err := db.Get("content", []byte(b.Hash))
	if err != nil {
		panic(err)
	}
	// decode base64
	base, err := base64.StdEncoding.DecodeString(string(ori))
	if err != nil {
		panic(err)
	}
	// decode zstd
	data, err := gozstd.Decompress(nil, base)
	if err != nil {
		panic(err)
	}
	return data
}

func (b *Block) Equal(other *Block) bool {
	return b.Hash == other.Hash
}

func (b *Block) UnMarshal(data []byte) *Block {
	err := sonic.Unmarshal(data, b)
	if err != nil {
		panic(err)
	}
	return b
}

func (b *Block) Marshal() []byte {
	data, err := sonic.Marshal(b)
	if err != nil {
		panic(err)
	}
	return data
}
