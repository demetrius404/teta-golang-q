package main

import (
	"fmt"
	"reflect"
)

// type alias
type str = string

func typeSwitch(v interface{}) {
	switch v.(type) {
		case int: 
			fmt.Println("type: int")
		case string:
			fmt.Println("type: str / string")
		default:
			fmt.Println("type: unknown")
	}
}

func main() {
	var i int = 1
	var s str = "text"
	
	fmt.Printf("i=%v\n", i)
	fmt.Printf("s=%v\n", s)
	fmt.Println(reflect.TypeOf(i)) // int
	fmt.Println(reflect.TypeOf(s)) // string

	typeSwitch(i) // type: int
	typeSwitch(s) // type: str / string
}
