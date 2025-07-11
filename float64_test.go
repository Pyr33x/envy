package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetFloat64(t *testing.T) {
	var fallback float64 = 10.5
	envValue := "12.5"

	t.Run("GetNonExistingFloat64WithFallback", func(t *testing.T) {
		assert.InDelta(t, fallback, envy.GetFloat64("KEY", fallback), 0)
	})

	t.Run("GetInvalidFloat64WithFallback", func(t *testing.T) {
		t.Setenv("KEY", envValue)
		assert.InDelta(t, float64(12.5), envy.GetFloat64("KEY", fallback), 0)
	})

	t.Run("GetValidFloat64WithFallback", func(t *testing.T) {
		t.Setenv("KEY", envValue)
		assert.InDelta(t, float64(12.5), envy.GetFloat64("KEY", fallback), 0)
	})
}

func TestMustGetFloat64(t *testing.T) {
	envValue := "12.5"

	t.Run("MustGetNonExistingFloat64", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetFloat64("KEY")
		})
	})

	t.Run("MustGetInvalidFloat64", func(t *testing.T) {
		assert.Panics(t, func() {
			t.Setenv("KEY", "INVALID")
			envy.MustGetFloat64("KEY")
		})
	})

	t.Run("MustGetValidFloat64", func(t *testing.T) {
		t.Setenv("KEY", envValue)
		assert.InDelta(t, float64(12.5), envy.MustGetFloat64("KEY"), 0)
	})
}
