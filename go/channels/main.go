package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.golang.org",
		"https://www.microsoft.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}

	for l := range c {
		go checkLink(l, c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	c <- link
	fmt.Println(link, "is up!")
}
