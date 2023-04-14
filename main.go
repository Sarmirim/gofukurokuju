package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sarmirim/gofukurokuju/pkg/api"
	"github.com/sarmirim/gofukurokuju/pkg/auth"
	"github.com/sarmirim/gofukurokuju/pkg/config"
)

// var (
// 	AuthConfig *config.AuthConfig
// )

// const (
// 	userAgent = "universal:ENIGMA23:v0.1"
// )

func init() {
	gin.SetMode(gin.ReleaseMode)
	print(config.GetConfig())
	// AuthConfig.GetInstance()
}

func main() {
	r := gin.Default()
	api.Connect(r.Group("/api"))
	go auth.Auth()
	r.Run(":9876") // listen and serve on 0.0.0.0:9876 (for windows "localhost:9876")
}
