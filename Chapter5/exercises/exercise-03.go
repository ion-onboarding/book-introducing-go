// Given the following array, what would x[2:5] give you?
// x := [6]string{"a","b","c","d","e","f"}

package main

import "fmt"

func main(){
	x := [6]string{"a","b","c","d","e","f"}
	fmt.Println(x[2:5]) // prints [c d e]
}