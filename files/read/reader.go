package read

import (
	"fmt"
	"os"
)

// function to read the content of a text file
func ReadTextFile(filename string) (string, error) {

	// read the content of the file using 'os' library
	content, err := os.ReadFile(filename)

	// handle errors
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	// return the content of the file
	return string(content), err
}
