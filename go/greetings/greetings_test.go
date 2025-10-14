package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Puffy"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Puffy")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello(Puffy) = %q, %v, want for %#q, nil`, msg, err, want)
	}
}

func TestEmptyHello(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
