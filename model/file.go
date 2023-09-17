package model

import (
	"bufio"
	"os"
	"time"
)

type File struct {
	Name    string // 文件名/路径 指向文件系统内唯一的文件
	Weak    []byte // 弱哈希，md5，碰撞概率小
	Version int64  // 版本，备份用，可用时间戳
	blocks  *BlockList
}

// 对文件进行分块，这里返回结果应当是一个链表
func (f *File) SplitBlocks() (*BlockList, error) {
	file, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	blocks := NewBlockList()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Bytes()
		blocks.Append(NewBlock(data))
	}
	return blocks, nil
}

func (f *File) UpdateVersion() {
	f.Version = time.Now().Unix()
}

func (f *File) Sync() error {
	f.UpdateVersion()
	blocks, err := f.SplitBlocks()
	if err != nil {
		return err
	}
	
}

func (f *File) Open() (*os.File, error) {
	return os.Open(f.Name)
}

func (f *File) Create() (*os.File, error) {
	return os.Create(f.Name)
}
