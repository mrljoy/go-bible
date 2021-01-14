package main

import (
	"fmt"
)

func main() {
	s := [...]int{1, 5, 7, 3, 2, 8, 9, 6, 3, 5}
	reverse(&s)
	fmt.Println(s)
}

func reverse(s *[10]int) {
	for i, j := 0, 9; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

