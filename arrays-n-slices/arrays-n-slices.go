// arrays-n-slices example

package main

import "fmt"

const gCap int = 5          // total number of array elements
var gLen int                // interger number of items in our array currently
var gGroceries [gCap]string //array of strings

func addGrocery(a string) {
	if gLen < gCap {
		gGroceries[gLen] = a
		gLen++
	} else {
		fmt.Println("Error!!! we've reached the upper bound of the array!!!")
	}
}

func listGroceries() {
	fmt.Println("Grocery list is as follows:")
	for i := 0; i < gLen; i++ {
		fmt.Println(gGroceries[i])
	}
}

func main() {
	listGroceries()
	addGrocery("Coffee")
	addGrocery("Milk")
	addGrocery("French Fries")
	addGrocery("Pizza")
	listGroceries()
	addGrocery("Celery")

}
