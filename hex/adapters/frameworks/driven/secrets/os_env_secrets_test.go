package secrets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdapterOS_GetTeleToken(t *testing.T) {
	secrets := NewAdapterOS()

	teleBotSecret := secrets.GetTeleToken()

	require.NotEmpty(t, teleBotSecret)
}
