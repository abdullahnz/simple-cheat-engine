package region

import (
	"fmt"
	"strings"

	"github.com/simple-cheat-engine/helper"
)

// MapEntry represents a memory map entry.
type MapEntry struct {
	StartAddress uintptr
	EndAddress   uintptr
	Permissions  string
	Size         uintptr
	File         string
}

func (e *MapEntry) String() string {
	return fmt.Sprintf("%x-%x %s %06x %s", e.StartAddress, e.EndAddress, e.Permissions, e.Size, e.File)
}

func parseMapEntry(line string) (*MapEntry, error) {
	fields := strings.Fields(line)

	addrRange := strings.Split(fields[0], "-")
	if len(addrRange) != 2 {
		return nil, fmt.Errorf("invalid address range: %s", line)
	}

	startAddress := helper.ParseUintptr(addrRange[0])
	endAddress := helper.ParseUintptr(addrRange[1])
	permissions := fields[1]
	file := fields[len(fields)-1]

	return &MapEntry{
		StartAddress: startAddress,
		EndAddress:   endAddress,
		Permissions:  permissions,
		Size:         endAddress - startAddress,
		File:         file,
	}, nil
}
