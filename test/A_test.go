//package main
//
//import (
//	"fmt"
//	"os"
//	"strings"
//)
//
//func main() {
//	var s, sep string
//	for i := 1; i < len(os.Args); i++ {
//		s += sep + os.Args[i]
//		sep = " "
//	}
//	fmt.Println(s)
//	s1, sep1 := "", ""
//	for _, arg := range os.Args[1:] {
//		s1 += sep1 + arg
//		sep1 = " "
//	}
//	fmt.Println(s1)
//	fmt.Println(strings.Join(os.Args[1:], " "))
//}

package test

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkA1(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	}
}

func BenchmarkA2(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		s1, sep1 := "", ""
		for _, arg := range os.Args[1:] {
			s1 += sep1 + arg
			sep1 = " "
		}
		fmt.Println(s1)
	}
}

func BenchmarkA3(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}
}