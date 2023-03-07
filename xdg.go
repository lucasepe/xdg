package xdg

import (
	"os"
	"path/filepath"

	"github.com/lucasepe/xdg/pathutil"
)

// XDG Base Directory environment variables.
const (
	envDataHome   = "XDG_DATA_HOME"
	envConfigHome = "XDG_CONFIG_HOME"
	envCacheHome  = "XDG_CACHE_HOME"
)

func xdgPath(name, defaultPath string) string {
	dir := pathutil.ExpandHome(os.Getenv(name), HomeDir())
	if dir != "" && filepath.IsAbs(dir) {
		return dir
	}

	return defaultPath
}
