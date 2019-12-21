package config

import (
	"os"
)

type EmbyConfig struct {
	EmbyUrl string 
	EmbyToken string 
	EmbyUserName string
}

type Config struct {
	Emby EmbyConfig
}

func New() *Config  {
	return &Config {
		Emby: EmbyConfig{
			EmbyUrl: getEnv("EMBY_URL", ""),
			EmbyToken: getEnv("EMBY_TOKEN", ""),
			EmbyUserName: getEnv("EMBY_USER", ""),
		},
	}

}

func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
	return value
    }

    return defaultVal
}