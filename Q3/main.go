package main

import (
	"fmt"
)

func boolSwitch(v int) bool {
	switch v != 0 {
		case true:
			return true
		case false:
			return false
	}
	return false
}

func main() {
	
	var f int = 1

	// solution (1) int > bool
	var b bool = f != 0
	
	fmt.Printf("f=%v, b=%v\n", f, b)
	if !b {
		fmt.Println("!b")
	} else {
		fmt.Println("b")
	}

	// solution (2) with switch
	if boolSwitch(1) {
		fmt.Println("boolSwitch(1)")
	}

	if !boolSwitch(0) {
		fmt.Println("!boolSwitch(0)")
	}

	fmt.Printf("boolSwitch(1)=%v\n", boolSwitch(1))
	fmt.Printf("boolSwitch(2)=%v\n", boolSwitch(2))
	fmt.Printf("boolSwitch(-1)=%v\n", boolSwitch(-1))
	fmt.Printf("boolSwitch(0)=%v\n", boolSwitch(0))

}