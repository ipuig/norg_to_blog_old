package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
    ContentPath string `json:"content_path"`
    Port int `json:"port"`
    Ip string `json:"ip"`
    CertPath string `json:"cert"`
    KeyPath string `json:"key"`
    Author string `json:"author"`
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
        return
    }

    if !validateConfigIp(config.Ip) {
        fmt.Printf("config.json error: %s\ndefaulting to HTTP mode\nreading posts from '%s'\nlistening at port %d\n", "Invalid IP", SiteConfig.ContentPath, SiteConfig.Port)
        return
    }

    SiteConfig = config;
}

func validateConfigIp(ip string) bool {
    if ip == "" {
        return true
    }

    octets := strings.Split(ip, ".")
    if len(octets) != 4 {
        return false
    }

    for _, octet := range octets {
        v, err := strconv.Atoi(octet)
        if err != nil {
            return false;
        }

        if v < 0 || v > 255 {
            return false
        }
    }

    return true
}
