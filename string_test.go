package envy_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/pyr33x/envy"
)

func TestGetString(t *testing.T) {
	fallback := "envy"

	t.Run("GetNonExistingStringWithFallback", func(t *testing.T) {
		assert.Equal(t, fallback, envy.GetString("KEY", fallback))
	})

	t.Run("GetValidStringWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "VALID")
		assert.Equal(t, "VALID", envy.GetString("KEY", fallback))
	})
}

func TestMustGetString(t *testing.T) {
	t.Run("MustGetNonExistingString", func(t *testing.T) {
		assert.Panics(t, func() {
			envy.MustGetString("KEY")
		})
	})

	t.Run("MustGetValidString", func(t *testing.T) {
		t.Setenv("KEY", "VALID")
		assert.Equal(t, "VALID", envy.MustGetString("KEY"))
	})
}
