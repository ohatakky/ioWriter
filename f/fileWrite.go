package f

import "os"

// OutputFile is func
func OutputFile() {
	file, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File sample\n"))
	file.Close()
}
