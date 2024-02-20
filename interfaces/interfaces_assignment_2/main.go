package main

import (
	"io"
	"log"
	"os"
)

func main() {

	arg := os.Args

	file, err := os.Open(arg[1])
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, file)
}
