package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	s1, sep1 := "", ""
	for _, arg := range os.Args[1:] {
		s1 += sep1 + arg
		sep1 = " "
	}
	fmt.Println(s1)
	fmt.Println(strings.Join(os.Args[1:], " "))
}