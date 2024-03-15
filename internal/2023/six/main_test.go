package six_test

import (
	"os"
	"testing"

	"github.com/jamesTait-jt/go-aoc/internal/2023/six"
	"github.com/jamesTait-jt/go-aoc/internal/parse"
	"github.com/stretchr/testify/assert"
)

func Test_PartOne(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 288,
		"real.txt": 505494,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := six.PartOne(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}

func Test_PartTwo(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 71503,
		"real.txt": 23632299,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := six.PartTwo(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}
