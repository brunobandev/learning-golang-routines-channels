package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// This function will run through the links and finish.
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// To check every time we need to create a infite loop.
	// for {
	// 	go checkLink(<-c, c)
	// }

	// another alternative, to be more explicit in the value being passed to the function
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// Adding a time call a func instead of call it every single moment.
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
