package model

import (
	"bufio"
	"os"
	"time"
)

type File struct {
	Name    string // 文件名/路径 指向文件系统内唯一的文件
	Hash    []byte // blake3
	Version int64  // 版本，时间戳表示
	blocks  *BlockList
}

// 对文件进行分块，返回区块列表
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

func (f *File) Open() (*os.File, error) {
	return os.Open(f.Name)
}
