package greetings

import (
	"testing"

	"regexp"
)

// to run tests: go test
// to run tests with verbose output: go test -v

// for info on the T param: https://pkg.go.dev/testing#T
func TestHelloName(t *testing.T) {
	name := "daniel"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("daniel")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("daniel") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}