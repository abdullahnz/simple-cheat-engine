package region

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/simple-cheat-engine/constants"
	"github.com/simple-cheat-engine/helper"
)

type Maps interface {
	Read() ([]*MapEntry, error)
}

type MapsImpl struct {
	Pid int
}

const (
	mapsFileFormat = "/proc/%d/maps"
)

func New(pid int) Maps {
	return &MapsImpl{
		Pid: pid,
	}
}

func (m *MapsImpl) Read() ([]*MapEntry, error) {
	mapsFile := fmt.Sprintf(mapsFileFormat, m.Pid)

	regions, err := ioutil.ReadFile(mapsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read maps file: %v", err)
	}

	var entries []*MapEntry

	lines := strings.Split(string(regions), constants.NEWLINE)

	for _, line := range lines {
		if !helper.IsEmptyString(line) {
			entry, err := parseMapEntry(line)
			if err != nil {
				return nil, fmt.Errorf("failed to parse map entry: %v", err)
			}
			entries = append(entries, entry)
		}
	}

	return entries, nil
}
