package main

import "fmt"

func reverse(s []string) []string {
	l := 0
	r := len(s) - 1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	return s
}

func main() {
	str := []string{"B", "R", "Y", "A", "N"}
	fmt.Println(reverse(str))
}
