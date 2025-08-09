package pathhandler

import (
	"os"
	"path/filepath"
)

func ProjectPath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	return exeDir, nil
}

func FontDirPath() (string, error) {
	exeDir, err := ProjectPath()
	if err != nil {
		return "", err
	}
	fontPath := filepath.Join(exeDir, "cfg", "fonts")
	return fontPath, nil
}
