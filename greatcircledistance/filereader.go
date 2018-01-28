package greatcircledistance

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
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

		if js, ok := isValidJSON(l); ok {
			lines = append(lines, js)
		} else {
			err := errors.New(fmt.Sprintf("Not valid json `%s`", js))
			return nil, err
		}

	}
	return lines, nil
}

func isValidJSON(s []byte) (string, bool) {
	var js map[string]interface{}
	castedStr := string(s)
	return castedStr, json.Unmarshal([]byte(s), &js) == nil
}
