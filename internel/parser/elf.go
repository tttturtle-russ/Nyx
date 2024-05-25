package parser

import (
	"fmt"
	"github.com/saferwall/elf"
	"github.com/tidwall/gjson"
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
	jsonString, err := parser.DumpJSON()
	for _, symbol := range parser.F.NamedSymbols {
		fmt.Println(symbol.Name)
	}
	if err != nil {
		panic(err)
	}
	gjson.Get(jsonString, "")
	return nil
}
