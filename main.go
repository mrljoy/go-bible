package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"test_chat/book/ch2/lengconv"
)

var pwd string

func main() {
	/*for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\r\n", err)
			continue
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\r\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}*/
	/*if len(os.Args) == 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if input.Text() == "end" {
				fmt.Println("type OK")
				break
			}
			transfer(input.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			transfer(arg)
		}
	}*/
	/*for _, arg := range os.Args[1:] {
		t , _ := strconv.Atoi(arg)
		x := uint64(t)
		pc := popcount.PopCount(x)
		println(pc)
		pc = popcount.PopByClearing(x)
		println(pc)
		pc = popcount.PopByShifting(x)
		println(pc)
	}*/
	/*x := "hello"
	for _, x := range x {
		x := x + 'A' + 'a'
		//fmt.Printf("%c", x)
		fmt.Println(x)
	}*/
	fmt.Println(pwd)
	for x := 0; x < 28; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("not this wd, err : %v", err)
	}
}

func transfer(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mf: %v\r\n", err)
		return
	}
	m := lengconv.Meter(t)
	f := lengconv.Foot(t)
	fmt.Printf("%s = %s, %s = %s\r\n", f, lengconv.FToM(f), m, lengconv.MToF(m))
}
