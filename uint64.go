package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

// Returns uint64 value from env. If key is not set, it returns the fallback.
func GetUint64(key string, fallback uint64) uint64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseUint(value, decimalBase, bitSize64)
	if err != nil {
		return fallback
	}

	return uint64(result)
}

// Returns uint64 value from env. If key is not set, it will panic.
func MustGetUint64(key string) uint64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, key)
	}

	result, err := strconv.ParseUint(value, decimalBase, bitSize64)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return uint64(result)
}
