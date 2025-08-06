package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type englishBot struct{}
type spanishBot struct{}

type logWritter struct{}

func (logWritter) Write(p []byte) (n int, err error) {
	fmt.Println((string(p)))
	return len(p), nil
}

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
	io.Copy(logWritter{}, resp.Body)
	// bs := make([]byte, 99999)
	// n, respReadError := resp.Body.Read(bs)

	// if respReadError != nil {
	// 	fmt.Println("Error reading response body:", respReadError)
	// 	os.Exit(1)
	// }
	// fmt.Printf("Readed %v bytes, result is %v", n, string(bs))
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
