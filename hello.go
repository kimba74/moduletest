package moduletest

import (
	"fmt"
	"strings"
)

type stringer interface {
	String() string
}

// Hello greets the name provided or "world" if name is an empty string
func Hello(name interface{}) (string, bool) {
	greet := "world"
	if name != nil {
		switch s := name.(type) {
		case string:
			if strings.TrimSpace(s) != "" {
				greet = s
			}
		case stringer:
			greet = s.String()
		default:
			return "", false
		}
	}
	return fmt.Sprintf("Hello, %s!", greet), true
}

// HelloWorld says hello to the world
func HelloWorld() string {
	greeting, _ := Hello(nil)
	return greeting
}
