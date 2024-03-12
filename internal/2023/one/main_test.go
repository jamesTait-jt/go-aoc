package one_test

import (
	"os"
	"testing"

	"github.com/jamesTait-jt/go-aoc/internal/2023/one"
	"github.com/jamesTait-jt/go-aoc/internal/parse"
	"github.com/stretchr/testify/assert"
)

func Test_PartOne(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 142,
		"real.txt": 56042,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata/one"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := one.PartOne(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}

func Test_PartTwo(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 281,
		"real.txt": 55358,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata/two"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := one.PartTwo(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}
