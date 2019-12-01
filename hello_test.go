package moduletest

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	want := "Hello, world!"
	if got := HelloWorld(); got != want {
		t.Errorf(fmt.Sprintf("Hello(\"\")=%q; want %q", got, want))
	}
}

func TestHelloYou(t *testing.T) {
	want := "Hello, you!"
	if got := Hello("you"); got != want {
		t.Errorf(fmt.Sprintf("Hello(\"you\")=%q; want %q", got, want))
	}
}
