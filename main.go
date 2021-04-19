package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sarmirim/gofukurokuju/reddit"
)

// Port - server port
var Port rune = 9876

// Req - struct to parse json from request
type Req struct {
	URL string
}

// var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/api", API)
	http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}

// FirstPart -
var FirstPart string = "To use gofukurokuju please use /api"

// SecondPart -
var SecondPart string = `
Get request with body like: 
{
"url": "https://www.reddit.com/r/memes/comments/ltkhxe/well/"
}`

// Hello -
func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		w.Write([]byte(FirstPart + SecondPart))
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write(reqBody)
	}

	switch r.Method {
	case "POST":
		// w.Write([]byte("Received a POST request\n"))
		w.Write([]byte(`Use only GET request`))
	}

	for k, v := range r.URL.Query() {
		fmt.Printf("%s: %s\n", k, v)
	}

	w.Write([]byte(FirstPart + SecondPart))
}

// API -
func API(w http.ResponseWriter, r *http.Request) {
	req := Req{}

	if r.URL.Path != "/api" {
		w.Write([]byte(FirstPart + SecondPart))
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write(reqBody)
	}

	switch r.Method {
	case "POST":
		w.Write([]byte("Received a POST request\n"))
	}
	for k, v := range r.URL.Query() {
		fmt.Printf("%s: %s\n", k, v)
	}
	jsonErr := json.Unmarshal(reqBody, &req)
	if jsonErr != nil {
		js, err := json.Marshal(reddit.Data{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
		return
	}

	var ans reddit.Data = MyRequest(req.URL)
	js, err := json.Marshal(ans)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(js)
	// w.Write(reqBody)
	print(string(reqBody))
}

// MyRequest -
func MyRequest(link string) reddit.Data {
	problem := reddit.Data{}
	post := []reddit.Post{}
	client := &http.Client{}
	lastPart := ".json?raw_json=1"
	req, err := http.NewRequest("GET", link+lastPart, nil)
	if err != nil {
		log.Fatalln(err)
		return reddit.Data{}
	}

	req.Header.Set("User-Agent", "GO Enigma23")

	resp, err := client.Do(req)
	if err != nil {
		return problem
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return problem
	}

	jsonErr := json.Unmarshal(body, &post)
	if jsonErr != nil {
		return problem
	}

	answer := post[0].Data.Children[0].Data
	// answer.UTC = time.Unix(int64(answer.Created_utc), 0)
	return *answer
}
