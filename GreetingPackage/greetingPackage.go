package GreetingPackage

import (
	"fmt"

	"bufio"

	"os"

	"errors"

	"math/rand"

	"time"
)

var randFormats = [3]string{
	"Hi, %v Welcome from GreetingPackage",
	"Nice to meet you, %v",
	"Hail Hydra, %v	",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Greeting(name string) (string, error) {
	message := fmt.Sprintf(randFormats[rand.Intn(len(randFormats))], name)
	if name == "\n" || name == "" {
		return "", errors.New("empty name")
	}
	return message, nil
}

func AskTheName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is you name, sweetheart")
	input, _ := reader.ReadString('\n')
	return input
}

func MultipleGreeting(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Greeting(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
