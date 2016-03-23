package main

import "fmt"

func main(){
	var x [5]int
	var y [5]float64

	z := [5]int{1,2,3,4,5} // Shorter Sintax

	w := [5]int{
	1,
	2,
	3,
	4,
//	5, // this last comma is required by Go, it allows us to easily remove an element from the array
	   // commenting out the line
	}

	x[4] = 100

	y[0] = 99
	y[0] = 93
	y[0] = 77
	y[0] = 82
	y[0] = 83

	var total float64 = 0

//	for i := 0 ; i < 5 ; i++ {
//	for i := 0 ; i < len(y) ; i++ {
//	for i, value := range y { // Go compiler won't allow create variables that never use
	for _, value := range y { // with '_' we tell to the compiler that we don't need this

		total += value
//		total += y[i]
	}
	fmt.Println(total/ float64(len(y)) )

}
