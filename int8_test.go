package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetInt8(t *testing.T) {
	var fallback int8 = 10

	t.Run("GetNonExistingInt8WithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetInt8("KEY", fallback))
	})

	t.Run("GetInvalidInt8WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetInt8("KEY", fallback))
	})

	t.Run("GetValidInt8WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int8(12), envy.GetInt8("KEY", fallback))
	})
}

func TestMustInt8(t *testing.T) {
	t.Run("MustGetNonExistingInt8", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetInt8("KEY")
		})
	})

	t.Run("MustGetInvalidInt8", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetInt8("KEY")
		})
	})

	t.Run("MustGetValidInt8", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int8(12), envy.MustGetInt8("KEY"))
	})
}
