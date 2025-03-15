package main

import (
	"fmt"
	"net/http"
	"time"
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

	start := time.Now()
	for _, link := range links {
		go checkLink(link, c)

	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	// fmt.Println(<-c)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Time elapsed: ", elapsed.Milliseconds(), "ms")
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	c <- link
	fmt.Println(link, "is up!")
}
