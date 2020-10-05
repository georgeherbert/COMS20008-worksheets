package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, j := range slice {
		slice[i] = f(j)
	}
}

func mapArray(f func(a int) int, array [5]int) [5]int {
	for i, j := range array {
		array[i] = f(j)
	}
	return array
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(intsSlice)
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)
	fmt.Println()

	intsArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(intsArray)
	intsArray = mapArray(addOne, intsArray)
	fmt.Println(intsArray)
	fmt.Println()

	newSlice := intsSlice[1:3]
	fmt.Println(newSlice)
	mapSlice(square, newSlice)
	fmt.Println(newSlice)
	fmt.Println(intsSlice)
	fmt.Println()

	fmt.Println(intsSlice)
	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)
	fmt.Println()

}
