// # sigo-parser

// sigo-parser will take a template and parse some text and extract the parts as indicated by the template.
// BUT: for now it will take some text (output of method calls from phpstorm) and using the amount of tab spacing
// at the beginning of a line output a different format.
// The solution is inspired by the outline2 program from the Go Programming Language book.
// `go get gopl.io/ch5/outline2`

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]
	text := readLines(path)

}

// readLines reads a whole file into memory
// and returns a slice of its lines.
// http://stackoverflow.com/questions/5884154/golang-read-text-file-into-string-array-and-write
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
