package m

import "fmt"

// StrByte is func
func StrByte() {
	str := "string"
	bstr := []byte(str)
	fmt.Println(str)
	fmt.Println(bstr)
	fmt.Println(&bstr[0])
	fmt.Println(len(bstr))
}
