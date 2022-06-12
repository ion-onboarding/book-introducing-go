// Write another program that converts from feet into meters (1 ft = 0.3048 m).

package main

import "fmt"

func main() {
	fmt.Println("distance in feet: ")
	var input float64

	fmt.Scanf("%f", &input)

	fmt.Println("feet transformed into meters: ", input / 3.2808)
}