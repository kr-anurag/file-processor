package processor

type FileProcessor interface {
	Process(filePath string) FileSummary
}

type FileSummary struct {
	FileName string
	Lines    int
	Words    int
	Error    error
}
