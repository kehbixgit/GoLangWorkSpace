//My New Year countdown clock in GO!!!
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 10
	for i > 0 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i = i - 1
	}
	fmt.Println("Happy new years 2021!")
}
