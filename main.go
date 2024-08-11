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
    log.Print(config.SiteConfig)
    addr := fmt.Sprintf(":%d", config.SiteConfig.Port)

    log.Printf("Server running at port %d\n", config.SiteConfig.Port)
    if config.SiteConfig.CertPath == "" && config.SiteConfig.KeyPath == "" {
        http.ListenAndServe(addr, routes.Routes())
    } else {
        http.ListenAndServeTLS(addr, config.SiteConfig.CertPath, config.SiteConfig.KeyPath, routes.Routes())
    }

}

