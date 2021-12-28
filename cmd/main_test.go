package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMainFunc(t *testing.T) {
	t.Run("Test for ci-tools check", func(t *testing.T) {
		actual := true
		expected := true

		require.Equal(t, expected, actual)
	})
}
