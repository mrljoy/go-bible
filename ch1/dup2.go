package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//从输入中获取文件名
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
			for _, n := range counts {
				if n > 1 {
					fmt.Printf("%s\n", arg)
					break
				}
			}
			counts = map[string]int{}
		}
	}
	//for line, n := range counts {
	//	if n > 1 {
	//		fmt.Printf("%d\t%s\n", n, line)
	//	}
	//}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
