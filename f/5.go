package f

import (
	"bytes"
	"fmt"
)

// bufferもio.Writerを実装している
func BufferWrite() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer sample"))
	fmt.Println(buffer.String())
}
