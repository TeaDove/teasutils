package settings_utils

import (
	"os"
	"testing"
	"time"

	"github.com/teadove/teasutils/utils/test_utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func writeToEnvFile(t *testing.T, v string) {
	t.Helper()

	filePath := getEnvFilePath()
	_ = os.Remove(filePath)

	file, err := os.Create(filePath)
	require.NoError(t, err)

	_, err = file.WriteString(v)
	require.NoError(t, err)
}

func TestUnit_Settings_Init_Ok(t *testing.T) {
	type Settings struct {
		User     string `env:"user"     envDefault:"masha"                             json:"user"`
		Password string `env:"password" envDefault:"thebestpasswordeverveeeeeeerylong" json:"password"`
	}

	t.Setenv("teas_user", "julia")

	settings, err := GetSettings[Settings](test_utils.GetLoggedContext(), "teas_")
	require.NoError(t, err)
	assert.Equal(t, "julia", settings.User)
}

//nolint: paralleltest // working with files
func TestUnit_Settings_InitFromFile_Ok(t *testing.T) {
	type Settings struct {
		User     string `env:"user"     envDefault:"masha"               json:"user"`
		Password string `env:"password" envDefault:"thebestpasswordever" json:"password"`
	}

	writeToEnvFile(t, `teas_user=julia`)

	settings, err := GetSettings[Settings](test_utils.GetLoggedContext(), "teas_")
	require.NoError(t, err)
	assert.Equal(t, "julia", settings.User)

	_ = os.Remove(getEnvFilePath())
}

//nolint: paralleltest // working with files
func TestUnit_Settings_PanicFromCorruptedFile_Ok(t *testing.T) {
	type Settings struct {
		User     string `env:"user"     envDefault:"masha"               json:"user"`
		Password string `env:"password" envDefault:"thebestpasswordever" json:"password"`
	}

	writeToEnvFile(t, `teas_user;julia`)

	_, err := GetSettings[Settings](test_utils.GetLoggedContext(), "teas_")
	require.Error(t, err)

	_ = os.Remove(getEnvFilePath())
}

//nolint: paralleltest // working with files
func TestUnit_Settings_TimeSetted_Ok(t *testing.T) {
	assert.NotEmpty(t, ServiceSettings.StartedAt)
}

func TestUnit_Settings_SetServiceName_Ok(t *testing.T) {
	t.Setenv("HOSTNAME", "device-get-event-bf6ff5d47-qbs4f")

	settings := MustGetSetting[serviceSettings](test_utils.GetLoggedContext(), "BASE_")
	setServiceName(settings)

	assert.Equal(t, "device-get-event", settings.ServiceName)
}

//nolint: paralleltest // working with files
func TestUnit_Settings_Refresh_Ok(t *testing.T) {
	//nolint: tenv // needs
	err := os.Setenv("ENV_REFRESH_INTERVAL_S", "1")
	require.NoError(t, err)

	writeToEnvFile(t, `BASE_RELEASE=false`)

	settings := MustGetSetting[serviceSettings](test_utils.GetLoggedContext(), "BASE_")
	assert.False(t, settings.Release)

	writeToEnvFile(t, `BASE_RELEASE=true`)
	time.Sleep(2 * time.Second)
	assert.True(t, settings.Release)

	_ = os.Remove(getEnvFilePath())
}
