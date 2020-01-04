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

type File struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Language string `json:"language"`
	RawURL   string `json:"raw_url"`
	Size     int    `json:"size"`
}
type Gist struct {
	URL         string    `json:"url"`
	ForksURL    string    `json:"forks_url"`
	CommitsURL  string    `json:"commits_url"`
	GitPullURL  string    `json:"git_pull_url"`
	GitPushURL  string    `json:"git_push_url"`
	HTMLURL     string    `json:"html_url"`
	Public      bool      `json:"public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Comments    int       `json:"comments"`
	CommentsURL string    `json:"comments_url"`
	Owner       *Owner
	Files       map[string]File
}

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
	githubAccessToken = "GITHUB_ACCESS_TOKEN"
)

type Client interface {
	//	GetGist() (*Snippet, error)
	UploadSnippet(string) error
}

// Snippet is the remote snippet
type Snippet struct {
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
	accessToken, err := getGithubAccessToken()
	if err != nil {
		return nil, fmt.Errorf(`access is not provided $%v`, githubAccessToken)
	}
	client := GistClient{
		Client: githubClient(accessToken),
		ID:     config.Gc.GistID,
	}
	return client, nil
}

func getGithubAccessToken() (string, error) {
	if config.Gc.AccessToken != "" {
		return config.Gc.AccessToken, nil
	} else if os.Getenv(githubAccessToken) != "" {
		return os.Getenv(githubAccessToken), nil
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
	fmt.Println(g.ID)
	if g.ID == "" {
		gistID, err := g.CreateGist(context.Background(), gist)
		if err != nil {
			return err
		}
		fmt.Printf("Gist ID: %s\n", *gistID)
	}
	return nil
}

func (g GistClient) CreateGist(ctx context.Context, gist *github.Gist) (gistID *string, err error) {
	fmt.Println("Create Gist")
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
	s.Suffix = "Createing Gist..."
	defer s.Stop()
	retGist, _, err := g.Client.Gists.Create(ctx, gist)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create gist")
	}
	return retGist.ID, nil
}
