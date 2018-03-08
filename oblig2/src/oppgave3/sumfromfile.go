package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) (s string) {

	var res int
	var temp string
	s = ""

	data, err := ioutil.ReadFile(path)
	checkErr(err)

	go func() {
		checkSigint()
	}()

	tempData := string(data)
	stringData := strings.Split(tempData, "\n")
	stringData = stringData[:len(stringData)-1]

	for _, v := range stringData {
		temp = string(v)
		tempInt, err := strconv.Atoi(temp)
		checkErr(err)
		res += tempInt
	}
	file, err2 := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	checkErr(err2)

	defer file.Close()

	_, err2 = file.WriteString(fmt.Sprintln(res))
	checkErr(err2)

	err2 = file.Sync()

	return s
}
