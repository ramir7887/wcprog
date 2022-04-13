package wc

import (
	"bufio"
	"io"
	"os"
)

// CountLines return count lines in named file.
func CountLines(name string) (int, error) {
	count := 0
	f, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		count++
	}
	if err := s.Err(); err !=nil {
		return 0, err
	}
	return count, nil
}

// CountWords return count words in named file.
func CountWords(name string) (int, error) {
	count := 0
	f, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			count += lineCountWords(line)
			break
		}
		if err != nil {
			return 0, err
		}
		count += lineCountWords(line)
	}
	return count, nil
}

func lineCountWords(s string) int {
	count := 0
	out := false
	for idx := 0; idx < len(s); idx++ {
		if (s[idx] != ' ' && s[idx] != '\r' && s[idx] != '\n') && !out {
			count++
			out = true
		} else if s[idx] == ' ' || s[idx] == '\r' || s[idx] == '\n' {
			out = false
		}
	}
	return count
}