package hex

import (
	"testing"

	secrets2 "github.com/neodejack/CaptainJack/hex/adapters/frameworks/driven/secrets"
	"github.com/stretchr/testify/require"
)

func TestSecrets_GetTeleToken(t *testing.T) {
	secrets := secrets2.NewAdapterGCP()

	teleBotSecret, err := secrets.GetTeleToken()

	require.Nil(t, err)
	require.Equal(t, teleBotSecret,
		"5402144533:AAGzoxfahqea1L3SyFgrHftm-uGkUWkj2xw")
}
