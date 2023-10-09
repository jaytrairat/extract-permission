package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	targetFiles := findManifestFile("./")
	result := []string{}
	if len(targetFiles) != 0 {
		for _, currentFile := range targetFiles {
			file, err := os.Open(currentFile)
			if err != nil {
				os.Exit(1)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				currentLine := scanner.Text()
				if strings.Contains(currentLine, "uses-permission") {
					result = append(result, strings.TrimSpace(currentLine))
				}
			}
		}
	}

	fmt.Printf(strings.Join(result, "\n"))
}
