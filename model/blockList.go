package model

type BlockList struct {
	head *BlockNode
}

type BlockNode struct {
	this *Block
	next *Block
}
