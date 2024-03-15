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

const INPUT_DIR_FMT = "%s/internal/%d/.input/real"
const inputFileFormatString = "%s/internal/%d/.input/real/%d.txt"

func Download(year, day int, forceDownload bool) error {
	// If the file already exists, we don't need to download it again
	inputFilePath := fmt.Sprintf(inputFileFormatString, config.ROOT_DIR, year, day)
	_, err := os.Stat(inputFilePath)
	if err == nil && !forceDownload {
		fmt.Printf("the input for year=%d day=%d already exists. use -force if you want to overwrite it\n", year, day)

		return nil
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := prepareRequest(url)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	err = writeToFile(resp.Body, inputFilePath)
	if err != nil {
		return err
	}

	return nil
}

func Read(year, day int) ([]string, error) {
	inputDir := fmt.Sprintf(INPUT_DIR_FMT, config.ROOT_DIR, year)
	inputFileName := fmt.Sprintf("%d.txt", day)
	input, err := parse.Lines(os.DirFS(inputDir), inputFileName)
	if err != nil {
		return []string{}, err
	}

	return input, nil
}

func prepareRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "github.com/jamesTait-jt/go-aoc")

	sessionCookieVal, err := config.GetSessionCookie()
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionCookieVal,
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
