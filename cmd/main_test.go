package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Just mock
func TestMainFunc(t *testing.T) {
	t.Run("FuncUnderTest", func(t *testing.T) {
		actual := true
		expected := true

		require.Equal(t, expected, actual)
	})
}
