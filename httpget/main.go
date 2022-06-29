package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	httpMethod := "GET"
	//url := "https://api.github.com"
	url := "http://localhost:8080"

	client := http.Client{} // default http client

	request, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/xml")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}
