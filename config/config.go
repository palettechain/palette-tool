package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ipfs/go-log"
)

type Config struct {
	Rpc          string
	Admin        *Account
	TestAccounts []*Account
}

type Account struct {
	KeyFile    string
	Passphrase string
}

func GenerateConfig(dir string) *Config {
	cfg := &Config{}

	if err := log.SetLogLevel("*", "DEBUG"); err != nil {
		panic(fmt.Sprintf("failed to initialize logger: [%v]", err))
	}

	if _, err := toml.DecodeFile(dir, cfg); err != nil {
		panic(fmt.Sprintf("failed to decode config file: [%v]", err))
	}

	return cfg
}
