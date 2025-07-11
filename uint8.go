package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

// Returns uint8 value from env. If key is not set, it returns the fallback.
func GetUint8(key string, fallback uint8) uint8 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseUint(value, decimalBase, bitSize8)
	if err != nil {
		return fallback
	}

	return uint8(result)
}

// Returns uint8 value from env. If key is not set, it will panic.
func MustGetUint8(key string) uint8 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, key)
	}

	result, err := strconv.ParseUint(value, decimalBase, bitSize8)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return uint8(result)
}
