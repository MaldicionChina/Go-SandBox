package main

import "fmt"

func main(){
	i := 1
	var englishNumber string = ""
	for i <= 10{
// ------------ Switch Stament
		switch i {
			case 1: englishNumber = "One"
			case 2: englishNumber = "Two"
			case 3: englishNumber = "Three"
			case 4: englishNumber = "Four"
			case 5: englishNumber = "Five"
			case 6: englishNumber = "Six"
			default : englishNumber = "Unknown Number"
		}
		fmt.Println(englishNumber)
		i = i + 1
	}
}
