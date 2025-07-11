package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetUint16(t *testing.T) {
	var fallback uint16 = 10

	t.Run("GetNonExistingUint16WithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetUint16("KEY", fallback))
	})

	t.Run("GetInvalidUint16WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetUint16("KEY", fallback))
	})

	t.Run("GetValidUint16WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint16(12), envy.GetUint16("KEY", fallback))
	})
}

func TestMustGetUint16(t *testing.T) {
	t.Run("MustGetNonExistingUint16", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetUint16("KEY")
		})
	})

	t.Run("MustGetInvalidUint16", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetUint16("KEY")
		})
	})

	t.Run("MustGetValidUint16", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, uint16(12), envy.MustGetUint16("KEY"))
	})
}
