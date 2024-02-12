package kzg

import (
	"testing"

	"github.com/KP-Universe/prysm/v4/testing/require"
)

func TestStart(t *testing.T) {
	require.NoError(t, Start())
	require.NotNil(t, kzgContext)
}
