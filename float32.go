package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

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
