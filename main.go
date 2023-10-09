package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func findManifestFile(rootDir string) string {
	manifestFiles := ""
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		manifestFiles = "TEST"
		return nil
	})
	return manifestFiles
}

func main() {
	targetFile := findManifestFile("./")
	fmt.Println(targetFile)
}
