package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

// Returns uint16 value from env. If key is not set, it returns the fallback.
func GetUint16(key string, fallback uint16) uint16 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseUint(value, decimalBase, bitSize16)
	if err != nil {
		return fallback
	}

	return uint16(result)
}

// Returns uint16 value from env. If key is not set, it will panic.
func MustGetUint16(key string) uint16 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, key)
	}

	result, err := strconv.ParseUint(value, decimalBase, bitSize16)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return uint16(result)
}
