package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin(((f - 32) * 5 / 9) + 273.15) }

func main() {
	fmt.Println("Celcius to Fahrenheit: ", CToF(FreezingC))
	fmt.Println("Fahrenheit	to Celcius: ", FToC(Fahrenheit(BoilingC)))
	fmt.Println("Celcius to Kelvin: ", CToK(BoilingC))
	fmt.Println("Fahrenheit to Kelvin: ", FToK(Fahrenheit(BoilingC)))
}
