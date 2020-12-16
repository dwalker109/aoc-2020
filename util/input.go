package util

import (
	"bufio"
	"bytes"
	"os"
)

func StreamInput(i chan<- string, fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic("No file!")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		i <- s.Text()
	}

	close(i)

	if s.Err() != nil {
		panic("File parse error!")
	}
}

func StreamInputCustomSplit(i chan<- string, fn string, sf func(data []byte, atEOF bool) (advance int, token []byte, err error)) {
	f, err := os.Open(fn)
	if err != nil {
		panic("No file!")
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(sf)
	for s.Scan() {
		i <- s.Text()
	}

	close(i)

	if s.Err() != nil {
		panic("File parse error!")
	}
}

func SplitOnString(s string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	var (
		searchBytes = []byte(s)
		searchLen   = len(searchBytes)
	)

	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLen, data[0:i], nil
		}

		if atEOF {
			return dataLen, data, nil
		}

		// Moar data
		return 0, nil, nil
	}
}
