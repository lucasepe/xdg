//go:build windows
// +build windows

package xdg_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lucasepe/xdg"
	"github.com/stretchr/testify/require"
)

func TestDefaultBaseDirs(t *testing.T) {
	home := xdg.HomeDir()
	systemDrive := `C:\`

	localAppData := filepath.Join(home, "AppData", "Local")

	envSamples := []*envSample{
		{
			name:     "XDG_DATA_HOME",
			expected: localAppData,
			actual:   xdg.DataDir,
		},

		{
			name:     "XDG_CONFIG_HOME",
			expected: localAppData,
			actual:   xdg.ConfigDir,
		},

		{
			name:     "XDG_CACHE_HOME",
			expected: filepath.Join(localAppData, "cache"),
			actual:   xdg.CacheDir,
		},
	}

	// Test default environment.
	testDirs(t, envSamples...)

	// Test system drive not set.
	envSystemDrive := os.Getenv("SystemDrive")
	require.NoError(t, os.Unsetenv("SystemDrive"))
	testDirs(t, envSamples...)
	require.NoError(t, os.Setenv("SystemDrive", envSystemDrive))
}
