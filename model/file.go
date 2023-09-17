package model

type File struct {
	Name    string // 文件名/路径？指向文件系统内唯一的文件
	Weak    string // 弱哈希，md5，碰撞概率极小
	Version string // 版本，备份用，可用时间戳
}

// 对文件进行分块，这里返回结果应当是一个链表
func (f *File) SplitBlocks() []*Block{

}
