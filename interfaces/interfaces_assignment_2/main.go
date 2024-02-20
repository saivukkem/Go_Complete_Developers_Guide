package main

import (
	"io"
	"log"
	"os"
)

/*
1. Load file from hd
2. Read its contents
3. Print them to the terminal
*/

func main() {

	arg := os.Args

	file, err := os.Open(arg[1])
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, file)
}
