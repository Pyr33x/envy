package envy_test

import (
	"testing"

	"github.com/pyr33x/envy"
	"github.com/stretchr/testify/assert"
)

func TestGetInt16(t *testing.T) {
	var fallback int16 = 10

	t.Run("GetNonExistingInt16WithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetInt16("KEY", fallback))
	})

	t.Run("GetInvalidInt16WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Equal(t, fallback, envy.GetInt16("KEY", fallback))
	})

	t.Run("GetValidInt16WithFallback", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int16(12), envy.GetInt16("KEY", fallback))
	})
}

func TestMustInt16(t *testing.T) {
	t.Run("MustGetNonExistingInt16", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetInt16("KEY")
		})
	})

	t.Run("MustGetInvalidInt16", func(t *testing.T) {
		t.Setenv("KEY", "INVALID")
		assert.Panics(t, func() {
			envy.MustGetInt16("KEY")
		})
	})

	t.Run("MustGetValidInt16", func(t *testing.T) {
		t.Setenv("KEY", "12")
		assert.Equal(t, int16(12), envy.MustGetInt16("KEY"))
	})
}
