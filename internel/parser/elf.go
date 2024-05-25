package parser

import (
	"fmt"
	"github.com/saferwall/elf"
)

type Elf struct {
	elf  *elf.Parser
	os   string
	arch int
	mode int
}

func NewElf(filePath string) *Elf {
	parser, err := elf.New(filePath)
	if err != nil {
		panic(err)
	}
	err = parser.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(parser.F.Machine.String())
	fmt.Println(parser.DumpJSON())
	return nil
}
