// Go stament 'defer' schedules a FUNCTION call to be run after
// function completes

package main 

import "fmt"

func second(){
	fmt.Println("2nd")
}

func first(){
	fmt.Println("1st")
}

func main(){
	defer second()
	first()
	// Will print 1st and after 2nd
}
