package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i, j := range slice {
		slice[i] = f(j)
	}
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	for i, j := range array {
		array[i] = f(j)
	}
	return array
}

func main() {
	intsSlice := []int{1, 2, 3}
	fmt.Println(intsSlice)
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	fmt.Println()

	intsArray := [3]int{1, 2, 3}
	fmt.Println(intsArray)
	intsArray = mapArray(addOne, intsArray)
	fmt.Println(intsArray)
}
