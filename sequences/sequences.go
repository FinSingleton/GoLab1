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
	for i, x := range slice {
		slice[i] = f(x)
	}

}

func mapArray(f func(a int) int, array *[3]int) {
	for i, x := range array {
		array[i] = f(x)
	}
}

func main() {

	intsSlice := []int{1, 2, 3}
	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	intsArray := [3]int{1, 2, 3}
	mapArray(addOne, &intsArray)
	fmt.Println(intsArray)
	/*
		newSlice := intsSlice[1:3]
		mapSlice(square, newSlice)
		fmt.Println(newSlice)
		fmt.Println(intsSlice)
	*/
	double(intsSlice)
	fmt.Println(intsSlice)
}
