package parser

import (
	"strconv"
	"strings"
)

type Section struct {
	Paddr      int32
	Vaddr      int32
	Size       int32
	Vsize      int32
	Permission string
	Name       string
}

func parseSections(pipeOut string) []*Section {
	var result []*Section

	// remove the head 4 line and the NULL Section
	outputs := strings.Split(pipeOut, "\n")[5:]

	for _, output := range outputs {
		fields := strings.Fields(output)
		paddr, err := strconv.ParseInt(fields[1], 0, 32)
		if err != nil {
			panic(err)
		}
		size, err := strconv.ParseInt(fields[2], 0, 32)
		if err != nil {
			panic(err)
		}
		vaddr, err := strconv.ParseInt(fields[3], 0, 32)
		if err != nil {
			panic(err)
		}
		vsize, err := strconv.ParseInt(fields[4], 0, 32)
		if err != nil {
			panic(err)
		}
		result = append(result, &Section{
			Name:       fields[7],
			Paddr:      int32(paddr),
			Vaddr:      int32(vaddr),
			Size:       int32(size),
			Vsize:      int32(vsize),
			Permission: fields[5],
		})
	}
	return result
}
