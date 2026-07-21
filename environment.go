package main

import (
	"fmt"
	"path/filepath"
)

// getOverlayDir returns the overlay directory path for a given chroot directory and name
func getOverlayDir(chrootDir string, overlayName string) string {
	if overlayName == "" {
		return ""
	}
	return fmt.Sprintf("%s.%s", chrootDir, overlayName)
}

// isOverlaySetup checks if a specific overlay is already set up
func isOverlaySetup(chrootDir string, overlayName string) bool {
	overlayDir := getOverlayDir(chrootDir, overlayName)
	if overlayDir == "" {
		return false
	}
	mergedPath := filepath.Join(overlayDir, MergedDir)

	// Check if overlay directory exists and merged is mounted
	return dirExists(overlayDir) && isMounted(mergedPath)
}

// validateChrootStructure validates that the chroot directory has the required structure
func validateChrootStructure(chrootDir string) error {
	// Check if the directory exists
	if !dirExists(chrootDir) {
		return fmt.Errorf("chroot directory %s does not exist", chrootDir)
	}

	// For a minimal check, just ensure it's a directory
	// Additional validation can be added here if needed
	return nil
}
