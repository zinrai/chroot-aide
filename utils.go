package main

import (
	"fmt"
	"os"
)

// fileExists checks if a file exists
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// dirExists checks if a directory exists
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// pathExists checks if a path (file or directory) exists
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ensureDir creates a directory if it doesn't exist
func ensureDir(path string, perm os.FileMode) error {
	if dirExists(path) {
		return nil
	}

	if err := os.MkdirAll(path, perm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}

	return nil
}

// removeIfExists removes a file or directory if it exists
func removeIfExists(path string) error {
	if !pathExists(path) {
		return nil
	}

	return os.RemoveAll(path)
}
