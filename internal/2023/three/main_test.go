package three_test

import (
	"os"
	"testing"

	"github.com/jamesTait-jt/go-aoc/internal/2023/three"
	"github.com/jamesTait-jt/go-aoc/internal/parse"
	"github.com/stretchr/testify/assert"
)

func Test_PartOne(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 4361,
		"real.txt": 536576,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := three.PartOne(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}

func Test_PartTwo(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 467835,
		"real.txt": 75741499,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := three.PartTwo(input)

		// Assert
		assert.Equal(t, expected, solution)
	}
}