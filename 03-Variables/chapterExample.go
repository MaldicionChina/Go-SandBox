package main

import "fmt"

func main(){
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	ouput := input * 2
	fmt.Println(ouput)
}
