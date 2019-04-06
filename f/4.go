package f

import "os"

// fmt.Println()では、最終的にos.StdoutのWrite()を呼び出している。
// 下記はそれと等価なコード
func StdoutWrite() {
	os.Stdout.Write([]byte("os.Stdout sample\n"))
}
