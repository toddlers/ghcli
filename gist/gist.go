package gist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/toddlers/ghcli/config"
	"github.com/toddlers/ghcli/models"
)

type Client interface {
	GetGist() (*Snippet, error)
	UploadSnippet(string) error
}

// Snippet is the remote snippet
type Snippet struct {
	Content   string
	UpdatedAt time.Time
}

func GetGists(username string) ([]*models.Gist, error) {
	//https://api.github.com/users/toddlers/gists
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

	var result []*models.Gist
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

func GistUpload(client Client) (err error) {
	//	var snippets models.Snippets
	body := "this is test snippet"
	if err = client.UploadSnippet(body); err != nil {
		return errors.Wrap(err, "Failed to upload snippet")
	}

	gist := &github.Gist{
		Description: github.String("description"),
		Public: false,
		Files: map["temp.go"]github.GistFile{
			github.GistFilename("temp.go"): github.GistFile{
				Content: body,
			},
		},
	}
	fmt.Println("Upload success")
	return nil
}
