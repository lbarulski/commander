package github

import (
	"net/http"
	"commander/libs/config"
	"bufio"
	"commander/libs/userProviders"
)

const PROVIDER_GITHUB = "github"

type github struct {
	userProviders.Provider

	config *config.Config
}

func Initialize(config *config.Config) *github {
	return &github {
		config: config,
	}
}

func (gh *github) GetUsers() []*userProviders.User {
	var usersList []*userProviders.User

	for _, u := range gh.config.GetACLByProvider(PROVIDER_GITHUB) {
		keys := getKeys(u.Username)
		if len(u.PublicKeys) > 0 {
			for _, k := range u.PublicKeys {
				keys = append(keys, k)
			}
		}

		usersList = append(usersList, &userProviders.User{
			Username: u.Username,
			PublicKeys: keys,
		})
	}

	return usersList
}

func getKeys(username string) []string {
	var keys []string
	rsp, err := http.Get("https://github.com/" + username + ".keys")
	if err != nil {
		panic(err.Error())
	}
	if rsp.StatusCode == http.StatusOK {
		scanner := bufio.NewScanner(rsp.Body)

		for scanner.Scan() {
			keys = append(keys, scanner.Text())
		}
	}

	return keys
}