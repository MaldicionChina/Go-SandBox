// What is the lenght od the slice created using make([]int,3,9) ... ¿11?

package main

import "fmt"

func main(){
	sli := make([]int,3,9)
	fmt.Println(len(sli))
}
