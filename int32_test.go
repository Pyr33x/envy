package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetInt32(t *testing.T) {
	var fallback int32 = 10

	t.Run("GetNonExistingInt32WithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetInt32("KEY", fallback))
	})

	t.Run("GetInvalidInt32WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetInt32("KEY", fallback))
	})

	t.Run("GetValidInt32WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int32(12), envy.GetInt32("KEY", fallback))
	})
}

func TestMustInt32(t *testing.T) {
	t.Run("MustGetNonExistingInt32", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetInt32("KEY")
		})
	})

	t.Run("MustGetInvalidInt32", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetInt32("KEY")
		})
	})

	t.Run("MustGetValidInt32", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int32(12), envy.MustGetInt32("KEY"))
	})
}
