package echoing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEchoing(t *testing.T) {
	coreAdapter := NewAdapter()

	echoedWords, err := coreAdapter.Echoing("test string")
	require.Nil(t, err)
	require.Equal(t, echoedWords, "test string, so what?")

}
