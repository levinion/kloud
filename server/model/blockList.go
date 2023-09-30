package model

import (
	"bytes"
)

type BlockList struct {
	List []*Block `json:"list"`
}

func NewBlockList() *BlockList {
	return &BlockList{
		make([]*Block, 0),
	}
}

func (b *BlockList) Append(block *Block) *BlockList {
	b.List = append(b.List, block)
	return b
}

func (b *BlockList) GetContent() []byte {
	buf := bytes.NewBuffer([]byte{})
	for _, block := range b.List {
		content := block.GetContent()
		buf.Write(content)
	}
	return buf.Bytes()
}
