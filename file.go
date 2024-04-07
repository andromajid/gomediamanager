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

/*
   // Walk the directory tree
   err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
       if err != nil {
           return err
       }

       // Check if the current path is a symbolic link
       if info.Mode()&os.ModeSymlink != 0 {
           // Resolve the symbolic link
           targetPath, err := filepath.EvalSymlinks(path)
           if err != nil {
               log.Printf("Error resolving symlink %s: %v", path, err)
               return nil // Skip this symlink
           }

           // Check if the target is a directory
           targetInfo, err := os.Stat(targetPath)
           if err != nil {
               log.Printf("Error stating %s: %v", targetPath, err)
               return nil // Skip this symlink
           }

           if targetInfo.IsDir() {
               // If the target is a directory, walk it
               return filepath.Walk(targetPath, func(innerPath string, innerInfo os.FileInfo, innerErr error) error {
                   if innerErr != nil {
                       return innerErr
                   }

                   // Check if the inner path is a file and matches the media regex pattern
                   if !innerInfo.IsDir() && mediaRegEx.MatchString(innerInfo.Name()) {
                       fmt.Println(innerPath)
                   }

                   return nil
               })
           }
       } else if !info.IsDir() && mediaRegEx.MatchString(info.Name()) {
           fmt.Println(path)
       }

       return nil
   })

   if err != nil {
       log.Println(err)
   }
*/
// scanDirectory recursively scans the given directory for media files
func ScanDirectory(directory string) {
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
