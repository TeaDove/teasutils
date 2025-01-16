package settings_utils

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Settings_Init_Ok(t *testing.T) {
	type Settings struct {
		User     string `env:"user"     envDefault:"masha"               json:"user"`
		Password string `env:"password" envDefault:"thebestpasswordever" json:"password"`
	}

	t.Setenv("teas_user", "julia")

	settings, err := InitSetting[Settings](context.Background(), "teas_", "password")
	require.NoError(t, err)
	assert.Equal(t, "julia", settings.User)
}

// nolint: paralleltest // working with files
func TestUnit_Settings_InitFromFile_Ok(t *testing.T) {
	type Settings struct {
		User     string `env:"user"     envDefault:"masha"               json:"user"`
		Password string `env:"password" envDefault:"thebestpasswordever" json:"password"`
	}

	_ = os.Remove(envFile)

	file, err := os.Create(envFile)
	require.NoError(t, err)

	_, err = file.WriteString(`teas_user=julia`)
	require.NoError(t, err)

	settings, err := InitSetting[Settings](context.Background(), "teas_", "password")
	require.NoError(t, err)
	assert.Equal(t, "julia", settings.User)

	err = os.Remove(envFile)
	require.NoError(t, err)
}

// nolint: paralleltest // working with files
func TestUnit_Settings_PanicFromCorruptedFile_Ok(t *testing.T) {
	type Settings struct {
		User     string `env:"user"     envDefault:"masha"               json:"user"`
		Password string `env:"password" envDefault:"thebestpasswordever" json:"password"`
	}

	_ = os.Remove(envFile)

	file, err := os.Create(envFile)
	require.NoError(t, err)

	_, err = file.WriteString(`teas_user;julia`)
	require.NoError(t, err)

	assert.Panics(t, func() {
		_, _ = InitSetting[Settings](context.Background(), "teas_", "password")
	})

	err = os.Remove(envFile)
	require.NoError(t, err)
}

// nolint: paralleltest // working with files
func TestUnit_Settings_TimeSetted_Ok(t *testing.T) {
	assert.NotEmpty(t, BaseSettings.StartedAt)
}

func TestUnit_Settings_SetServiceName_Ok(t *testing.T) {
	t.Parallel()

	os.Setenv("HOSTNAME", "device-get-event-bf6ff5d47-qbs4f")
	settings := MustInitSetting[baseSettings](context.Background(), "BASE_")
	setServiceName(&settings)

	assert.Equal(t, "device-get-event", settings.ServiceName)
}
