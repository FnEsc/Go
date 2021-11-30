package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreeingC      Celsius = 0
	BoilingC      Celsius = 100
)

func C2F(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func F2C(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c)
	fmt.Printf("%g\n", BoilingC-FreeingC) // 100
	fmt.Println(c == 0)                   // true
	fmt.Println(f >= 0)                   // true
	fmt.Println(c == Celsius(f))          // true
	// fmt.Println(c == f)                   // compile error: type mismatch
}

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) } // 作打印处理
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
