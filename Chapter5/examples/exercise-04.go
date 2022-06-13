package main

import "fmt"

func main(){
	 x := [5]float64{
		98,
		93,
		77,
		82,
		// 83,
	}

	for i, _ := range x {
		fmt.Println(x[i])
	}
}