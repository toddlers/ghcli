package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/toddlers/ghcli/config"
	"github.com/toddlers/ghcli/models"
)

func GetUser(name string) *models.User {
	resp, err := http.Get(config.APIURL + config.UserEndpoint + name)
	if err != nil {
		log.Fatalf("Error receving data: %s\n", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error receving data: %s\n", err)
	}
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalf("Error receving data: %s\n", err)
	}
	return &user
}
