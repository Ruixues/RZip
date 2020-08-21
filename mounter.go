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
		return Modules.NewZip(FileName)
	}
	return nil
}
