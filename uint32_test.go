package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetUint32(t *testing.T) {
	var fallback uint32 = 10

	t.Run("GetNonExistingUint32WithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetUint32("KEY", fallback))
	})

	t.Run("GetInvalidUint32WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetUint32("KEY", fallback))
	})

	t.Run("GetValidUint32WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint32(12), envy.GetUint32("KEY", fallback))
	})
}

func TestMustGetUint32(t *testing.T) {
	t.Run("MustGetNonExistingUint32", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetUint32("KEY")
		})
	})

	t.Run("MustGetInvalidUint32", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetUint32("KEY")
		})
	})

	t.Run("MustGetValidUint32", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint32(12), envy.MustGetUint32("KEY"))
	})
}
