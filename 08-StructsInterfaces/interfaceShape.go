package main

import ( "fmt"; "math" )



type Shape interface {
	// Method set
	// A method set is a list of methods that a type must have
	// in order to "implement" the interface
	area() float64
	perimeter() float64 // remove it a look what happend...
	// another method over here...
	//...
	// method
}

type Circle struct {
        x float64
        y float64
        r float64
//      x , y , r float64
}

func (c *Circle) area() float64 {
        return math.Pi * c.r*c.r
}

func (c *Circle) perimeter() float64{
	return 2 * math.Pi * c.r
}
// Sums perimeters of all shapes that are passed by arguments
func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _,p := range shapes {
		perimeter += p.perimeter()
	}
	return perimeter
}

// Sums areas of all shapes that are passed by arguments
func totalArea(shapes ... Shape) float64 {
	var area float64
	for _,s := range shapes {
		area += s.area()
	}
	return area
}

func main (){
	c := Circle{0,0,5}
	fmt.Println(totalArea(&c))
	fmt.Println(totalPerimeter(&c))
}
