//go:build aix || dragonfly || freebsd || (js && wasm) || nacl || linux || netbsd || openbsd || solaris
// +build aix dragonfly freebsd js,wasm nacl linux netbsd openbsd solaris

package xdg

import (
	"os"
	"path/filepath"
)

func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}

	return "/"
}

func DataDir() string {
	return xdgPath(envDataHome, filepath.Join(HomeDir(), ".local", "share"))
}

func ConfigDir() string {
	return xdgPath(envConfigHome, filepath.Join(HomeDir(), ".config"))
}

func CacheDir() string {
	return xdgPath(envCacheHome, filepath.Join(HomeDir(), ".cache"))
}
