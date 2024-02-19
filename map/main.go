package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#ff1254",
		"blue":  "#ef0012",
	}

	fmt.Println(colours)

	printMap(colours)
}

func printMap(m map[string]string) {
	for colour, hexCode := range m {
		fmt.Println("Hex code for", colour, "is", hexCode)
	}
}
