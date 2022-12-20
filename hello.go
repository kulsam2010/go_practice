package main

import (
	"log"
	"net/http"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(name, language)
}

func greetingPrefix(name string, language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix + name
}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet("world", w)
}

func main() {
	// fmt.Println(Hello("Sameer", ""))
	// fmt.Println(integers.Add(5, 6))
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetingHandler)))
}
