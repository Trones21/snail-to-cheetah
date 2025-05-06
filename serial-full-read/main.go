// Note: This is a cut down version of the front matter checker... for the full version with all the flags, see the front matter checker github
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

func main() {
	// Define a command-line flag for the directory
	pathPtr := flag.String("path", "./", "Path to the directory or file to process")
	jsonTemplatePath := flag.String("json", "./template.json", "Location of the template json file")
	flag.Parse()

	path := *pathPtr
	// Check if the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Error: The path %s does not exist.\n", path)
		return
	}

	jsonUri := *jsonTemplatePath
	// Check if the path exists
	if _, err := os.Stat(jsonUri); os.IsNotExist(err) {
		fmt.Printf("Error: The file %s does not exist.\n", jsonUri)
		return
	}
	loadTemplate(jsonUri)

	if templateKeys == nil {
		fmt.Print("templateKeys is nil")
		return
	}

	go func() {
		log.Println("pprof available at http://localhost:6060/debug/pprof/")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Printf("Starting Serial Full Read for path: %s\n", path)

	// Collect all file paths
	paths, err := collectFilePaths(path, ".md")
	if err != nil {
		fmt.Printf("Error collecting file paths: %v\n", err)
		return
	}
	fmt.Printf("File Count: %v\n", len(paths))
	// Process each file
	start := time.Now()
	for _, path := range paths {
		err := processFile(path)
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", path, err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Processed %s in %v\n", path, elapsed)
}

// collectFilePaths collects all file paths from a given directory or single file that match a given extension (case-insensitive).
func collectFilePaths(path string, ext string) ([]string, error) {
	var paths []string
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Normalize the extension to lowercase for case-insensitive comparison
	ext = strings.ToLower(ext)

	if info.IsDir() {
		err = filepath.Walk(path, func(p string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Check file extension (case-insensitive)
			if !f.IsDir() && strings.ToLower(filepath.Ext(f.Name())) == ext {
				paths = append(paths, p)
			}
			return nil
		})
	} else if strings.ToLower(filepath.Ext(info.Name())) == ext {
		paths = append(paths, path) // Add single file to paths if it matches the extension
	}

	return paths, err
}

var templateKeys map[string]interface{}

// loadTemplate loads the JSON template once for all files.
func loadTemplate(templatePath string) error {
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read template file: %w", err)
	}

	templateKeys = make(map[string]interface{})
	err = json.Unmarshal(templateContent, &templateKeys)
	if err != nil {
		return fmt.Errorf("failed to parse JSON template: %w", err)
	}

	return nil
}

// processFile processes a file, checking its front matter against the loaded JSON template.
func processFile(path string) error {
	// Read the file content
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Extract front matter

	frontMatter, err := extractFrontMatterBoundary(string(content))
	if err != nil {
		return fmt.Errorf("failed to extract front matter. Error: %w", err)
	}

	// Parse the front matter into a map
	var frontMatterKVs map[string]interface{}
	err = yaml.Unmarshal([]byte(frontMatter), &frontMatterKVs)
	if err != nil {
		log.Fatalf("failed to parse YAML front matter: %v", err)
	}

	// Check for missing keys
	missingKeys := []string{}
	for key := range templateKeys {
		if _, exists := frontMatterKVs[key]; !exists {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) > 0 {
		// Currently they are all invalid as the template is pretty static
		//fmt.Printf("File: %s - Missing keys in front matter: %v\n", path, missingKeys)
	} else {
		fmt.Printf("File: %s - Front matter is valid.\n", path)
	}

	return nil
}

// extractFrontMatterRegex extracts the front matter using a regex pattern.
func extractFrontMatterRegex(content string) (string, error) {
	// Regular expression to match front matter enclosed by --- or +++
	regex := regexp.MustCompile(`(?s)^---\n(.*?)\n---`)
	matches := regex.FindStringSubmatch(content)
	if len(matches) < 2 {
		return "", errors.New("front matter not found")
	}
	return matches[1], nil
}

// extractFrontMatterBoundary extracts the front matter by reading up to the second ---.
func extractFrontMatterBoundary(content string) (string, error) {
	// Normalize line endings to \n to handle different platforms
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	lines := strings.Split(content, "\n")

	if len(lines) < 2 || lines[0] != "---" {
		return "", fmt.Errorf("front matter start delimiter not found. First line: %s", lines[0])
	}

	var frontMatterLines []string
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			return strings.Join(frontMatterLines, "\n"), nil
		}
		frontMatterLines = append(frontMatterLines, lines[i])
	}

	return "", errors.New("front matter end delimiter not found")
}
