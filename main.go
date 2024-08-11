package main

import (
	"blog/config"
	"blog/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
    config.LoadConfig();
    addr := fmt.Sprintf("%s:%d", config.SiteConfig.Ip, config.SiteConfig.Port)

    log.Printf("Server running at %s\n", addr)
    if config.SiteConfig.CertPath != "" && config.SiteConfig.KeyPath != "" {
        err := http.ListenAndServeTLS(addr, config.SiteConfig.CertPath, config.SiteConfig.KeyPath, routes.Routes())
        if err != nil {
            panic(err);
        }
    } else {
        http.ListenAndServe(addr, routes.Routes())
    }

}

