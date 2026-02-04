package greetings

import (
	"regexp"
	"testing"
)

// Test if the message contains the name
func TestHelloName(t *testing.T) {
	name := "Arthur"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Arthur")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Arthur") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
