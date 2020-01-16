package hello_test

import (
	"testing"

	"abc.xyz/hello"
)

func TestHello(t *testing.T) {

	x := hello.Hello("abc")
	if "Hello, abc" != x {
		t.Errorf("want %s, but %s", "Hello, abc", x)
	}

	x = hello.Hello("")
	if "Who are you?" != x {
		t.Errorf("want %s, but %s", "Who are you?", x)
	}
}
