package flags

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlags_Load(t *testing.T) {
	assert := assert.New(t)

	t.Run("when there are no params", func(t *testing.T) {
		expectedErr := errors.New("environment flag is required")
		t.Run("it returns an error", func(t *testing.T) {
			flags, err := Load()

			assert.Equal(expectedErr, err)
			assert.Nil(flags)
		})
	})

}
