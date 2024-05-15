package parser

import (
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
