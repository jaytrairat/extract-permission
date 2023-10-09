package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/jaytrairat/extract-permission/constant"
	"github.com/spf13/cobra"
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

func extractPermission() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " :: Start acquire permissions")
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

	preparedText := strings.Join(result, "\n")
	clipboard.WriteAll(preparedText)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " :: All permissions copied to clipboard")
}

var rootCmd = &cobra.Command{
	Use:   "extract-permission",
	Short: "A tool for extracting permissions from AndroidManifest.xml files",
	Run: func(cmd *cobra.Command, args []string) {
		extractPermission()
	},
}

func main() {
	rootCmd.Execute()
}
