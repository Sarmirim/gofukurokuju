package config

import (
	"fmt"
	"sync"
)

type Oauth2Config struct {
	Access_token string
}

var (
	authConfig *Oauth2Config
	once       sync.Once
)

func GetConfig() *Oauth2Config {
	once.Do(func() {
		authConfig = &Oauth2Config{}
	})
	return authConfig
}

func SetAccessToken(newToken string) {
	if len(newToken) > 0 {
		authConfig.Access_token = newToken
		fmt.Printf("New token is %s\n", authConfig.Access_token)
	} else {
		fmt.Println("Token size is zero")
	}
}
