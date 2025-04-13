package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Define required flags
	numFiles := flag.Int("num", 0, "Number of files to generate (required)")
	minSize := flag.Int("min", 0, "Minimum file size in bytes (required)")
	maxSize := flag.Int("max", 0, "Maximum file size in bytes (required)")
	outputDir := flag.String("dir", "", "Directory to store the generated files (required)")

	flag.Parse()

	// Validate flags
	if *numFiles <= 0 || *minSize < 0 || *maxSize < *minSize || *outputDir == "" {
		fmt.Println("Error: All flags are required and must have valid values.")
		fmt.Println("Usage: go run main.go --num <number> --min <size> --max <size> --dir <directory>")
		os.Exit(1)
	}

	// Ensure output directory exists
	err := os.MkdirAll(*outputDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", *outputDir, err)
		os.Exit(1)
	}

	// Generate files
	err = generateMarkdownFiles(*numFiles, *minSize, *maxSize, *outputDir)
	if err != nil {
		fmt.Printf("Error generating files: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %d Markdown files in %s\n", *numFiles, *outputDir)
}

// generateMarkdownFiles creates Markdown files with random sizes within the specified range.
func generateMarkdownFiles(numFiles, minSize, maxSize int, outputDir string) error {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numFiles; i++ {
		// Random file size
		size := rand.Intn(maxSize-minSize+1) + minSize

		// Generate file content
		content := generateRandomMarkdown(size)

		// File name
		fileName := filepath.Join(outputDir, fmt.Sprintf("file_%d.md", i+1))

		// Write file
		err := os.WriteFile(fileName, []byte(content), 0644)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %v", fileName, err)
		}
	}
	return nil
}

// generateRandomMarkdown creates random Markdown content of a given size with varied elements.
func generateRandomMarkdown(size int) string {
	// Example front matter
	frontMatter := `---
id: %d
title: "Sample Markdown File %d"
tags: [example, test, random]
---\n\n`
	content := fmt.Sprintf(frontMatter, rand.Intn(1000), rand.Intn(1000))

	// Markdown elements for variety
	elements := []string{
		"# Header Level 1\n",
		"## Header Level 2\n",
		"- A bullet list item\n",
		"1. A numbered list item\n",
		"`Inline code example`\n",
		"```\nCode block example\n```\n",
		"A paragraph with **bold** and *italic* text.\n",
		"A [link](https://example.com) in the Markdown content.\n",
		"> A blockquote for demonstration purposes.\n",
		"Regular text line with some random words.\n",
	}

	// Generate random content until the desired size
	for len(content) < size {
		line := elements[rand.Intn(len(elements))]
		content += line
	}

	// Truncate to exact size if necessary
	if len(content) > size {
		content = content[:size]
	}

	return content
}
