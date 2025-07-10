package envy

import (
	"os"

	"github.com/pyr33x/envy/pkg/err"
)

// Returns string value from env. If key is not set, it returns the fallback.
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

// Returns string value from env, If key is not set, it will panic.
func MustGetString(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		err.Throw(err.NotSet, key)
	}

	return val
}
