// 12 days of Christmas example for switch

package main

import "fmt"

func main() {
	fmt.Println("On the", 1, "of Christmas my true love sent to me")
	fmt.Println("a partridge in a pear tree")
	fmt.Println("")
	for day := 2; day <= 12; day++ {
		fmt.Println("On the", day, "of christmas my true love sent to me")
		switch day {
		case 12:
			fmt.Println("12 Drummers Drumming")
			fallthrough
		case 11:
			fmt.Println("Eleven Pipers piping")
			fallthrough
		case 10:
			fmt.Println("Ten Lords a leaping")
			fallthrough
		case 9:
			fmt.Println("Nine ladies Dancing")
			fallthrough
		case 8:
			fmt.Println("Eight Maids a milking")
			fallthrough
		case 7:
			fmt.Println("Seven swans a swimming")
			fallthrough
		case 6:
			fmt.Println("Six geese are laying")
			fallthrough
		case 5:
			fmt.Println("Five Golden rings")
			fallthrough
		case 4:
			fmt.Println(" Four calling birds")
			fallthrough
		case 3:
			fmt.Println("Three french Hens")
			fallthrough
		case 2:
			fmt.Println("Two turtle doves")
			fallthrough
		case 1:
			fmt.Println("And a patridge in a pear tree")
		}
		fmt.Println("")
	}
}
