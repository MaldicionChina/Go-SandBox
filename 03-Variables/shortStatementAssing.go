// Creating a new varaible with a STARTING VALUE is a so common
// Go also support a shorter statement
// ':=' 
package main

import "fmt"

func main(){
	x := "Hello World" // notice that no type was specified, Go compiler will detect type
	var y string = "Hello World"
	n := 5 // The same thins works for other types

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(n)
}

