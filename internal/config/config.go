package config

import (
	"errors"
	"os"

	"github.com/jamesTait-jt/go-aoc/internal/parse"
)

var ROOT_DIR = "../.."
var COOKIE_FILEPATH = ".cookie"

func GetSessionCookie() (string, error) {
	cookieFileContents, err := parse.Lines(os.DirFS(ROOT_DIR), COOKIE_FILEPATH)
	if err != nil {
		return "", err
	}

	if len(cookieFileContents) != 1 {
		return "", errors.New("session cookie file was not valid")
	}

	return cookieFileContents[0], nil
}
