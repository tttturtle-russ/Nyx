package parser

import (
	"github.com/radareorg/r2pipe-go"
	"strconv"
	"strings"
)

type function struct {
	addr int32
	name string
	size int32
}

func parseFunctions(pipeOut string) []function {
	var result []function

	outputs := strings.Split(pipeOut, "\n")
	for _, output := range outputs {
		fields := strings.Fields(output)
		addr, err := strconv.ParseInt(fields[0], 0, 32)
		if err != nil {
			panic(err)
		}
		size, err := strconv.ParseInt(fields[2], 0, 32)
		if err != nil {
			panic(err)
		}
		result = append(result, function{
			addr: int32(addr),
			name: fields[3],
			size: int32(size),
		})
	}
	return result
}

type Disassembler struct {
	pipe      *r2pipe.Pipe
	sections  []string
	functions []function
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
		newDisassembler.functions = parseFunctions(functions)
	}
	return newDisassembler
}

func (d *Disassembler) Functions() []function {
	return d.functions
}

func (d *Disassembler) Sections() []string {
	return d.sections
}
