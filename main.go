package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Repos struct {
	Name string `json:"Name"`
	URL  string `json:"url"`
}

func main() {

	userName := os.Args[1]
	res, _ := http.Get("http://www.mocky.io/v2/5b52fd522f0000510d3bb683")

	temp, _ := ioutil.ReadAll(res.Body)

	var repos []Repos
	err := json.Unmarshal(temp, &repos)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	fmt.Printf("User name: %s \n\n", userName)
	for _, repo := range repos {
		fmt.Printf("Name: %s \n", repo.Name)
		fmt.Printf("URL: %s \n\n", repo.URL)
	}
}
