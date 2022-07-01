package GreetingPackage

import (
	"testing"

	"regexp"
)

func TestGreeting(t *testing.T) {
	name := "test name"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greeting(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGreetingEmptyString(t *testing.T) {
	msg, err := Greeting("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
