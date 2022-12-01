package util

import (
	"bufio"
	"os"
)

func FileScanner(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file)
}
