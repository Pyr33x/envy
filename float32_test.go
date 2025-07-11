package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetFloat32(t *testing.T) {
	var fallback float32 = 10.5
	envValue := "12.5"

	t.Run("GetNonExistingFloat32WithFallback", func(t *testing.T) {
		assert.InDelta(t, fallback, envy.GetFloat32("KEY", fallback), 0)
	})

	t.Run("GetInvalidFloat32WithFallback", func(t *testing.T) {
		t.Setenv("KEY", envValue)
		assert.InDelta(t, float32(12.5), envy.GetFloat32("KEY", fallback), 0)
	})

	t.Run("GetValidFloat32WithFallback", func(t *testing.T) {
		t.Setenv("KEY", envValue)
		assert.InDelta(t, float32(12.5), envy.GetFloat32("KEY", fallback), 0)
	})
}

func MustGetFloat32(t *testing.T) {
	envValue := "12.5"

	t.Run("MustGetNonExistingFloat32", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetFloat32("KEY")
		})
	})

	t.Run("MustGetInvalidFloat32", func(t *testing.T) {
		assert.Panics(t, func() {
			t.Setenv("KEY", "INVALID")
			envy.MustGetFloat32("KEY")
		})
	})

	t.Run("MustGetValidFloat32", func(t *testing.T) {
		t.Setenv("KEY", envValue)
		assert.InDelta(t, float32(12.5), envy.MustGetFloat32("KEY"), 0)
	})
}
