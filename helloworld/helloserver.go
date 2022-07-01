package main

import (
	"fmt"

	"rsc.io/quote"

	"github.com/GreetingPackage"

	"log"

	"strings"
)

func main() {
	log.SetPrefix("Greeting:")
	log.SetFlags(0)

	fmt.Println("Hello server")
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())

	name := GreetingPackage.AskTheName()
	message, err := GreetingPackage.Greeting(name)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message)
	}

	namesInput := GreetingPackage.AskTheName()
	names := strings.Fields(namesInput)
	greetings, merr := GreetingPackage.MultipleGreeting(names)
	if merr != nil {
		log.Fatal(merr)
	} else {
		fmt.Println(greetings)
	}

}
