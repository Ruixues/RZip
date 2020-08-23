package main

import (
	"RZip/Modules"
	"RZip/abstract"
	"fmt"
	"github.com/h2non/filetype"
	"os"
)

func GetDealer(FileName string) abstract.Dealer {
	file, _ := os.Open(FileName)
	head := make([]byte, 261)
	file.Read(head)
	if !filetype.IsArchive(head) {
		return nil
	}
	t, err := filetype.Match(head)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	switch t.Extension {
	case "zip":
		ret := Modules.NewZip(FileName)
		return &ret
	case "rar":
		ret := Modules.NewRarTar7Z(FileName)
		return &ret
	case "7z":
		ret := Modules.NewRarTar7Z(FileName)
		return &ret
	case "tar":
		ret := Modules.NewRarTar7Z(FileName)
		return &ret
	default:
		return nil
	}
}
