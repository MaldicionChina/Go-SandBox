// Problem 4 of chapter 4
package main

import "fmt"

func main(){
	var faherenheit float64
	var celsius float64
	fmt.Println("Program to convert Faherenheit to Celsius")
	fmt.Println("Enter temperature: ")
	fmt.Scanf("%f", &faherenheit)
	celsius = (((faherenheit - 32)*5)/9)
	fmt.Printf("Celsius: %.5f \n", celsius)
}
