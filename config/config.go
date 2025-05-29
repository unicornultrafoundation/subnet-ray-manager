package config

import (
	"os"
)

// Config holds all configuration for the application.
type Config struct {
	EthRpcUrl       string
	ContractAddress string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() *Config {
	return &Config{
		EthRpcUrl:       os.Getenv("ETH_RPC_URL"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
	}
}
