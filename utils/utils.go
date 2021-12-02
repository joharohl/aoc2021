package utils

import (
	"bufio"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(path string) (lines []string) {
	lines = make([]string, 0, 1000)
	file, err := os.Open(path)
	CheckError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	CheckError(scanner.Err())
	return
}
