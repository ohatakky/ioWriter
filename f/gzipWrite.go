package f

import (
	"compress/gzip"
	"io"
	"os"
)

// 書き込まれたデータをgzip圧縮して、os.Fileに中継する
func GzipWrite() {
	file, err := os.Create("gzip-sample.txt.zip")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "gzip-sample.txt"
	io.WriteString(writer, "gzip.Writer sample\n")
	writer.Close()
}
