package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	// oppgave 2
	fileCount(getArg())
	arr := runeCount(getArg())
	arr2 := runeCount(getArg())

	// manuell run --> kommenter ut neste tre linjene under
	/*fileCount("text.txt")
	arr := runeCount("text.txt")
	arr2 := runeCount("text.txt")*/
	var sortedCount []int = Bubble_sort_modified(arr)
	var slice []int = sortedCount[123:]
	getASCII(arr2, slice)

}

func getArg() string {
	args := os.Args[1:]
	var st string

	for _, v := range args {
		st += v
	}
	return st
}

func fileCount(s string) {

	fmt.Println("Information about: ", s)
	fmt.Println("Number of lines in file: ", lineCount(s))

}

func lineCount(s string) int {
	fi, _ := os.Open(s)
	fiScanner := bufio.NewScanner(fi)
	count := 0

	for fiScanner.Scan() {
		count++
	}
	return count
}

func runeCount(s string) []int {
	// leser fil, teller opp antall forekomster for hver verdi
	// og legger dem til i array og returnerer array
	var countArr []int
	var count int
	data, err := ioutil.ReadFile(s)

	for i := 0; i < 128; i++ {
		count = 0
		for _, v := range data {
			if err != nil {
				panic(err)
			}
			if int(v) == i {
				count++
			}
		}
		countArr = append(countArr, count)
	}
	return countArr
}

// used for testing
func checkChar() {
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}
	count := strings.Count(string(data), "A")
	count2 := strings.Count(string(data), "1")
	fmt.Println(count, " ", count2)
}

func Bubble_sort_modified(list []int) []int {
	// bruker bubleSort til aa sortere array med verdier slik
	// at de siste 5 slots i array er topp 5
	n := len(list)
	swapped := true
	for swapped {
		swapped = false
		for index := 1; index < n; index++ {
			if list[index-1] > list[index] {
				temp := list[index]
				list[index] = list[index-1]
				list[index-1] = temp
				swapped = true
			}
		}
		n -= 1
	}
	return list
}

func getASCII(data []int, sortedCount []int) {
	// finner hexa verdi ved aa hente indeks fra hovedarray der
	// topp 5 verdiene er like (indeks brukes til aa vite hvilke ASCII verdi
	// som er knyttet til hver verdi)
	var hexa []int
	for i := 0; i < 5; i++ {
		for j := 0; j < 128; j++ {
			if sortedCount[i] == data[j] {
				hexa = append(hexa, j)
			}
		}
	}
	fmt.Println("Most common runes: ")
	var place int = 1
	for i := 4; i >= 0; i-- {
		if hexa[i] == 32 {
			fmt.Println(place, ". Rune: space", ", ", "Counts: ", sortedCount[i])
			i -= 1
			place += 1
		}
		fmt.Println(place, ". Rune: ", string(hexa[i]), ", ", "Counts: ", sortedCount[i])
		place += 1
	}
}
