package envy

import (
	"os"
	"strconv"

	er "github.com/pyr33x/envy/pkg/err"
)

func GetFloat64(key string, fallback float64) float64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := strconv.ParseFloat(value, bitSize64)
	if err != nil {
		return fallback
	}

	return float64(result)
}

func MustGetFloat64(key string) float64 {
	value, ok := os.LookupEnv(key)
	if !ok {
		er.Throw(er.NotSet, value)
	}

	result, err := strconv.ParseFloat(value, bitSize64)
	if err != nil {
		er.Throw(er.NotSet, key)
	}

	return float64(result)
}
