package main

import (
	"files/read"
	"fmt"
	"os"
)

func main() {

	// get working directory
	rootPath, _ := os.Getwd()
	fmt.Println(rootPath)

	// Read sample file
	message, err := read.ReadTextFile(rootPath + "/data/sample.txt")

	// handle errors
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(message)

}
