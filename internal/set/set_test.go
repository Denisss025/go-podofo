package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/denisss025/go-podofo/internal/set"
)

func TestSet(t *testing.T) {
	t.Parallel()

	t.Run("put", func(t *testing.T) {
		t.Parallel()

		const n = 10

		s := set.New[int]()
		assert.False(t, s.Contains(n))

		s.Put(n)
		assert.True(t, s.Contains(n))

		s.Put(n)
		assert.True(t, s.Contains(n))

		s.Remove(n)
		assert.False(t, s.Contains(n))

		s.Remove(n)
		assert.False(t, s.Contains(n))
	})
}
