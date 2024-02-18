package main

import "fmt"

type intSlices []int

func main() {
	is := intSlices{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	is.printMessage()
}

func (i intSlices) printMessage() {
	for _, num := range i {
		if num%2 == 0 {
			fmt.Println(num, "is Even")
		} else {
			fmt.Println(num, "is Odd")
		}
	}
}
