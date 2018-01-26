package harverine

import (
	"bufio"
	"io"
	"os"
)

type FileReader struct{}

func NewFileReader() FileReader {
	return FileReader{}
}

func (r FileReader) ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	reader := bufio.NewReader(file)
	for {
		l, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		lines = append(lines, string(l))
	}
	return lines, nil
}
