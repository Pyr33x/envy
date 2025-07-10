package envy

import (
	"os"

	"github.com/pyr33x/envy/pkg/err"
)

// Returns boolean value from env. If key is not set, it returns the fallback.
func GetBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val == "1" || val == "true"
}

// Returns boolean value from env. If key is not set, it will panic.
func MustGetBool(key string) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		err.Throw(err.NotSet, key)
	}

	return val == "1" || val == "true"
}
