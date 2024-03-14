package four_test

import (
	"os"
	"testing"

	"github.com/jamesTait-jt/go-aoc/internal/parse"
	"github.com/jamesTait-jt/go-aoc/internal/2023/four"
	"github.com/stretchr/testify/assert"
)

func Test_PartOne(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 13,
		"real.txt": 24706,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := four.PartOne(input)

		// Assert
		assert.Equal(t, expected, solution)
	}	
}

func Test_PartTwo(t *testing.T) {
	testCases := map[string]int{
		"mini.txt": 30,
		"real.txt": 13114317,
	}

	for filePath, expected := range testCases {
		// Arrange
		input, err := parse.Lines(os.DirFS("./testdata"), filePath)
		if err != nil {
			t.Fatal(err)
		}

		// Act
		solution := four.PartTwo(input)

		// Assert
		assert.Equal(t, expected, solution)
	}	
}