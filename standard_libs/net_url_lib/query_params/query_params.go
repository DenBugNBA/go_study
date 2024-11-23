package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL := "https://example.com/api/resource"

	params := url.Values{}
	params.Add("param1", "value1")
	params.Add("param2", "value2")

	fullURL := baseURL + "?" + params.Encode()
	fmt.Println("fullUrl: ", fullURL)
}
