package api

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sarmirim/gofukurokuju/pkg/auth"
	"github.com/sarmirim/gofukurokuju/reddit"
)

// Req - struct to parse json from request
type Req struct {
	URL string `json:"url"`
}

// API -
func post(c *gin.Context) {
	var req Req

	if err := c.ShouldBind(&req); err != nil {
		log.Println(err, "c.BindJSON")
	}

	fmt.Println(req.URL)

	// var answer map[string]interface{} = MapRedditRequestController(req.URL)
	var responseJSON reddit.Post = RedditRequestController(req.URL)

	// if len(responseJSON.Data.Children) == 0 {
	// 	c.JSON(404, gin.H{"message": "Data is empty"})
	// 	return
	// }

	var answer = responseJSON.Data.Children[0]
	// js, err := json.Marshal(answer)

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(js)

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(200, answer.Data)
}

// RedditRequestController -
func RedditRequestController(link string) reddit.Post {
	lastPart := `.json?raw_json=1`
	var oauthLink = strings.Replace(link, "https://www", "https://oauth", 1)
	var answer reddit.Post
	var err error
	var f func()
	var counter int = 0
	f = func() {
		counter += 1
		answer, err = RedditRequest(oauthLink+lastPart, counter)
	}

	f()

	if err.Error() == "status code" {
		log.Println(err)
		auth.Auth()
		f()
	}

	return answer
}

// // MapRedditRequestController -
// func MapRedditRequestController(link string) map[string]interface{} {
// 	client := &http.Client{}

// 	lastPart := `.json?raw_json=1?client_id=fukurokuju&response_type=code&state=RANDOM_STRING&redirect_uri=http://nouri.com/&duration=temporary&scope=read`
// 	req, err := http.NewRequest("GET", link+lastPart, nil)
// 	if err != nil {
// 		log.Println(err, "req")
// 		return nil
// 	}

// 	bearer := fmt.Sprintf("bearer %s", config.GetConfig().Access_token)
// 	req.Header.Set("Authorization", bearer)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println(err, "resp")
// 		return nil
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println(err, "body")
// 		return nil
// 	}
// 	fmt.Println("body", string(body))

// 	jsonMap := make(map[string]interface{})
// 	jsonErr := json.Unmarshal(body, &jsonMap)
// 	if jsonErr != nil {
// 		log.Println(jsonErr, "jsonErr")
// 	}

// 	return jsonMap
// }
