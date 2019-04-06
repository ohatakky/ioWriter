package m

import "fmt"

func strByte() {
	str := "string"
	bstr := []byte(str)
	fmt.Println(str)
	fmt.Println(bstr)
	fmt.Println(&bstr[0])
	fmt.Println(len(bstr))
}
