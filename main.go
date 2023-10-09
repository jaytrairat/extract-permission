package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func findManifestFile(rootDir string) []string {
	manifestFiles := []string{}
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "AndroidManifest.xml" {
			manifestFiles = append(manifestFiles, path)
		}
		return nil
	})
	return manifestFiles
}

func main() {
	targetFile := findManifestFile("./")
	fmt.Println(targetFile)
}
