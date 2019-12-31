package gist

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/toddlers/ghcli/config"
	"github.com/toddlers/ghcli/models"
)

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
