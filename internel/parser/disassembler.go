package parser

import (
	"github.com/radareorg/r2pipe-go"
)

type Disassembler struct {
	pipe *r2pipe.Pipe
}

func NewDisassembler(file string) *Disassembler {
	pipe, err := r2pipe.NewPipe(file)
	if err != nil {
		panic(err)
	}
	_, err = pipe.Cmd("aaaa")
	if err != nil {
		panic(err)
	}
	return &Disassembler{
		pipe: pipe,
	}
}

func (d *Disassembler) Disassemble(segment string) {
	d.pipe.Cmd("s " + segment)
}

func (d *Disassembler) DisassembleText() {

}
