package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetInt64(t *testing.T) {
	var fallback int64 = 10

	t.Run("GetNonExistingInt64WithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetInt64("KEY", fallback))
	})

	t.Run("GetInvalidInt64WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetInt64("KEY", fallback))
	})

	t.Run("GetValidInt64WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int64(12), envy.GetInt64("KEY", fallback))
	})
}

func TestMustInt64(t *testing.T) {
	t.Run("MustGetNonExistingInt64", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetInt64("KEY")
		})
	})

	t.Run("MustGetInvalidInt64", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetInt64("KEY")
		})
	})

	t.Run("MustGetValidInt64", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int64(12), envy.MustGetInt64("KEY"))
	})
}
