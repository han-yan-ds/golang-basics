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

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	i := 0
	for msg := range ch { // this syntax looks at the channel and gets messages, assigns to msg.  It's nonstop
		fmt.Println(msg)
		i++
		if i == len(links) { // this syntax is non-stop which is why we need this
			return
		}
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		channel <- fmt.Sprint(link, " might be down :(")
		return
	}
	channel <- fmt.Sprint(link, " is up :)")
}
