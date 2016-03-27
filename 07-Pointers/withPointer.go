package main

import "fmt"

func main(){
	x := 5
	zero(&x) // '&' operator to find the address of a variable
	fmt.Println(x)
}

func zero(xPointer *int) {
	*xPointer = 0
}
