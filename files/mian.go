package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func getFileInfo(path string) {
	fifo, err := os.Stat(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	fattr := fifo.Sys().(*syscall.Win32FileAttributeData)
	fmt.Print(time.Unix(fattr.CreationTime.Nanoseconds()/1000000000, 0))
}

func main() {
	getFileInfo("F:/test.yml")
}
