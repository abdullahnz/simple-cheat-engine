package main

import (
	"flag"
	"fmt"

	"github.com/simple-cheat-engine/core"
)

type Args struct {
	Pid   int
	Maps  bool
	Read  uint64
	Size  uint64
	Write uint64
	Bit   uint64
	Value uint64
	Str   string
}

func setupArgs() Args {
	args := Args{}
	flag.IntVar(&args.Pid, "pid", -1, "PID to attach to (required)")
	flag.BoolVar(&args.Maps, "maps", false, "Print all memory regions")
	flag.Uint64Var(&args.Read, "read", 0, "Read memory at address")
	flag.Uint64Var(&args.Write, "write", 0, "Write memory at address")
	flag.Uint64Var(&args.Size, "size", 0, "Size of memory to read")
	flag.Uint64Var(&args.Bit, "bit", 0, "Bit size of memory to read")
	flag.Uint64Var(&args.Value, "val", 0, "Value to write to memory")
	flag.StringVar(&args.Str, "str", "", "String to write to memory")
	flag.Parse()

	return args
}

func main() {
	args := setupArgs()
	process := core.New(args.Pid)

	switch {
	case args.Maps:
		if err := process.Dump(); err != nil {
			fmt.Println(err)
		}

	case args.Read > 0:
		if err := process.Read(args.Read, args.Size, args.Bit); err != nil {
			fmt.Println(err)
		}

	case args.Write > 0:
		var (
			n   int
			err error
		)

		if args.Str != "" {
			n, err = process.WriteString(args.Write, args.Str)
		} else {
			n, err = process.WriteBytes(args.Write, args.Value, args.Bit)
		}

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Wrote %d bytes to 0x%x\n", n, args.Write)
		}

	default:
		flag.Usage()
	}
}
