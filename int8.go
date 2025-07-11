package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

// Returns int8 value from env. If key is not set, it returns the fallback.
func GetInt8(key string, fallback int8) int8 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseInt(value, decimalBase, bitSize8)
	if err != nil {
		return fallback
	}

	return int8(result)
}

// Returns int8 value from env. If key is not set, it will panic.
func MustGetInt8(key string) int8 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, key)
	}

	result, err := strconv.ParseInt(value, decimalBase, bitSize8)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return int8(result)
}
