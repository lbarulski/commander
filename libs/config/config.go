package config

import (
	"github.com/spf13/viper"
	"fmt"
)

type Config struct {
	ACL []*access `yaml:"acl"`
	SupportedNamespaces []string `mapstructure:"supported_namespaces"` // all if empty
}

type access struct {
	Provider string `yaml:"provider"`
	Username string `yaml:"username"`
	Namespaces []string `mapstructure:"namespaces"`
	Commands []string `mapstructure:"commands"`
	PublicKeys []string `mapstructure:"public_keys"`
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/commander")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	c := &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		panic(err.Error())
	}

	return c
}

func (config *Config) GetACLByProvider(provider string) []*access {
	var acl []*access
	for _, a := range config.ACL {
		if a.Provider == provider {
			acl = append(acl, a)
		}
	}

	return acl
}