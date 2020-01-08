package utils

import (
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/google/go-github/github"
	"github.com/toddlers/ghcli/config"
)

func DaysAgo(t github.Timestamp) int {
	return int(time.Since(t.Time).Hours() / 24)
}

func Scan(message string) (string, error) {
	tempFile := "/tmp/tempgist.tmp"
	l, err := readline.NewEx(&readline.Config{
		Prompt:            message + "\033[31m»\033[0m ",
		HistoryFile:       tempFile,
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
	})

	if err != nil {
		return "", err
	}

	defer l.Close()

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		return line, nil
	}
	return "", errors.New("Canceled")
}

func GetGithubAccessToken() (string, error) {
	if config.Gc.AccessToken != "" {
		return config.Gc.AccessToken, nil
	} else if os.Getenv(config.GithubAccessToken) != "" {
		return os.Getenv(config.GithubAccessToken), nil
	}
	return "", errors.New("Github AccessToken not found")
}
