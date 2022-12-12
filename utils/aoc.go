package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

const URL = "https://adventofcode.com/2022/day/%d/input"

func GetInputScanner(day int) *bufio.Scanner {
	input := GetInput(day)
	return StringScanner(input)
}

func StringScanner(input string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(input))
}

func FetchInput(day int) string {
	cwd, _ := os.Getwd()
	data, err := os.ReadFile(path.Join(cwd, ".env"))
	check(err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(URL, day), nil)
	check(err)
	req.AddCookie(&http.Cookie{Name: "session", Value: strings.Trim(string(data), "\n")})

	fmt.Println("Fetching data...")
	resp, err := client.Do(req)
	check(err)
	fmt.Println("Fetched")

	body, err := io.ReadAll(resp.Body)
	cacheInput(day, body)
	check(err)

	return strings.TrimSpace(string(body))
}

func GetInput(day int) string {
	if existsInCache(day) {
		return getCachedInput(day)
	}
	return FetchInput(day)
}

func cacheInput(day int, data []byte) {
	path := cachePath(day)
	err := os.WriteFile(path, data, 0444)
	check(err)
}

func existsInCache(day int) bool {
	path := cachePath(day)
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	check(err)

	return true
}

func getCachedInput(day int) string {
	data, err := os.ReadFile(cachePath(day))
	check(err)
	return strings.TrimSpace(string(data))
}

func cachePath(day int) string {
	cwd, err := os.Getwd()
	check(err)
	return path.Join(cwd, "inputs", fmt.Sprint(day))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
