// Map is unordered colletion of key-value pairs
package main

import "fmt"

func main(){
//	var x map[string]int // 'x' is a map of 'strings' to 'ints'
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
// We can delete items from map using the built-in 'delete' function
	// delete(x,"key")
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"

	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"]) // returns the zero value for the value type

//	name, ok := elements["Un"]
//	fmt.Println(name,ok)

/*	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
*/
// ------ Shoter Syntax to declare map
/*	elementsShorter := map[string]string {
		"H" : "Hydrogen",
		"He" : "Helium",
		"Li" : "Lithium",
		"Be" : "Berillium",
		"B" : "Boron", // This last comma is mandatory for Go...
	}
*/

// ------ Shorter Syntax to declare map and nested maps
	elementsShorterNestedMap := map[string]map[string]string {
		"H" : map[string]string{
			"name" : "Hydrogen",
			"state" : "gas",
			},
		"He" : map[string]string{
			"name" : "Helium",
			"state" : "gas",
			},
		"Li" : map[string]string{
			"name" : "Litium",
			"state" : "solid",
			},
		"Be" : map[string]string{
			"name" : "Berillium",
			"state" : "solid",
			},
	}

	if el, ok := elementsShorterNestedMap["Li"]; ok {
		fmt.Println("Testins Nested Maps!!!!")
		fmt.Println(el["name"], el["state"])
	}
}
