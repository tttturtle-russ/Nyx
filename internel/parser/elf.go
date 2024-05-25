package parser

import (
	"fmt"
	"github.com/bnagy/gapstone"
	"github.com/saferwall/elf"
)

type Elf struct {
	elf   *elf.Parser
	arch  int
	mode  int
	_type string
}

func NewElf(filePath string) *Elf {
	parser, err := elf.New(filePath)
	if err != nil {
		panic(err)
	}
	return &Elf{
		elf: parser,
	}
}

func (e *Elf) Parse() {
	err := e.elf.Parse()
	if err != nil {
		panic(err)
	}
	switch e.elf.F.Machine {
	case elf.EM_X86_64:
		switch e.elf.F.Class() {
		case elf.ELFCLASS64:
			e.mode = gapstone.CS_MODE_64
		case elf.ELFCLASS32:
			e.mode = gapstone.CS_MODE_64
		}
		e.arch = gapstone.CS_ARCH_X86
	default:
		panic("Only support 32/64 elf file")
	}
	e._type = e.elf.F.Type.String()
}

func (e *Elf) StartAddress() uint64 {
	return e.elf.F.Entry
}

func (e *Elf) Code() []byte {
	for _, symbol := range e.elf.F.NamedSymbols {
		fmt.Println(symbol)
	}
	return nil
}
