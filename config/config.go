package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
    ContentPath string `json:"content_path"`
    Port int `json:"port"`
    CertPath string `json:"cert"`
    KeyPath string `json:"key"`
}

var SiteConfig Config = Config{
    ContentPath: "content/posts",
    Port: 8080,
}

func LoadConfig() {
    config := Config{};
    content, err := os.ReadFile("config.json")
    if err != nil {
        fmt.Printf("config.json not found.\ndefaulting to HTTP mode\nreading posts from '%s'\nlistening at port %d\n", SiteConfig.ContentPath, SiteConfig.Port)
        return
    }

    err = json.Unmarshal(content, &config)
    if err != nil {
        fmt.Printf("config.json error: %s\ndefaulting to HTTP mode\nreading posts from '%s'\nlistening at port %d\n", err.Error(), SiteConfig.ContentPath, SiteConfig.Port)
    }

    SiteConfig = config;
}
