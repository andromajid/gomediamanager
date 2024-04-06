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
func scanDirectory(directory string) {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isMediaFile(path) {
			fmt.Println("Media File:", path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error scanning directory:", err)
	}
}
