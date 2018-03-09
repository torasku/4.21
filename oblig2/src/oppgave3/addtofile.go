package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func writeToFile() (s string) {

	// scanner input, lager fil og skriver til fil
	var n1 int
	var n2 int
	s = ""

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
	return s
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readResult(path string) (s int) {

	// leser fra fil, lagrer siste tall i resultat og konverterer til int
	data, err := ioutil.ReadFile(path)
	checkErr(err)

	stringArr := strings.Split(string(data), "\n")
	temp := stringArr[len(stringArr)-2]

	res, err := strconv.Atoi(temp)
	checkErr(err)
	s = res
	return s
}
