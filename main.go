package main

import (
	"fmt"
	"net/http"
	"os"
)

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb)
	printGreating(sb)

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}

func printGreating(b bot) {
	println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello, World!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
