package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetInt(t *testing.T) {
	var fallback int = 10

	t.Run("GetNonExistingIntWithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetInt("KEY", fallback))
	})

	t.Run("GetInvalidIntWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetInt("KEY", fallback))
	})

	t.Run("GetValidIntWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int(12), envy.GetInt("KEY", fallback))
	})
}

func TestMustInt(t *testing.T) {
	t.Run("MustGetNonExistingInt", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetInt("KEY")
		})
	})

	t.Run("MustGetInvalidInt", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetInt("KEY")
		})
	})

	t.Run("MustGetValidInt", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int(12), envy.MustGetInt("KEY"))
	})
}
