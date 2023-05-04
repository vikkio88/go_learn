package h

import "fmt"

// shorthand syntax for fmt.Sprintf
func F(pattern string, params ...any) string {
	return fmt.Sprintf(pattern, params...)
}
