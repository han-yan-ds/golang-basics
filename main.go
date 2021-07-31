package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	numDone := 0

	ch := make(chan int)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for { // keep running until all channels have received a value
		numDone = numDone + <-ch
		if numDone == len(links) {
			return
		}
	}

}

func checkLink(link string, channel chan int) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down :(")
		channel <- 1
		return
	}
	fmt.Println(link, "is up :)")
	channel <- 1
}
