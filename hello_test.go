package moduletest

import (
	"fmt"
	"testing"
)

func TestHelloEmpty(t *testing.T) {
	want := "Hello, world!"
	if got, _ := Hello(""); got != want {
		t.Errorf(fmt.Sprintf("Hello(\"\")=%q; want %q", got, want))
	}
}

func TestHelloInvalid(t *testing.T) {
	want := "Hello, world!"
	if got, _ := Hello(23); got != want {
		t.Errorf(fmt.Sprintf("Hello(\"\")=%q; want %q", got, want))
	}
}

func TestHelloNil(t *testing.T) {
	want := "Hello, world!"
	if got, _ := Hello(nil); got != want {
		t.Errorf(fmt.Sprintf("Hello(\"\")=%q; want %q", got, want))
	}
}

func TestHelloWorld(t *testing.T) {
	want := "Hello, world!"
	if got := HelloWorld(); got != want {
		t.Errorf(fmt.Sprintf("Hello(\"\")=%q; want %q", got, want))
	}
}

func TestHelloYou(t *testing.T) {
	want := "Hello, you!"
	if got, _ := Hello("you"); got != want {
		t.Errorf(fmt.Sprintf("Hello(\"you\")=%q; want %q", got, want))
	}
}
