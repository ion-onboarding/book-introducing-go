// What is the length of a slice created using make([]int, 3, 9)?

package main

import "fmt"

func main(){
	array := make([]int, 3, 9)
	fmt.Println(len(array)) // prints 3, create 3 elements max capacity 9
	fmt.Println(array) // by default values are initialized with zero - prints [0 0 0 ]
}