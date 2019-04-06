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
	// https://qiita.com/awakia/items/ef8978fc33ffb1d27460
	// HandleFuncすると、DefaultServerMuxと言うものにマッピングが登録される
	// muxはマルチプレクサのこと
	// 複数のリクエストがあって、一致した1つを出力とするルーティングは、マルチプレクサの動作に似ている？？？
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
