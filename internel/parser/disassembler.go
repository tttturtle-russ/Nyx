package parser

import (
	"github.com/radareorg/r2pipe-go"
)

type Disassembler struct {
	pipe      *r2pipe.Pipe
	sections  []*Section
	functions []*Function
	code      string
}

func NewDisassembler(file string, lazyLoad bool) *Disassembler {
	newDisassembler := &Disassembler{}
	pipe, err := r2pipe.NewPipe(file)
	if err != nil {
		panic(err)
	}
	if !lazyLoad {
		_, err = pipe.Cmd("aaaa")
		if err != nil {
			panic(err)
		}
		sections, err := pipe.Cmd("iS")
		if err != nil {
			panic(err)
		}
		newDisassembler.sections = parseSections(sections)
		functions, err := pipe.Cmd("afl")
		if err != nil {
			panic(err)
		}
		newDisassembler.functions = parseFunctions(functions)
		getFunctionCode(pipe, newDisassembler.functions)
	}
	return newDisassembler
}

func (d *Disassembler) Functions() []*Function {
	return d.functions
}

func (d *Disassembler) Sections() []*Section {
	return d.sections
}
