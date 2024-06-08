package parser

import (
	"github.com/radareorg/r2pipe-go"
	"strings"
)

type Disassembler struct {
	pipe      *r2pipe.Pipe
	sections  []string
	functions []string
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
		newDisassembler.sections = strings.Split(sections, "\n")
		functions, err := pipe.Cmd("afl")
		if err != nil {
			panic(err)
		}
		newDisassembler.functions = strings.Split(functions, "\n")
	}
	return newDisassembler
}

func (d *Disassembler) Functions() []string {
	return d.functions
}

func (d *Disassembler) Sections() []string {
	return d.sections
}
