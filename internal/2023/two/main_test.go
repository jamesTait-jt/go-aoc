package two_test

import (
	"os"
	"testing"

	"github.com/jamesTait-jt/go-aoc/internal/2023/two"
	"github.com/jamesTait-jt/go-aoc/internal/parse"
	"github.com/stretchr/testify/assert"
)

func Test_PartOne(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 8,
		"real.txt": 2085,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := two.PartOne(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}

func Test_PartTwo(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 2286,
		"real.txt": 71274,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := two.PartTwo(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}

