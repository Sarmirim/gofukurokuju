package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sarmirim/gofukurokuju/pkg/config"
	"github.com/sarmirim/gofukurokuju/reddit"
)

// RedditRequestController -
func RedditRequest(link string, counter int) (reddit.Post, error) {
	// post := &reddit.Post{}
	post := []reddit.Post{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Println("ERROR req", err)
		return reddit.Post{}, err
	}

	req.Header.Set("User-Agent", "GOFUKUROKUJU")
	bearer := fmt.Sprintf("bearer %s", config.GetConfig().Access_token+"1")
	if counter == 2 {
		bearer = fmt.Sprintf("bearer %s", config.GetConfig().Access_token)
	}

	req.Header.Set("Authorization", bearer)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ERROR resp", err)
		return reddit.Post{}, err
	}

	defer resp.Body.Close()
	// var post interface{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR body", err)
		return reddit.Post{}, err
	}

	fmt.Println("body", string(body))

	if resp.StatusCode == 401 {
		return reddit.Post{}, errors.New("status code")
	}

	jsonErr := json.Unmarshal(body, &post)
	if jsonErr != nil {
		log.Println("ERROR jsonErr", jsonErr)
		return reddit.Post{}, jsonErr
		// return problem
	}
	// var inter interface{}
	// jsonErr = json.Unmarshal(body, &inter)
	// if jsonErr != nil {
	// 	log.Println(jsonErr, "jsonErr")
	// 	// return problem
	// }
	// fmt.Println(post)
	// fmt.Println(inter)

	// return problem
	// answer := post[0].Data.Children[0].Data
	// answer.UTC = time.Unix(int64(answer.Created_utc), 0)
	return post[0], nil
}
