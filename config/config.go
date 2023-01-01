package config

import "os"

type Config struct {
	SecretKey string
}

func New() *Config {
	secret := os.Getenv("secret_key")
	config := &Config{SecretKey: secret}
	return config
}

func SetSecretKeyEnv(secret string) {
	os.Setenv("secret_key", secret)
}
