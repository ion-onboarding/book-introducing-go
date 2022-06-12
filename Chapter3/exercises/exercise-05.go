// Using the example program as a starting point, write a program that converts
// from Fahrenheit into Celsius (C = (F − 32) * 5/9)

package main

import "fmt"

func main() {
	fmt.Println("provide the temperature in Fahrenheit: ")
	var F float64
	fmt.Scanf("%f", &F)

	fmt.Println("in Celsius =", (F - 32) * 5/9)
}