package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func main() {
	formData := url.Values{
		"name":     {"hello"},
		"surename": {"golang post form"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result["form"])
}
