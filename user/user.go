package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/toddlers/ghcli/config"
	"github.com/toddlers/ghcli/repos"
)

func GetUser(name string) *github.User {
	resp, err := http.Get(config.APIURL + config.UserEndpoint + name)
	if err != nil {
		log.Fatalf("Error receving data: %s\n", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error receving data: %s\n", err)
	}
	var user github.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalf("Error receving data: %s\n", err)
	}
	return &user
}

func GetStarredRepos(username string) (*repos.ReposSearchResult, error) {
	url := config.APIURL + config.UserEndpoint + username + "/starred"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("query failed : %s", resp.Status)
	}

	var result []*github.Repository
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &repos.ReposSearchResult{Items: result}, nil
}

func GetZen() ([]byte, error) {
	url := config.APIURL + "/zen"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("query failed : %s", resp.Status)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "Can't read from the body")
	}
	return bodyBytes, nil
}
