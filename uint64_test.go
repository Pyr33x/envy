package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetUint64(t *testing.T) {
	var fallback uint64 = 10

	t.Run("GetNonExistingUint64WithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetUint64("KEY", fallback))
	})

	t.Run("GetInvalidUint64WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetUint64("KEY", fallback))
	})

	t.Run("GetValidUint64WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint64(12), envy.GetUint64("KEY", fallback))
	})
}

func TestMustGetUint64(t *testing.T) {
	t.Run("MustGetNonExistingUint64", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetUint64("KEY")
		})
	})

	t.Run("MustGetInvalidUint64", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetUint64("KEY")
		})
	})

	t.Run("MustGetValidUint64", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint64(12), envy.MustGetUint64("KEY"))
	})
}
