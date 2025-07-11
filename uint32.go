package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

// Returns uint32 value from env. If key is not set, it returns the fallback.
func GetUint32(key string, fallback uint32) uint32 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseUint(value, decimalBase, bitSize32)
	if err != nil {
		return fallback
	}

	return uint32(result)
}

// Returns uint32 value from env. If key is not set, it will panic.
func MustGetUint32(key string) uint32 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, key)
	}

	result, err := strconv.ParseUint(value, decimalBase, bitSize32)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return uint32(result)
}
