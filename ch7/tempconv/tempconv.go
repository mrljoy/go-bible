package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit     { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius     { return Celsius((f - 32.0) * 5.0 / 9.0) }
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

type celsiusFlag struct{ Celsius }
type fahrenheitFlag struct{ Fahrenheit }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (f *fahrenheitFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Fahrenheit = CToF(Celsius(value))
		return nil
	case "F", "°F":
		f.Fahrenheit = Fahrenheit(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func FahrenheitFlag(name string, value Fahrenheit, usage string) *Fahrenheit {
	f := fahrenheitFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Fahrenheit
}
