package f

import "syscall"

// f.fdがファイルディスクリプタに該当

func (f *File) write(b []byte) (n int, err error) {
	for {
		bcap := b
		if needsMaxRW && len(bcap) > maxRW {
			bcap = bcap[:maxRW]
		}
		m, err := fixCount(syscall.Write(f.fd, bcap))
		n += m
		//.
		//.
	}
}
