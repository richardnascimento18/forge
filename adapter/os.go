/*
Package adapter solves cross-platform issues
*/
package adapter

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func ResolveLogDir() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return linuxLogDir()
	case "darwin":
		return macosLogDir()
	case "windows":
		return windowsLogDir()
	default:
		return "", errors.New("unsupported OS")
	}
}

func linuxLogDir() (string, error) {
	if v := os.Getenv("XDG_STATE_HOME"); v != "" {
		return filepath.Join(v, "forge", "logs"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".local", "state", "forge", "logs"), nil
}

func macosLogDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, "Library", "Logs", "Forge"), nil
}

func windowsLogDir() (string, error) {
	if v := os.Getenv("LOCALAPPDATA"); v != "" {
		return filepath.Join(v, "Forge", "Logs"), nil
	}

	return "", errors.New("LOCALAPPDATA not set")
}
