// Package config provides functionality for initializing and managing configuration parameters for Advent of Code challenges.
package config

import (
	"errors"
	"flag"
	"os"
	"path/filepath"

	"github.com/jamesTait-jt/go-aoc/internal/parse"
)

// AppConfig represents the configuration parameters for running Advent of Code challenges.
type AppConfig struct {
	// Year of the Advent of Code challenge.
	Year int

	// Days of the Advent of Code challenge.
	Days []int

	// ForceDownload specifies whether to force download the input data even if it already exists.
	ForceDownload bool

	// SessionCookie is the session cookie used to log in to the Advent of Code account.
	SessionCookie string

	// Input is the directory where input files will be saved.
	Input string
}

// Init initializes the application configuration based on command line arguments.
// It parses command line flags, validates them, and returns an AppConfig.
func Init() (AppConfig, error) {
	var year *int
	var days IntSliceValue
	var forceDownload *bool
	var sessionCookiePath *string
	var input *string

	year = flag.Int("year", -1, "the year you would like to run")
	flag.Var(&days, "days", "the days you would like to run")
	forceDownload = flag.Bool("force-download", false, "whether you would like to download the test input even if it already exists on your system (please be mindful of spamming the aoc servers!)")
	sessionCookiePath = flag.String("session-cookie", "", "the file path of your session cookie to log in to your aoc account")
	input = flag.String("input", "", "the directory that the input files will be saved to")
	flag.Parse()

	config := AppConfig{}

	// Set the year
	if *year == -1 {
		return AppConfig{}, errors.New("you must specify a year with -year")
	}
	config.Year = *year

	// Set the days
	parsedDays, err := parseDays(days)
	if err != nil {
		return AppConfig{}, err
	}
	config.Days = parsedDays

	// Set the session cookie
	sessionCookie, err := parseSessionCookie(*sessionCookiePath)
	if err != nil {
		return AppConfig{}, err
	}
	config.SessionCookie = sessionCookie

	// Set the directory for the challenge input
	if *input == "" {
		return AppConfig{}, errors.New("you must provide a directory to save the input files to")
	}
	config.Input = *input

	// Set remaining config values
	config.ForceDownload = *forceDownload

	return config, nil
}

func parseDays(days IntSliceValue) ([]int, error) {
	if len(days) == 0 {
		days = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	}

	for _, day := range days {
		if day > 25 {
			return []int{}, errors.New("no day can be greater than 25")
		}
	}

	return days, nil
}

func parseSessionCookie(path string) (string, error) {
	if path == "" {
		return "", errors.New("you must provide the path of the file containing your session cookie")
	}

	dir := filepath.Dir(path)
	filename := filepath.Base(path)

	cookieFileContents, err := parse.Lines(os.DirFS(dir), filename)
	if err != nil {
		return "", err
	}

	if len(cookieFileContents) != 1 {
		return "", errors.New("the session cookie file was not valid")
	}

	return cookieFileContents[0], nil
}
