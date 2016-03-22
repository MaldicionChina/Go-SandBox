package main

import "fmt"

func main(){
	var x string = "hello"
	var y string = "world"
	fmt.Println( x == y ) // Should print 'false' because 'hello' is no the same as 'word'
}
