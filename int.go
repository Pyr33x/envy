package envy

// Returns int value from env. If key is not set, it returns the fallback.
func GetInt(key string, fallback int) int {
	return int(GetInt64(key, int64(fallback)))
}

// Returns int value from env. If key is not set, it will panic.
func MustGetInt(key string) int {
	return int(MustGetInt64(key))
}
