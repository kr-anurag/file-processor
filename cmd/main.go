package main

import (
	"file-processor/internal/processor"
	"file-processor/internal/service"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify the directory containing .txt files")
		return
	}

	dir := os.Args[1]
	files, err := filepath.Glob(filepath.Join(dir, "*.txt"))
	if err != nil || len(files) == 0 {
		fmt.Println("No text files found or error accessing directory.")
		return
	}

	textProcessor := &processor.TextFileProcessor{}
	processorService := &service.FileProcessorService{Processor: textProcessor}
	summaries := processorService.ProcessFiles(files)

	var totalFiles, totalLines, totalWords int
	fmt.Printf("-----------------------------------------\n")
	fmt.Printf("-----------------------------------------\n")
	for _, result := range summaries {
		fmt.Printf("File: %s - Lines: %d, Words: %d\n", result.FileName, result.Lines, result.Words)
		totalFiles++
		totalLines += result.Lines
		totalWords += result.Words
	}
	fmt.Printf("-----------------------------------------\n")
	fmt.Printf("-----------------------------------------\n")
	fmt.Printf("Total Files Processed: %d\n", totalFiles)
	fmt.Printf("Total Lines: %d, Total Words: %d\n", totalLines, totalWords)
}
