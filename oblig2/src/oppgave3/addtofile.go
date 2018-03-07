package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func writeToFile() {

	var n1 int
	var n2 int

	fmt.Println("Enter num: ")
	fmt.Scan(&n1)
	fmt.Println("Enter num: ")
	fmt.Scan(&n2)

	file, err := os.Create("3b.txt")
	checkErr(err)

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintln(n1))
	_, err = file.WriteString(fmt.Sprintln(n2))
	checkErr(err)

	err = file.Sync()
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readResult(path string) {
	data, err := ioutil.ReadFile(path)
	checkErr(err)

	tempData := string(data)
	stringData := strings.Split(tempData, "\n")
	temp := stringData[len(stringData)-2]

	res, err := strconv.Atoi(temp)
	checkErr(err)

	fmt.Println("Result from file: ", res)
}
