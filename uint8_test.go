package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetUint8(t *testing.T) {
	var fallback uint8 = 10

	t.Run("GetNonExistingUint8WithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetUint8("KEY", fallback))
	})

	t.Run("GetInvalidUint8WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetUint8("KEY", fallback))
	})

	t.Run("GetValidUint8WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint8(12), envy.GetUint8("KEY", fallback))
	})
}

func TestMustGetUint8(t *testing.T) {
	t.Run("MustGetNonExistingUint8", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetUint8("KEY")
		})
	})

	t.Run("MustGetInvalidUint8", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetUint8("KEY")
		})
	})

	t.Run("MustGetValidUint8", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint8(12), envy.MustGetUint8("KEY"))
	})
}
