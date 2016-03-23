package main

import "fmt"

func main(){
/*	panic("PANIC") // function to cause a run time error
	str := recover() // stops the panic and returns the value that was passed to the call to panic
	fmt.Println(str)
	// recover will never happen because the call to panic immediately stops execution of the function
*/

	defer func(){
		str := recover()
		fmt.Println(str)
	} () //

	panic("PANIC")
}
