package f

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample\n")
}

// ブラウザへの書き込みに使えるhttp.ResponseWriterも、io.Writerインタフェースの同じWrite()メソッドで書き出すことができる。
func HttpWrite() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
