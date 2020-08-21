package Modules

import (
	"RZip/abstract"
	"archive/zip"
	"fmt"
	"strings"
)

type DealerZip struct {
	Reader   *zip.ReadCloser
	NowNode  abstract.FileTreeNode
	NowPath  []string
	AllNodes []*zip.File
	RootNode abstract.FileTreeNode
}

func NewZip(File string) (ret DealerZip) {
	ret.Init(File)
	return
}
func (z DealerZip) Init(name string) error {
	var err error
	z.Reader, err = zip.OpenReader(name)
	if err != nil {
		return err
	}
	//开始初始化树结构
	z.AllNodes = z.Reader.File
	z.NowPath = []string{""}
	z.RootNode = abstract.FileTreeNode{
		FileInfo: nil,
		Sons:     make(map[string]*abstract.FileTreeNode),
	}
	z.NowNode = z.RootNode
	for _, v := range z.AllNodes {
		z.InsertNode(&z.RootNode, v.Name, v)
	}
	return nil
}
func (z DealerZip) GetNowPath() string {
	var ret string
	len := len(z.NowPath) - 1
	for i, v := range z.NowPath {
		ret += v
		if i != len {
			ret += "/"
		}
	}
	return ret
}
func (z DealerZip) InsertNode(NowNode *abstract.FileTreeNode, Name string, FileInfo *zip.File) bool {
	list := strings.Split(Name,"/")
	fmt.Println (list)
	for _, v := range list {
		if _, ok := NowNode.Sons[v]; !ok {
			var sons map[string]*abstract.FileTreeNode
			if !FileInfo.FileInfo().IsDir() { //是文件
				sons = nil
			} else {
				sons = make(map[string]*abstract.FileTreeNode)
			}
			NowNode.Sons[v] = &abstract.FileTreeNode{
				FileInfo: FileInfo,
				Sons:     sons,
			}
		}
	}
	return true
}
func (z DealerZip) GetNowFiles() []*abstract.FileTreeNode {
	ret := make([]*abstract.FileTreeNode, 0)
	for _, v := range z.NowNode.Sons {
		ret = append(ret, v)
	}
	return ret
}

func (z DealerZip) GetIntoRelativeDirectory(name string) bool {
	if _, ok := z.NowNode.Sons[name]; !ok {
		return false
	}
	v := z.NowNode.Sons[name]
	if !v.FileInfo.(*zip.File).FileInfo().IsDir() {
		return false
	}
	return true
}
