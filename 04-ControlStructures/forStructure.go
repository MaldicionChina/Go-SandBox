package main

import "fmt"

func main(){
	i := 1
	fmt.Println("One way to do a 'for'")
	for i <= 10{
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("another way to do a 'for'")
	for i := 11; i <= 20; i++{
		fmt.Println(i)
	}
}
