package gist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/toddlers/ghcli/config"
	"github.com/toddlers/ghcli/utils"
	"golang.org/x/oauth2"
)

type Client interface {
	DownloadSnippet(string) (*Snippet, error)
	UploadSnippet(string) error
	GetSnippets() ([]*github.Gist, error)
}

type GistClient struct {
	Client *github.Client
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
	accessToken, err := utils.GetGithubAccessToken()
	if err != nil {
		return nil, fmt.Errorf(`access is not provided $%v`, config.GithubAccessToken)
	}
	client := GistClient{
		Client: githubClient(accessToken),
	}
	return client, nil
}

func GetGists(username string) ([]*github.Gist, error) {
	if os.Getenv(config.GithubAccessToken) != "" {
		fmt.Println("Listing for Authenticated user")
		client, err := NewGistClient()
		gists, err := client.GetSnippets()
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to fetch the gists")
		}
		return gists, nil
	}
	url := config.APIURL + config.UserEndpoint + username + "/gists"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed : %s", resp.Status)
	}

	var result []*github.Gist
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

func GistUpload(body string) (err error) {
	client, err := NewGistClient()
	if err != nil {
		return errors.Wrap(err, "Failed to initialize gist client")
	}
	if err = client.UploadSnippet(body); err != nil {
		return errors.Wrap(err, "Failed to upload snippet")
	}
	return nil
}

func GistDownload(id string) (err error) {
	client, err := NewGistClient()
	if err != nil {
		return errors.Wrap(err, "Failed to initialize gist client")
	}
	snippet, err := client.DownloadSnippet(id)
	if err != nil {
		return errors.Wrap(err, "could not able to download the snippet")
	}
	fmt.Println("saving the file")
	f, err := os.Create(snippet.Filename)
	if err != nil {
		return errors.Wrap(err, "Could not able to create the file")
	}
	defer f.Close()
	// Use MultiWriter so we can write to stdout and
	// a file on the same operation
	dest := io.MultiWriter(os.Stdout, f)
	_, err = io.Copy(dest, bytes.NewBufferString(snippet.Content))
	if err != nil {
		return errors.Wrap(err, "Could not able to save the file")
	}
	return nil
}
