package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"sample": "data",
	}

	gw := gzip.NewWriter(w)
	defer gw.Flush()
	mw := io.MultiWriter(gw, os.Stdout)
	je := json.NewEncoder(mw)
	je.SetIndent("", "	")
	je.Encode(source)
}

func main() {
	// f.OutputFile()
	// f.StdoutWrite()
	// f.BufferWrite()
	// f.NetWrite()
	// f.HttpWrite()
	// f.Decorator()
	// f.GzipWrite()
	// f.JsonEncoderWrite()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
