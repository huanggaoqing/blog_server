package tools

import (
	"bufio"
	"io"
)

func ReadAllByLine(r io.Reader) (string, error) {
	read := bufio.NewReader(r)
	content := ""
	for {
		line, err := read.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		content += line
	}
	return content, nil
}
