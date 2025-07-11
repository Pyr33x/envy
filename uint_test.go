package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetUint(t *testing.T) {
	var fallback uint = 10

	t.Run("GetNonExistingUintWithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetUint("KEY", fallback))
	})

	t.Run("GetInvalidUintWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetUint("KEY", fallback))
	})

	t.Run("GetValidUintWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint(12), envy.GetUint("KEY", fallback))
	})
}

func TestMustGetUint(t *testing.T) {
	t.Run("MustGetNonExistingUint", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetUint("KEY")
		})
	})

	t.Run("MustGetInvalidUint", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetUint("KEY")
		})
	})

	t.Run("MustGetValidUint", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint(12), envy.MustGetUint("KEY"))
	})
}
