package memory

import (
	"encoding/binary"
	"fmt"
	"os"
)

const (
	memFileFormat = "/proc/%d/mem"
)

type Memory interface {
	Read(address, size uint64) ([]byte, error)
	Write64(address, value uint64) (int, error)
	Write32(address uint64, value uint32) (int, error)
	Write16(address uint64, value uint16) (int, error)
	Write8(address uint64, value uint8) (int, error)
	WriteString(address uint64, value string) (int, error)
}

type MemoryImpl struct {
	Pid int
}

func New(pid int) Memory {
	return &MemoryImpl{
		Pid: pid,
	}
}

func (m *MemoryImpl) Read(address, size uint64) ([]byte, error) {
	memFile := fmt.Sprintf(memFileFormat, m.Pid)
	fp, err := os.Open(memFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open mem file: %v", err)
	}
	defer fp.Close()

	data := make([]byte, size)

	n, err := fp.ReadAt(data, int64(address))
	if err != nil {
		return nil, fmt.Errorf("failed to read memory: %v", err)
	}
	if uint64(n) != size {
		return nil, fmt.Errorf("failed to read memory: expected %d bytes, got %d", size, n)
	}

	return data, nil
}

func (m *MemoryImpl) Write64(address, value uint64) (int, error) {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, value)
	return m.writeData(address, data)
}

func (m *MemoryImpl) Write32(address uint64, value uint32) (int, error) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, value)
	return m.writeData(address, data)
}

func (m *MemoryImpl) Write16(address uint64, value uint16) (int, error) {
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, value)
	return m.writeData(address, data)
}

func (m *MemoryImpl) Write8(address uint64, value uint8) (int, error) {
	data := make([]byte, 1)
	data[0] = value
	return m.writeData(address, data)
}

func (m *MemoryImpl) WriteString(address uint64, value string) (int, error) {
	return m.writeData(address, []byte(value))
}

// writeMemoryData writes binary data to the specified address of a process with the given PID.
func (m *MemoryImpl) writeData(address uint64, data []byte) (int, error) {
	memFile := fmt.Sprintf(memFileFormat, m.Pid)

	fp, err := os.OpenFile(memFile, os.O_WRONLY, 0)
	if err != nil {
		return 0, fmt.Errorf("failed to open mem file: %v", err)
	}
	defer fp.Close()

	n, err := fp.WriteAt(data, int64(address))
	if err != nil {
		return 0, fmt.Errorf("failed to write memory: %v", err)
	}
	if n != len(data) {
		return 0, fmt.Errorf("failed to write memory: expected %d bytes, wrote %d", len(data), n)
	}

	return n, nil
}
