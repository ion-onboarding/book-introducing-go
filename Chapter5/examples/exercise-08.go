package main

import "fmt"

func main(){
	elements := make(map[string]string)
		elements["H"] = "Hydrogen"
		elements["He"] = "Helium"
		elements["Li"] = "Lithium"
		elements["Be"] = "Beryllium"
		elements["B"] = "Boron"
		elements["C"] = "Carbon"
		elements["N"] = "Nitrogen"
		elements["O"] = "Oxygen"
		elements["F"] = "Fluorine"
		elements["Ne"] = "Neon"
	
	fmt.Println(elements["Li"])
	fmt.Println(elements["NONEXISTENT"])

	name, ok := elements["NONEXISTENT"]
	fmt.Println(name, ok)

	if name, ok := elements["NONEXISTENT"]; ok {
		fmt.Println(name, ok)
	}
	if name, ok := elements["Ne"]; ok {
		fmt.Println(name, ok)
	}
}