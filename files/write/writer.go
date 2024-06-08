package write

import "os"

func WriteToFile(filename string, content string) error {

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}
