package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, website := range websites {
		go checkwebsite(website, c)
	}

	for l := range c {
		go func(website string) {
			time.Sleep(5 * time.Second)
			checkwebsite(website, c)
		}(l)
	}

}

func checkwebsite(website string, c chan string) {
	_, err := http.Get(website)
	if err != nil {
		fmt.Println(website, "might be down!")
		c <- website
		return
	}
	fmt.Println(website, "is up!")
	c <- website
}
