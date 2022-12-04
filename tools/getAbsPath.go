package tools

import (
	"os"
	"path/filepath"
)

func GetAbsPath(name string) (string, error) {
	absPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(absPath, name), nil
}
