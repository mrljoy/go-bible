package main

import (
	"flag"
	"fmt"
	"test_chat/go-bible/ch7/tempconv"
)

//var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
var f = tempconv.FahrenheitFlag("f", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*f)
}
