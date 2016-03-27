package main

import "fmt"

func main(){
	x := 5
	zero(x)
	fmt.Println(x)
}


// Won't work... this 'x' variable is a local variable... it has not global scope 
func zero(x int) {
	x = 0
}
