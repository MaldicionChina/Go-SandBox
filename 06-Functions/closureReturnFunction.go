// Closure could be used as a return of an function

package main

import "fmt"

// Function thats retunr a closure function
func makeEvenGenerator() func() uint {
	i := uint(0)
	// closure to be returned
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}

func main(){
	//Assing to a variable the function
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
	// Persist local variable i because the variable nextEven store the makeEvenGenerator
}

