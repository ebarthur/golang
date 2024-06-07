package main

import (
	// alias the packages (r,w for read and write)
	r "files/read"
	w "files/write"
	"fmt"
	"os"
)

func main() {

	// get working directory
	rootPath, _ := os.Getwd()
	fmt.Println(rootPath)

	// Read sample file
	message, err := r.ReadTextFile(rootPath + "/data/sample.txt")

	// handle errors
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(message)

	// write the file

	content := "\n Well, it's another day in paradise!"

	newContent := fmt.Sprint(message, "\n", content)

	err = w.WriteToFile(rootPath+"/data/sample.txt", newContent)

	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	fmt.Println("File written successfully!")

}
