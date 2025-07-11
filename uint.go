package envy

// Returns uint value from env. If key is not set, it returns the fallback.
func GetUint(key string, fallback uint) uint {
	return uint(GetUint64(key, uint64(fallback)))
}

// Returns uint value from env. If key is not set, it will panic.
func MustGetUint(key string) uint {
	return uint(MustGetUint64(key))
}
