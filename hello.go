package moduletest

import (
	"fmt"
	"strings"
)

// Hello greets the name provided or "world" if name is an empty string
func Hello(name string) string {
	greet := "world"
	if str := strings.TrimSpace(name); str != "" {
		greet = str
	}
	return fmt.Sprintf("Hello, %s!", greet)
}
