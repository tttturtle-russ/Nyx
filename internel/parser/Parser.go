package parser

import "github.com/bnagy/gapstone"

type Parser interface {
	Parse() error
}

type Disassembler struct {
	engine *gapstone.Engine
	parser Parser
}

func (d *Disassembler) Disassemble() {
	//d.engine.Disasm()
}
