// Package input provides functionality to download and read input data for Advent of Code challenges.
package input

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jamesTait-jt/go-aoc/internal/config"
	"github.com/jamesTait-jt/go-aoc/internal/parse"
)

// fileExtension is used in .gitignore to prevent comitting test data
const fileExtension = "aocin"

// inputURLfmt id the location of the test input with year and day as string formatting arguments
const inputURLfmt = "https://adventofcode.com/%d/day/%d/input"

// Download downloads the input data for specified year and days from the Advent of Code website.
// It saves the downloaded inputs to files in the specified input directory.
// If a file already exists for a year and day, and ForceDownload is not set to true in the AppConfig,
// it does not download the input again.
func Download(appConfig config.AppConfig) error {
	for _, day := range appConfig.Days {
		inputFilePath := fmt.Sprintf("%s/%d/%d.%s", appConfig.Input, appConfig.Year, day, fileExtension)

		if fileExists(inputFilePath) && !appConfig.ForceDownload {
			fmt.Printf("the input for year=%d day=%d already exists (-force to overwrite)\n", appConfig.Year, day)

			continue
		}

		url := fmt.Sprintf(inputURLfmt, appConfig.Year, day)
		if err := downloadDay(appConfig, url, inputFilePath); err != nil {
			return err
		}
	}

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func downloadDay(appConfig config.AppConfig, url, path string) error {
	req, err := prepareRequest(url, appConfig.SessionCookie)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	err = writeToFile(resp.Body, path)
	if err != nil {
		return err
	}

	return nil
}

func prepareRequest(url, sessionCookie string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "github.com/jamesTait-jt/go-aoc")

	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	}
	req.AddCookie(cookie)

	return req, nil
}

func writeToFile(content io.ReadCloser, pathToWrite string) error {
	err := os.MkdirAll(filepath.Dir(pathToWrite), os.ModePerm)
	if err != nil {
		return err
	}

	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(pathToWrite, flags, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return err
	}

	return nil
}

// Read reads input data from a file corresponding to the given year and day.
// It returns a slice of strings, where each string represents a line in the input file.
func Read(year, day int, inputDir string) ([]string, error) {
	inputPath := fmt.Sprintf("%d/%d.%s", year, day, fileExtension)
	input, err := parse.Lines(os.DirFS(inputDir), inputPath)
	if err != nil {
		return []string{}, err
	}

	return input, nil
}
