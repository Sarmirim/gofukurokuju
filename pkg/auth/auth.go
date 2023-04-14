package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sarmirim/gofukurokuju/pkg/config"
)

type AuthResponse struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

func getEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Printf("%s not set\n", key)
	} else {
		fmt.Printf("%s=%s\n", key, val)
	}
	return val
}

func Auth() {
	client := &http.Client{}
	url := `https://www.reddit.com/api/v1/access_token?grant_type=client_credentials&device_id=423rwer2asdas1231eqe`
	authResponse := &AuthResponse{}

	req, err := http.NewRequest(http.MethodPost, url, http.NoBody)
	if err != nil {
		log.Println(err, "req")
		// return reddit.Data{}
	}
	req.Header.Set("User-Agent", "UNIVERSAL:ENIGMA23:V0.01TELEBOT")

	var username = getEnv("username")
	var password = getEnv("password")

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err, "resp")
	}

	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err, "body")
	// }

	json.NewDecoder(resp.Body).Decode(authResponse)
	// config.SetAccessToken(string)
	fmt.Println("Auth response", authResponse)
	config.SetAccessToken(authResponse.Access_token)
}
