package helper

import (
	"strconv"

	"github.com/simple-cheat-engine/constants"
)

func ParseHex(s string) uint64 {
	num, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		num = 0
	}
	return num
}

func ParseUintptr(s string) uintptr {
	num, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		num = 0
	}
	return uintptr(num)
}

func IsEmptyString(s string) bool {
	return s == constants.EMPTY_STRING
}
