package moduletest

import (
	"fmt"
	"testing"
)

func TestHelloEmpty(t *testing.T) {
	want := "Hello, world!"
	if got, _ := Hello(""); got != want {
		t.Errorf(fmt.Sprintf("Hello(\"\")=%q,true; want %q,true", got, want))
	}
}

func TestHelloInvalid(t *testing.T) {
	if got, ok := Hello(23); ok {
		t.Errorf(fmt.Sprintf("Hello(23)=%q,%t; want \"\",false", got, ok))
	}
}

func TestHelloNil(t *testing.T) {
	want := "Hello, world!"
	if got, _ := Hello(nil); got != want {
		t.Errorf(fmt.Sprintf("Hello(nil)=%q; want %q", got, want))
	}
}

func TestHelloWorld(t *testing.T) {
	want := "Hello, world!"
	if got := HelloWorld(); got != want {
		t.Errorf(fmt.Sprintf("HelloWorld()=%q; want %q", got, want))
	}
}

func TestHelloYou(t *testing.T) {
	want := "Hello, you!"
	if got, _ := Hello("you"); got != want {
		t.Errorf(fmt.Sprintf("Hello(\"you\")=%q; want %q", got, want))
	}
}
