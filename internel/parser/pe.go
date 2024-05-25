package parser

import (
	"github.com/saferwall/pe"
)

func NewPE(filePath string) *pe.File {
	parser, err := pe.New(filePath, &pe.Options{})
	if err != nil {
		panic(err)
	}
	err = parser.Parse()
	if err != nil {
		panic(err)
	}
	return parser
}
