//go:build darwin
// +build darwin

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

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			expected: filepath.Join(home, "Library", "Application Support"),
			actual:   xdg.DataDir,
		},

		&envSample{
			name:     "XDG_CONFIG_HOME",
			expected: filepath.Join(home, "Library", "Application Support"),
			actual:   xdg.ConfigDir,
		},

		&envSample{
			name:     "XDG_CACHE_HOME",
			expected: filepath.Join(home, "Library", "Caches"),
			actual:   xdg.CacheDir,
		},
	)
}

func TestHomeNotSet(t *testing.T) {
	envHomeVar := "HOME"
	envHomeVal := os.Getenv(envHomeVar)
	require.NoError(t, os.Unsetenv(envHomeVar))
	require.Equal(t, "/", xdg.HomeDir())

	require.NoError(t, os.Setenv(envHomeVar, envHomeVal))
	require.Equal(t, envHomeVal, xdg.HomeDir())
}
