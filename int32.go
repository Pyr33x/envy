package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

// Returns int32 value from env. If key is not set, it returns the fallback.
func GetInt32(key string, fallback int32) int32 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseInt(value, decimalBase, bitSize32)
	if err != nil {
		return fallback
	}

	return int32(result)
}

// Returns int32 value from env. If key is not set, it will panic.
func MustGetInt32(key string) int32 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, key)
	}

	result, err := strconv.ParseInt(value, decimalBase, bitSize32)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return int32(result)
}
