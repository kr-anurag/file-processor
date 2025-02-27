package processor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextFileProcessor struct{}

func (fp *TextFileProcessor) Process(filename string) FileSummary {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return FileSummary{FileName: filename, Lines: 0, Words: 0, Error: err}
	}
	defer file.Close()

	var lines, words int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		words += len(strings.Fields(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return FileSummary{FileName: filename, Lines: 0, Words: 0, Error: err}
	}

	return FileSummary{FileName: filename, Lines: lines, Words: words}
}
