package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

type logWriter struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	//htttp request
	res, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(res)

	// bs byte slice k input hisebe pathaise, function oitake update koire dise, so, bs er value change hoise
	// so, bs er value print korle, updated value print hobe
	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	ls := logWriter{}
	io.Copy(ls, res.Body)

	//using io pkg
	// body, err := io.ReadAll(res.Body)
	// res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%s", body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
