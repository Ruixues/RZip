package Modules

import (
	"RZip/abstract"
	"archive/zip"
	"fmt"
	"github.com/gen2brain/go-unarr"
)
type RarTar7Z struct {
	Archive *unarr.Archive
	NowNode  *abstract.FileTreeNode
	NowPath  []string
	AllNodes []*zip.File
	RootNode *abstract.FileTreeNode
	DirectVisit map [string]*abstract.FileTreeNode
}
func NewRarTar7Z (File string) RarTar7Z {
	ret := RarTar7Z{}
	ret.Init(File)
	return ret
}
func (z *RarTar7Z) Init (File string) error {
	var err error
	z.Archive,err = unarr.NewArchive(File)
	if err != nil {
		return err
	}
	//开始构建树
	z.RootNode = &abstract.FileTreeNode{
		FileInfo: nil,
		Sons:     make (map[string]*abstract.FileTreeNode),
		Message:  abstract.FileMessage{},
		Proxy:    "",
		Father:   nil,
	}
	contents,err := z.Archive.List()
	if err != nil {
		panic(err)
	}
	for _,file := range contents {
		fmt.Println (file)
	}
	return nil
}
func (z *RarTar7Z) GetNowPath () string {
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
func (z *RarTar7Z) GetNowFiles () []*abstract.FileTreeNode {
	return nil
}
func (z *RarTar7Z) GetIntoRelativeDirectory (Dictionary string) bool {
	return true
}