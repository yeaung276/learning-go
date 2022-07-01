package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("hello"))
	fmt.Println(stringutil.ToUpper("all are written in lower case"))
}
