package custom

import (
	"commander/libs/config"
	"commander/libs/userProviders"
)

const PROVIDER_CUSTOM = "custom"

type custom struct {
	userProviders.Provider

	config *config.Config
}

func Initialize(config *config.Config) *custom {
	return &custom{
		config: config,
	}
}

func (gh *custom) GetUsers() []*userProviders.User {
	var usersList []*userProviders.User

	for _, u := range gh.config.GetACLByProvider(PROVIDER_CUSTOM) {
		usersList = append(usersList, &userProviders.User{
			Username: u.Username,
			PublicKeys: u.PublicKeys,
		})
	}

	return usersList
}
