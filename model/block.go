package model

/*
块是最基本的存储单位。每个块属于一个文件，拥有不固定的尺寸，表示文件中的一部分内容。
*/
type Block struct {
	Weak   string // 弱哈希

	// 以下指向块具体内容
	file   *File  // 块表示的文件入口
	offset int64  // 偏移
	size   int64  // 尺寸
}

func (b *Block) getContent() string{
	
}

func (b *Block) GetStrongHash() string {

}

func (b *Block) Equal(other *Block) bool {
	return b.GetStrongHash() == other.GetStrongHash()
}
