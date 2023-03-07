package xdg

import (
	"path/filepath"

	"github.com/lucasepe/xdg/pathutil"
	"golang.org/x/sys/windows"
)

func HomeDir() string {
	return pathutil.KnownFolder(
		windows.FOLDERID_Profile,
		[]string{"USERPROFILE"},
		nil,
	)
}

func DataDir() string {
	return xdgPath(envDataHome, localAppData())
}

func ConfigDir() string {
	return xdgPath(envConfigHome, localAppData())
}

func CacheDir() string {
	return xdgPath(envCacheHome, filepath.Join(localAppData(), "cache"))
}

func localAppData() string {
	return pathutil.KnownFolder(
		windows.FOLDERID_LocalAppData,
		[]string{"LOCALAPPDATA"},
		[]string{filepath.Join(HomeDir(), "AppData", "Local")},
	)
}
