package main

import (
	"fmt"
)

func main() {
	list := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(list[1])

	newList := BubbleSort(list)
	fmt.Println(newList)

	newList2 := Bubble_sort_modified(list)
	fmt.Println(newList2)
}

func Bubble_sort_modified(list []int) []int {
	// Skriv din kode her
	n := len(list)
	swapped := true
	for swapped {
		swapped = false
		for index := 1; index < n-1; index++ {
			if list[index-1] > list[index] {
				temp := list[index-1]
				list[index-1] = list[index]
				list[index] = temp
				swapped = true
			}
		}
	}
	return list

	/*n := len(list)
	var swapped bool
	for swapped == false {
		for i := 1; i < n-1; i++ {
			if list[i-1] > list[i] {
				temp := list[i-1]
				list[i-1] = list[i]
				list[i] = temp
				swapped = true
			}
		}
	}
	return list*/
}

func BubbleSort(list []int) []int {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
	return list
}
