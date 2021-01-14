package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma3(arg))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	rs := ""
	for i := 0; i < n; i++ {
		if i%3 == 0 && i != 0 {
			rs = string(s[n-1-i]) + "," + rs
		} else {
			rs = string(s[n-1-i]) + rs
		}
	}
	return rs
}

func comma3(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		//if i%3 == 0 && i != 0 {
		//	buf.WriteString(",")
		//}
		//buf.WriteString(string(s[n-1-i]))
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(string(s[i]))
	}
	return buf.String()
}
