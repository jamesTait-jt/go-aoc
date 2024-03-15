// parse provides some functions for parsing input from files
package parse

import (
	"bufio"
	"fmt"
	"io/fs"
	"strconv"
	"strings"
)

// Lines returns a slice of strings, each element containing a single line from a file
func Lines(fsys fs.FS, path string) ([]string, error) {
	f, err := fsys.Open(path)
	if err != nil {
		return []string{}, fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
		return []string{}, fmt.Errorf("error scanning lines (did you pass a directory?): %w", err)
    }

	return lines, nil
}

// Nums parses a string representation of a list of numbers, to a slice of ints
func Nums(s string, delimiter string) []int {
	strs := strings.Split(s, delimiter)

	ns := []int{}
	for _, str := range strs {
		if str == "" {
			continue
		}
		
		n, _ := strconv.Atoi(str)
		ns = append(ns, n)
	}

	return ns
}