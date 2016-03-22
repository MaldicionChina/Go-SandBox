package main

import "fmt"

func main(){
	i := 1
	fmt.Println("One way to do a 'for'")
	for i <= 10{
// ------------- If Stament
		if i % 2 == 0 {
			fmt.Println("Even --")
// ------------- Else Stament
		} else {
			fmt.Println("Odd --")
		}
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("another way to do a 'for'")
	for i := 11; i <= 20; i++{
//---------------- If Stament
		if i % 2 == 0 {
			fmt.Println("Divisible by 2")
//---------------- Else if Stament
		} else if i % 3 == 0{
			fmt.Println("Divisible by 3")
		} else {
			fmt.Println("Isn't divisible by 3 neither by 2")
		}
		fmt.Println(i)
	}
}
