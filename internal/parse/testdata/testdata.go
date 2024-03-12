package testdata

import (
	"errors"
	"io/fs"
	"testing/fstest"
)

type LinesTestdata_GivenValidFile struct {
	Name string

	InputFileSystem fs.FS
	InputPath       string

	WantLines []string
}

func GetLinesTestdata_GivenValidFile() []LinesTestdata_GivenValidFile {
	return []LinesTestdata_GivenValidFile{
		{
			Name: "GivenValidTextFile",
			InputFileSystem: fstest.MapFS{
				"file.txt": &fstest.MapFile{
					Data: []byte("Line1\nLine2\nLine3"),
				},
			},
			InputPath: "file.txt",
			WantLines: []string{"Line1", "Line2", "Line3"},
		},
		{
			Name: "GivenValidTextFileInDirectory",
			InputFileSystem: fstest.MapFS{
				"files/file.txt": &fstest.MapFile{
					Data: []byte("Line1\nLine2\nLine3"),
				},
			},
			InputPath: "files/file.txt",
			WantLines: []string{"Line1", "Line2", "Line3"},
		},

	}
}

type LinesTestdata_GivenCantParseLines struct {
	Name string

	InputFileSystem fs.FS
	InputPath       string

	WantError error
}

func GetLinesTestdata_GivenCantParseLines() []LinesTestdata_GivenCantParseLines {
	return []LinesTestdata_GivenCantParseLines{
		{
			Name: "GivenCouldNotOpenFile_ThenErrorOpeningFile",
			InputFileSystem: fstest.MapFS{
				"files/file.txt": &fstest.MapFile{
					Data: []byte("Line1\nLine2\nLine3"),
				},
			},
			InputPath: "file.txt",
			WantError: errors.New("error opening file: open file.txt: file does not exist"),
		},
		{
			Name: "GivenDirectory_ThenErrorScanningLines",
			InputFileSystem: fstest.MapFS{
				"files/file.txt": &fstest.MapFile{
					Data: []byte("Line1\nLine2\nLine3"),
				},
			},
			InputPath: "files",
			WantError: errors.New("error scanning lines (did you pass a directory?): read files: invalid argument"),
		},
	}
}
