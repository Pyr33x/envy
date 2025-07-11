package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

// Returns int64 value from env. If key is not set, it returns the fallback.
func GetInt64(key string, fallback int64) int64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseInt(value, decimalBase, bitSize64)
	if err != nil {
		return fallback
	}

	return int64(result)
}

// Returns int64 value from env. If key is not set, it will panic.
func MustGetInt64(key string) int64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, key)
	}

	result, err := strconv.ParseInt(value, decimalBase, bitSize64)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return int64(result)
}
