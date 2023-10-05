package model

import (
	"kloud/db"
	"time"

	"github.com/bytedance/sonic"
)

type File struct {
	User    string   `json:"user"`
	Path    string   `json:"path"`    // 文件的云端路径
	Version int64    `json:"version"` // 版本，时间戳表示
	Hashs   []string `json:"hashs"`   // 块哈希列表
}

func NewFile(user, path string) *File {
	file := &File{
		User:  user,
		Path:  path,
		Hashs: make([]string, 0),
	}
	file.updateVersion()
	return file
}

func (f *File) updateVersion() {
	f.Version = time.Now().Unix()
}

func (f *File) Update(hashs []string) {
	f.Hashs = hashs // 更新哈希列表
	f.updateVersion()
}

func (f *File) UnMarshal(data []byte) *File {
	err := sonic.Unmarshal(data, f)
	if err != nil {
		panic(err)
	}
	return f
}

func (f *File) Marshal() []byte {
	data, err := sonic.Marshal(f)
	if err != nil {
		panic(err)
	}
	return data
}

func (f *File) GetBlocks() *BlockList {
	blocks := NewBlockList()
	for _, hash := range f.Hashs {
		data, err := db.Get("blocks", []byte(hash))
		if err != nil {
			panic(err)
		}
		block := new(Block).UnMarshal(data)
		blocks.Append(block)
	}
	return blocks
}

func (f *File) GetContent() []byte {
	blocks := f.GetBlocks()
	return blocks.GetContent()
}
