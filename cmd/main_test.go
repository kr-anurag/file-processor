package main

import (
	"file-processor/internal/processor"
	"file-processor/internal/service"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessFile(t *testing.T) {
	// dir := t.TempDir()
	dir := "./../sample-files/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	files := []struct {
		name    string
		content string
	}{
		{"file1.txt", "Hello World\nThis is a test file."},
		{"file2.txt", "Another file\nWith multiple lines\nAnd words."},
		{"empty.txt", ""},
		{"whitespace.txt", "   \n\t\n"},
		{"special_chars.txt", "!@#$%^&*()_+\n{}|:\"<>?`~"},
		{"large.txt", strings.Repeat("word ", 10000) + "\n" + strings.Repeat("line\n", 50000)},
	}

	for _, file := range files {
		err := os.WriteFile(filepath.Join(dir, file.name), []byte(file.content), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", file.name, err)
		}
	}

	tests := []struct {
		name     string
		fileName string
		lines    int
		words    int
		hasError bool
	}{
		{"Success Read File 1", "file1.txt", 2, 7, false},
		{"Success Read File 2", "file2.txt", 3, 7, false},
		{"Success Empty Text", "empty.txt", 0, 0, false},
		{"Success Whitespace Text", "whitespace.txt", 2, 0, false},
		{"Success Special Characters", "special_chars.txt", 2, 2, false},
		{"Success Large File", "large.txt", 50001, 60000, false},
		// {"Error Missing Text", "missing.txt", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := filepath.Join(dir, tt.fileName)

			textProcessor := &processor.TextFileProcessor{}
			processorService := &service.FileProcessorService{Processor: textProcessor}
			summaries := processorService.ProcessFiles([]string{file})

			for _, result := range summaries {
				assert.Equal(t, tt.lines, result.Lines)
				assert.Equal(t, tt.words, result.Words)
				assert.Equal(t, tt.hasError, result.Error != nil)
			}
		})
	}
}
