package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

// Returns float32 value from env. If key is not set, it returns the fallback.
func GetFloat32(key string, fallback float32) float32 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseFloat(value, bitSize32)
	if err != nil {
		return fallback
	}

	return float32(result)
}

// Returns float32 value from env. If key is not set, it will panic.
func MustGetFloat32(key string) float32 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, key)
	}

	result, err := strconv.ParseFloat(value, bitSize32)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return float32(result)
}
