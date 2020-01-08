package models

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/toddlers/ghcli/config"
	"golang.org/x/oauth2"
)

type Snippets struct {
	Snippets []SnippetInfo
}

type SnippetInfo struct {
	Description string
	Command     string
	Tag         []string
	Output      string
}

const (
	GithubAccessToken = "GITHUB_ACCESS_TOKEN"
)

type Client interface {
	DownloadSnippet(string) (*Snippet, error)
	UploadSnippet(string) error
	GetSnippets() ([]*github.Gist, error)
}

// Snippet is the remote snippet
type Snippet struct {
	Filename  string
	Content   string
	UpdatedAt time.Time
}

type GistClient struct {
	Client *github.Client
	ID     string
}

func githubClient(accessToken string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	return client

}

func NewGistClient() (Client, error) {
	accessToken, err := GetGithubAccessToken()
	if err != nil {
		return nil, fmt.Errorf(`access is not provided $%v`, GithubAccessToken)
	}
	client := GistClient{
		Client: githubClient(accessToken),
		ID:     config.Gc.GistID,
	}
	return client, nil
}

func GetGithubAccessToken() (string, error) {
	if config.Gc.AccessToken != "" {
		return config.Gc.AccessToken, nil
	} else if os.Getenv(GithubAccessToken) != "" {
		return os.Getenv(GithubAccessToken), nil
	}
	return "", errors.New("Github AccessToken not found")
}

func (g GistClient) UploadSnippet(content string) error {
	gist := &github.Gist{
		Description: github.String("description"),
		Public:      github.Bool(config.Gc.Public),
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(config.Gc.FileName): github.GistFile{
				Content: github.String(content),
			},
		},
	}
	if g.ID == "" {
		gistID, err := g.createGist(context.Background(), gist)
		if err != nil {
			return err
		}
		fmt.Printf("Gist ID: %s\n", *gistID)
	}
	return nil
}

func (g GistClient) DownloadSnippet(id string) (*Snippet, error) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
	s.Suffix = "Downloading Gist..."
	defer s.Stop()
	gist, res, err := g.Client.Gists.Get(context.Background(), id)
	if err != nil {
		if res.StatusCode == 404 {
			return nil, errors.Wrapf(err, "No gist ID (%s)", id)
		}
		return nil, errors.Wrapf(err, "Failed to get the gist")
	}

	var snippet Snippet
	for _, file := range gist.Files {
		snippet.Filename = *file.Filename
		snippet.Content = *file.Content
	}

	if snippet.Content == "" {
		return nil, fmt.Errorf("gist id %s is empty", id)
	}
	snippet.UpdatedAt = *gist.UpdatedAt

	return &snippet, nil
}

func (g GistClient) GetSnippets() ([]*github.Gist, error) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
	s.Suffix = "Getting starred Gists..."
	defer s.Stop()
	gists, _, err := g.Client.Gists.ListStarred(context.Background(), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get the gists")
	}
	return gists, nil

}

func (g GistClient) createGist(ctx context.Context, gist *github.Gist) (gistID *string, err error) {
	fmt.Println("Create Gist")
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
	s.Suffix = "Creating Gist..."
	defer s.Stop()
	retGist, _, err := g.Client.Gists.Create(ctx, gist)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create gist")
	}
	return retGist.ID, nil
}
