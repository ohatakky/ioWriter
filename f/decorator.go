package f

import (
	"io"
	"os"
)

// io.MultiWriterはio.Writerをラップしたデコレータ
// デコレータ：オブジェクトをラップして追加の機能を実装すること
func Decorator() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter sample")
}
