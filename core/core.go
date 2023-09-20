package core

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/simple-cheat-engine/constants"
	"github.com/simple-cheat-engine/core/memory"
	"github.com/simple-cheat-engine/core/region"
)

type Core struct {
	Pid int
}

func New(pid int) *Core {
	return &Core{
		Pid: pid,
	}
}

func (c *Core) Dump() error {
	mapEntries, err := region.New(c.Pid).Read()
	if err != nil {
		return fmt.Errorf("core.Dump(): %v", err)
	}

	for i, entry := range mapEntries {
		fmt.Printf("[%02d] %s\n", i, entry.String())
	}

	return nil
}

func (c *Core) Read(address, size, bit uint64) error {
	if size <= 0 {
		size = constants.DEFAULT_READ_SIZE
	}

	data, err := memory.New(c.Pid).Read(address, size)

	if err != nil {
		return fmt.Errorf("core.Read(): %v", err)
	}

	buf := bytes.NewBuffer(data)

	for i := 0; i < buf.Len(); i += (int(bit) / 8) * constants.LINE_SIZE {

		out := fmt.Sprintf("%016x: ", address+uint64(i))
		for j := 0; j < constants.LINE_SIZE; j++ {
			val := uint64(0)
			if err := binary.Read(buf, binary.LittleEndian, &val); err != nil {
				break
			}
			out += fmt.Sprintf("%016x ", val)
		}
		fmt.Println(out)
	}

	return nil
}

func (c *Core) WriteBytes(address, value, bit uint64) (int, error) {
	var (
		n   int
		err error
	)

	if bit <= 0 {
		bit = constants.DEFAULT_WRITE_BIT
	}

	switch {
	case bit == 64:
		n, err = memory.New(c.Pid).Write64(address, value)
	case bit == 32:
		n, err = memory.New(c.Pid).Write32(address, uint32(value))
	case bit == 16:
		n, err = memory.New(c.Pid).Write16(address, uint16(value))
	case bit == 8:
		n, err = memory.New(c.Pid).Write8(address, uint8(value))
	default:
		return 0, fmt.Errorf("core.WriteBytes(): invalid bit size: %d", bit)
	}

	if err != nil {
		return 0, fmt.Errorf("core.WriteBytes(): %v", err)
	}

	return n, nil
}

func (c *Core) WriteString(address uint64, data string) (int, error) {
	n, err := memory.New(c.Pid).WriteString(address, data)
	if err != nil {
		return 0, fmt.Errorf("core.WriteString(): %v", err)
	}

	return n, nil

}
