package operations_test

import (
	"learningmath/pkg/operations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		solution := operations.Add(1, 1)
		assert.Equal(t, 2, solution)
	})

	t.Run("should return 10", func(t *testing.T) {
		solution := operations.Add(5, 5)
		assert.Equal(t, 10, solution)

	})
}
