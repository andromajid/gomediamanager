package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type file struct {
	name     string
	location string
	Type     string
	hash     string
}

// isMediaFile checks if a given file is a media file based on its extension
func isMediaFile(file string) bool {
	mediaExtensions := []string{".mp3", ".mp4", ".avi", ".mkv", ".mov", ".flv"}
	for _, ext := range mediaExtensions {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}
	return false
}

// scanDirectory recursively scans the given directory for media files
func ScanDirectory(directory string) {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a symbolic link
		if info.Mode()&os.ModeSymlink != 0 {
			realPath, err := filepath.EvalSymlinks(path)
			if err != nil {
				fmt.Printf("Error resolving symlink %s: %v\n", path, err)
				return nil
			}
			if !isMediaFile(realPath) {
				return nil
			}
			fmt.Println("Media File (Symlink Resolved):", realPath)
			return nil
		}

		if !info.IsDir() && isMediaFile(path) {
			fmt.Println("Media File:", path)
		} else {
			fmt.Println("Not Media File : ", path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error scanning directory:", err)
	}
}
