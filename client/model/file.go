package model

import (
	"bufio"
	"kloud/client/user"
	"os"
	"path/filepath"
	"time"

	"github.com/imroc/req/v3"
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

const ServerAddr = "https://kloud.botland.top/"

// 从云端获取当前文件的块，仅获取块的哈希
func (f *File) FetchOnline() (*BlockList, error) {
	res, err := req.C().
		R().
		Get(filepath.Join(ServerAddr, "file", user.Me.ID, f.Name))
	if err != nil {
		return nil, err
	}
	blocks := NewBlockList().UnmarshalJson(res.Bytes())
	return blocks, nil
}

func (f *File) Sync() error {
	f.UpdateVersion()
	onlineBlocks, err := f.FetchOnline()
	if err != nil {
		return err
	}
	localBlocks, err := f.SplitBlocks()
	if err != nil {
		return err
	}
	diff := localBlocks.Diff(onlineBlocks)
	return f.PostOnline(diff)
}

// 上传差异块到云端
func (f *File) PostOnline(diff *BlockList) error {
	_, err := req.C().
		R().
		SetBody(diff).
		Post(filepath.Join(ServerAddr, "file", user.Me.ID, f.Name))
	if err != nil {
		return err
	}
	return nil
}

func (f *File) Open() (*os.File, error) {
	return os.Open(f.Name)
}
