package util

import (
	"bufio"
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
