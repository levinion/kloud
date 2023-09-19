package model

type BlockList struct {
	list []*Block
}

func NewBlockList() *BlockList {
	return &BlockList{
		make([]*Block, 0),
	}
}

func (b *BlockList) Append(block *Block) *BlockList {
	b.list = append(b.list, block)
	return b
}

func (b *BlockList) IntoHashMap() map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range b.list {
		m[string(v.hash)] = struct{}{}
	}
	return m
}

func (b *BlockList) Diff(other *BlockList) *BlockList {
	m := other.IntoHashMap()
	diff := NewBlockList()
	// 查找不存在于云端的块，此即为差异块
	for _, v := range b.list {
		if _, ok := m[string(v.hash)]; !ok {
			diff.list = append(diff.list, v)
		}
	}
	return diff
}

func (b *BlockList) IntoHashSlice() [][]byte {
	r := make([][]byte, 0)
	for _, v := range b.list {
		r = append(r, v.hash)
	}
	return r
}
