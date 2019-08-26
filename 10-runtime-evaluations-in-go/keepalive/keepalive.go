package main

import (
	"fmt"
	"runtime"
	"syscall"
)

func main() {
	type File struct{ d int }
	d, err := syscall.Open("./testfile.txt", syscall.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
	}
	p := &File{d}
	runtime.SetFinalizer(p, func(p *File) { syscall.Close(p.d) })
	var buf [10]byte
	n, err := syscall.Read(p.d, buf[:])
	if err != nil {
		fmt.Println(err)
	}
	runtime.KeepAlive(p)
	fmt.Println(n)
	fmt.Println(p)
}
