package main

import "fmt"

func main() {
//	var x []float64 // The diference with Array is the missing length between the brackets
//	y := make([]float64,5)
//
//	w := []float64{1,2,3,4,5}
//	z := w[0:5] // Creating slice with [low : high] expression

// Slice Functions... append 
	slice1 := []int{1,2,3,4}
	slice2 := append(slice1,5,6)
	fmt.Println(slice1,slice2)
// Slice Functions... copy
	slice3 := []int{1,2,3}
	slice4 := make([]int,2)
	copy(slice4,slice3)
	fmt.Println(slice3,slice4)
}

