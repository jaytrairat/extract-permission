package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jaytrairat/extract-permission/constant"
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
	fmt.Println()
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
					currentLine = strings.TrimSpace(currentLine)
					extractedPermission := regexp.MustCompile(`android:name="([^"]+)"`).FindStringSubmatch(currentLine)
					if len(extractedPermission) > 1 {
						currentLine = extractedPermission[1]
					}
					splitedPermission := strings.Split(currentLine, ".")
					if len(splitedPermission) == 3 {
						permissionIndex := -1
						for i, item := range constant.PERMISSION {
							if item["TYPE"] == splitedPermission[2] {
								permissionIndex = i
								break
							}
						}
						if permissionIndex != -1 {
							currentLine = fmt.Sprintf("%s\t%s", splitedPermission[2], constant.PERMISSION[permissionIndex]["DESCRIPTION"])
						} else {
							currentLine = fmt.Sprintf("%s\t", splitedPermission[2])
						}
					}
					result = append(result, currentLine)
				}
			}
		}
	}

	fmt.Printf(strings.Join(result, "\n"))
}
