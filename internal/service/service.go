package service

import (
	"sync"

	"file-processor/internal/processor"
)

type FileProcessorService struct {
	Processor processor.FileProcessor
}

func (c *FileProcessorService) ProcessFiles(files []string) []processor.FileSummary {
	var wg sync.WaitGroup

	results := make(chan processor.FileSummary, len(files))

	for _, file := range files {
		wg.Add(1)

		go func(file string) {
			defer wg.Done()

			p := &processor.TextFileProcessor{}
			results <- p.Process(file)
		}(file)
	}

	wg.Wait()
	close(results)

	var summaries []processor.FileSummary
	for result := range results {
		summaries = append(summaries, result)
	}

	return summaries
}
