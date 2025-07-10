package err_test

import (
	"errors"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/pyr33x/envy/pkg/err"
)

func TestThrow(t *testing.T) {
	t.Run("ThrowPanic", func(t *testing.T) {
		assert.Panics(t, func() {
			err.Throw(errors.New("Failed"), "KEY")
		})
	})
}
