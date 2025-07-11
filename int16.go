package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

// Returns int16 value from env. If key is not set, it returns the fallback.
func GetInt16(key string, fallback int16) int16 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseInt(value, decimalBase, bitSize16)
	if err != nil {
		return fallback
	}

	return int16(result)
}

// Returns int16 value from env. If key is not set, it will panic.
func MustGetInt16(key string) int16 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, key)
	}

	result, err := strconv.ParseInt(value, decimalBase, bitSize16)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return int16(result)
}
