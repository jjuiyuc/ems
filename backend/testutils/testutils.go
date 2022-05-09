package testutils

import (
	"path/filepath"
	"runtime"
)

func GetConfigDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "..", "config")
}
