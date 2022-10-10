package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	for _, website := range websites {
		checkwebsite(website)
	}
}

func checkwebsite(website string) {
	_, err := http.Get(website)
	if err != nil {
		fmt.Println(website, "might be down!")
		return
	}
	fmt.Println(website, "is up!")
}
