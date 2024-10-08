package settings_utils

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Settings_Init_Ok(t *testing.T) {
	t.Parallel()

	type Settings struct {
		User     string `env:"user"     envDefault:"masha"               json:"user"`
		Password string `env:"password" envDefault:"thebestpasswordever" json:"password"`
	}

	err := os.Setenv("teas_user", "julia")
	require.NoError(t, err)

	settings, err := InitSetting[Settings](context.Background(), "password")
	assert.NoError(t, err)
	assert.Equal(t, "julia", settings.User)
}
