package main

import "github.com/radareorg/r2pipe-go"

func main() {
	r2p, err := r2pipe.NewPipe("/home/russ/test")
	if err != nil {
		print("ERROR: ", err)

		return
	}
	defer r2p.Close()
	r2p.Cmd("aaaa")
	r2p.Cmd("s main")
	disasm, err := r2p.Cmd("pdf")
	if err != nil {
		print("ERROR: ", err)
	} else {
		print(disasm, "\n")
	}
}
