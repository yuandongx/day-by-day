package main

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"syscall"
	"time"
)

type Image struct {
	name  string
	ipath string
	year  int
	month int
	day   int
}

var DestPath string = "F:\\tmp"

func (i Image) move(dstpath string) {
	dst, err := os.Create(dstpath)
	if err != nil {
		return
	}
	defer dst.Close()
	src, err := os.Open(i.ipath)
	if err != nil {
		return
	}
	defer src.Close()
	io.Copy(dst, src)
}
func isImg(name string) bool {
	var stuffix [6]string = [6]string{".jpg", "jpeg", ".png", ".btmp", ".mov", ".gif"}
	var tmp string = strings.ToLower(name)
	for _, v := range stuffix {
		if strings.HasSuffix(tmp, v) {
			return true
		}
	}
	return false
}

func readDir(dirName string) {
	infos, errs := ioutil.ReadDir(dirName)
	if errs == nil {
		for _, info := range infos {
			if info.IsDir() {
				readDir(path.Join(dirName, info.Name()))
			} else if isImg(info.Name()) {
				t := info.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds()
				date := time.Unix(t/1000000000, 0)
				img := Image{name: info.Name(), ipath: path.Join(dirName, info.Name()), year: date.Year(), month: int(date.Month()), day: date.Day()}
				go img.move(DestPath)
			}
		}
	}
}
func main() {
	readDir("F:\\备份\\动图")
}
