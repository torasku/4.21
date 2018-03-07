package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	getFileInfo(getArg())

}

func getArg() string {
	args := os.Args[1:]
	var st string

	for _, v := range args {
		st += v
	}
	return st
}

func getFileInfo(s string) {
	fi, err := os.Lstat(s)
	if err != nil {
		log.Fatal(err)
	}
	mode := fi.Mode()

	info, err := os.Stat(s)
	if err != nil {
		log.Fatal(err)
	}

	var bytes int64 = info.Size()
	var kb int64 = (bytes / 1024)
	var mg float64 = (float64)(kb / 1024)
	var gb float64 = (mg / 1024)
	var modePerm = mode & os.ModePerm
	var append = mode & os.ModeAppend
	var device = mode & os.ModeDevice
	var symLink = mode & os.ModeSymlink

	fmt.Println("Information about: ", s)
	fmt.Println("Size: ", bytes, "bytes, ", kb, "kb, ", mg, "mg, ", gb, "gb")

	if mode.IsRegular() {
		fmt.Println("Regular file")
	} else {
		fmt.Println("Not a regular file")
	}
	if mode.IsDir() {
		fmt.Println("Is a directory")
	} else {
		fmt.Println("Is not a directory")
	}
	if modePerm == 0777 {
		fmt.Println("Has UNIX permission bits")
	} else {
		fmt.Println("Has not UNIX permission bits")
	}
	if append != 0 {
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}
	if device != 0 {
		fmt.Println("Is device file")
	} else {
		fmt.Println("Is not a device file")
	}
	if symLink != 0 {
		fmt.Println("Is a symlink")
	} else {
		fmt.Println("Is not a symlink")
	}
}
