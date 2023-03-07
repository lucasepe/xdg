//go:build aix || dragonfly || freebsd || (js && wasm) || nacl || linux || netbsd || openbsd || solaris
// +build aix dragonfly freebsd js,wasm nacl linux netbsd openbsd solaris

package xdg_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lucasepe/xdg"
	"github.com/stretchr/testify/require"
)

func TestCustomBaseDirs(t *testing.T) {
	home := xdg.HomeDir()

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			value:    "~/.local/share",
			expected: filepath.Join(home, ".local/share"),
			actual:   xdg.DataDir,
		},

		&envSample{
			name:     "XDG_CONFIG_HOME",
			value:    "~/.local/config",
			expected: filepath.Join(home, ".local/config"),
			actual:   xdg.ConfigDir,
		},

		&envSample{
			name:     "XDG_CACHE_HOME",
			value:    "~/.local/cache",
			expected: filepath.Join(home, ".local/cache"),
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
