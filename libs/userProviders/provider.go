package userProviders

import "commander/libs/config"

type Provider interface {
	Initialize(config *config.Config) *Provider
	GetUsers() []*User
}

