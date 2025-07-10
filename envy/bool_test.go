package envy_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/pyr33x/envy/envy"
)

func TestGetBool(t *testing.T) {
	t.Run("GetEmptyBoolWithFallback", func(t *testing.T) {
		assert.True(t, envy.GetBool("KEY", true))
	})

	t.Run("GetExistingBoolWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "false")
		assert.False(t, envy.GetBool("KEY", true))
	})

	t.Run("GetZeroAsBoolWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "0")
		assert.False(t, envy.GetBool("KEY", true))
	})

	t.Run("GetOneAsBoolWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "1")
		assert.True(t, envy.GetBool("KEY", false))
	})

	t.Run("GetTrueStringAsBoolWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "true")
		assert.True(t, envy.GetBool("KEY", false))
	})

	t.Run("GetInvalidBoolStringWithFallback", func(t *testing.T) {
		t.Setenv("KEY", "maybe")
		assert.False(t, envy.GetBool("KEY", false))
	})
}
