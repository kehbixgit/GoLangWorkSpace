// arrays-n-slices example

package main

import "fmt"

var gGroceries []string // slice of strings

func addGrocery(a string) {
	fmt.Println("Capacity:", cap(gGroceries))
	gGroceries = append(gGroceries, a)
}

/*func listGroceries() {
	fmt.Println("Grocery list is as follows:")
	for i := 0; i < len(gGroceries); i++ {
		fmt.Println(gGroceries[i])
	}
}*/

func listGroceries() {
	for i, d := range gGroceries {
		fmt.Println(i, d)
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
