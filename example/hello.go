package example

import "fmt"

// Hello returns a geeting when provided a name.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
