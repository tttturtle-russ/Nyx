package parser

import (
	"encoding/hex"
	"fmt"
	"github.com/saferwall/elf"
)

func NewElf(filePath string) *elf.Parser {
	parser, err := elf.New(filePath)
	if err != nil {
		panic(err)
	}
	err = parser.Parse()
	if err != nil {
		panic(err)
	}
	return parser
}

func GetTextSection(parser *elf.Parser) {
	for _, section := range parser.F.Sections64 {
		if section.SectionName == ".text" {
			data, err := section.Data()
			if err != nil {
				panic(err)
			}
			fmt.Println(hex.Dump(data))
		}
	}

}
