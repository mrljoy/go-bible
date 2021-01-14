package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"net/http"
	"strconv"
	l "test_chat/book/ch1/lissajous"
)

var palette1 = []color.Color{color.RGBA{255, 255, 255, math.MaxUint8}, color.Black, color.RGBA{0, 0, 128, math.MaxUint8}, color.RGBA{30, 144, 255, math.MaxUint8}, color.RGBA{72, 61, 139, math.MaxUint8}, color.RGBA{255, 0, 255, math.MaxUint8}, color.RGBA{124, 252, 0, math.MaxUint8}}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.ParseFloat(r.FormValue("cycles"), 64)
		if err != nil {
			fmt.Fprintf(w, "cycles is request as a parmer, err:%s", err)
			return
		}
		l.Lissajous(w, cycles)
	}
	http.HandleFunc("/", handler)
	//http.HandleFunc("/", handle3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, %s, %s \r\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\r\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\r\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\r\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\r\n", k, v)
	}
}
