package main

import ("fmt";"math")

type Circle struct {
	x float64
	y float64
	r float64
//	x , y , r float64
}

// Makes a copy of 'c' varaible
func circleArea(c Circle) float64 {
        return math.Pi * c.r*c.r
}

// Takes 'c' memory address
func circleAreaPtr(c *Circle) float64 {
        return math.Pi * c.r*c.r
}

func main(){
	// Initialization of Circle's Struct
	// create local Cricle varaible that is by default set to zero
	// For Struct means each of the fields is set to their corresponding
	// zero value 
	// var c Circle
	// c := new(Circle)

	// c := Circle{x: 0,y: 0,r: 5}
	c := Circle{0,0,5}

	// We can access fields using the . operator
	// fmt.Println(c.x,c.y,c.r)
	// c.x = 10
	// c.y = 5

	fmt.Println(circleArea(c))
	fmt.Println(circleAreaPtr(&c))

}
