package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Media struct {
	Id       int64
	Name     string
	Location string
	Type     string
	Hash     string
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
// scanDirectory recursively scans the given directory for media files

	func ScanDirectoryParallel(directory string) {
		cwalk.NumWorkers = runtime.GOMAXPROCS(1)
		var files = []file{}
		err := cwalk.WalkWithSymlinks(directory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && isMediaFile(path) {
				fmt.Println("Media File:", path)
				files = append(files, file{name: info.Name(), location: path, Type: "movie", hash: hashFile(path)})
			}
			return nil
		})

		if err != nil {
			//fmt.Println("Error scanning directory:", err)
		}
		fmt.Printf("%v", files)
	}
*/
func ScanDirectory(directory string, db Db) {
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
			sqlResult, err := db.AddFile(Media{Name: info.Name(), Location: filepath.Dir(path), Type: "movie", Hash: hashFile(path)})
			if err != nil {
				fmt.Println("Succesfully inserted files : ", path, " sql result : ", sqlResult)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error scanning directory:", err)
	}
}

func hashFile(filePath string) string {
	file, _ := os.Open(filePath)
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
