package m

import "fmt"

// Slice is func
func Slice() {
	s := make([]byte, 3)
	s = s[:2]
	fmt.Println(s)
}
