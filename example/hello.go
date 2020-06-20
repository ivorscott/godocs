// This overview will be displayed under the overview section
// of the example package documentation.
// https://github.com/ivorscott/godocs
package example

import "fmt"

// Hello returns a geeting when provided a name.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
