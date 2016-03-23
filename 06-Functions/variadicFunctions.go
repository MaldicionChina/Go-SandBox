// using '...' before the type name of the last parameter
// is posible indicate that ot takes zero or more of those parameters

package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _,value := range args {
		total += value
	}
	return total
}

func main(){
	fmt.Println(add(1,4,5,2,3))
} 
