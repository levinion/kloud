package model

type BlockList struct {
	head   *BlockNode
	length int
}

type BlockNode struct {
	this *Block
	next *BlockNode
}

func (b *BlockNode) Next() *BlockNode {
	b = b.next
	return b
}

func (b *BlockNode) Equal(other *BlockNode) bool {
	return b.this.Equal(other.this)
}

func NewBlockList() *BlockList {
	return &BlockList{
		head: &BlockNode{nil, nil},
	}
}

func (b *BlockList) Append(block *Block) *BlockList {
	p := b.head
	for p.next != nil {
		p = p.next
	}
	p.next = &BlockNode{block, nil}
	b.length++
	return b
}

func (b *BlockList) Iter() *BlockNode {
	return b.head
}

func (b *BlockList) Compare(other *BlockList) {
	thisIter, otherIter := b.Iter(), other.Iter()
	for otherIter.next != nil {
		if thisIter.next == nil || !thisIter.next.Equal(otherIter.next) {
			thisIter.next = &BlockNode{otherIter.next.this, thisIter.next}
		}
		thisIter.Next()
		otherIter.Next()
	}
}
