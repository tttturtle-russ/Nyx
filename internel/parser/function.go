package parser

import (
	"github.com/radareorg/r2pipe-go"
	"strconv"
	"strings"
)

type Function struct {
	StartOffset int32
	Name        string
	Size        int32
	Code        string
}

// parseFunctions parse the output of r2 command afl
func parseFunctions(pipeOut string) []*Function {
	var result []*Function

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
		result = append(result, &Function{
			StartOffset: int32(addr),
			Name:        fields[3],
			Size:        int32(size),
		})
	}
	return result
}

// getFunctionCode gets all the functions assemble code
func getFunctionCode(pipe *r2pipe.Pipe, functions []*Function) {
	for _, function := range functions {
		_, err := pipe.Cmdf("s %s", function.Name)
		if err != nil {
			panic(err)
		}
		code, err := pipe.Cmd("pdf")
		if err != nil {
			panic(err)
		}
		function.Code = code
	}
}
