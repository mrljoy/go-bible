package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var n = flag.Bool("n", true, "")

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	fmt.Println(x % y)
	fmt.Println(gcd(x, y))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func gcd2(x, y int) int {

}
