// age.go
// example if statement for age.go

package main

import "fmt"

func main() {

	age := 13

	if age < 13 {
		fmt.Println("wow you are very young!")
	} else if age < 20 {
		fmt.Println("wow you are a teenager")
	} else if age < 30 {
		fmt.Println("You are in your 20's")
	} else if age < 40 {
		fmt.Println(" You're in your thirties..")
	} else if age < 50 {
		fmt.Println("You are are getting there..")
	} else if age >= 50 {
		fmt.Println("Over the hill")
	}
}
