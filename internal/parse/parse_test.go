package parse_test

import (
	"testing"

	"github.com/jamesTait-jt/go-aoc/internal/parse"
	"github.com/jamesTait-jt/go-aoc/internal/parse/testdata"
	"github.com/stretchr/testify/assert"
)

func TestLines_GivenValidFile_ThenLinesReturned(t *testing.T) {
	for _, ts := range testdata.GetLinesTestdata_GivenValidFile() {
		t.Run(ts.Name, func(t *testing.T) {
			// Arrange
			// Act
			ls, err := parse.Lines(ts.InputFileSystem, ts.InputPath)

			// Assert
			assert.NoError(t, err, "unexpected error")
			assert.Equal(t, ts.WantLines, ls, "returned lines are not equal to expected")
		})
	}
}

func TestLines_GivenCantParseLines_ThenError(t *testing.T) {
	for _, ts := range testdata.GetLinesTestdata_GivenCantParseLines() {
		t.Run(ts.Name, func(t *testing.T) {
			// Arrange
			// Act
			ls, err := parse.Lines(ts.InputFileSystem, ts.InputPath)

			// Assert
			assert.EqualError(t, err, ts.WantError.Error(), "error mismatch")
			assert.Equal(t, []string{}, ls, "returned lines are not zero value")
		})
	}
}
