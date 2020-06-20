package example

import "fmt"

// Here's some description text. It will be rendered just before the example when documentation is rendered in go docs. The output result is just a comment. This example also functions as a test and will fail is the result does not match the expected output. Pretty cool.
func ExampleHello() {
	greeting := Hello("Jon")
	fmt.Println(greeting)
	// Output:
	// Hello, Jon
}
