package five_test

import (
	"os"
	"testing"

	"github.com/jamesTait-jt/go-aoc/internal/2023/five"
	"github.com/jamesTait-jt/go-aoc/internal/parse"
	"github.com/stretchr/testify/assert"
)

func Test_PartOne(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 35,
		"real.txt": 51752125,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := five.PartOne(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}

func Test_PartTwo(t *testing.T) {
	testCases := map[string]int{
		// "mini.txt": 46,
		"real.txt": 13114317,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := five.PartTwo(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}
