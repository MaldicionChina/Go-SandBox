//to put a variable outside of any function, make its scope global 
//and accesible for all function in the program
package main

import "fmt"

var x string = "Hello World" // This variable can be accesible for other functions, not just for the main function

func main(){
//	var x string
	fmt.Println(x)
}

func f(){
	fmt.Println("From function f: "+x)
}
