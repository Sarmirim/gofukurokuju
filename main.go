package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Sarmirim/gofukurokuju/reddit"
)

// Req - struct to parse json from request
type Req struct {
	URL string
}

// var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/api", API)
	http.ListenAndServe(":9876", nil)
}

// Hello -
func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
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

	w.Write([]byte("Hello"))
}

// API -
func API(w http.ResponseWriter, r *http.Request) {
	req := Req{}

	if r.URL.Path != "/api" {
		http.NotFound(w, r)
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
		log.Fatal(jsonErr)
	}

	var ans reddit.Data = MyRequest(req.URL)
	js, err := json.Marshal(ans)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
	// w.Write(reqBody)
	print(string(reqBody))
}

// MyRequest -
func MyRequest(link string) reddit.Data {
	post := []reddit.Post{}
	client := &http.Client{}
	lastPart := ".json?raw_json=1"
	req, err := http.NewRequest("GET", link+lastPart, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "GO Enigma")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	jsonErr := json.Unmarshal(body, &post)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	answer := post[0].Data.Children[0].Data
	return answer
	// log.Println(post[0].Data.Children[0].Data.URL)
}
