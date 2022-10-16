package main

import "fmt"

// int is a signed integer type that is at least 32 bits in size
func f(p *int) {
	c := 25
	// argument p is overwritten before first use (SA4009)
	p = &c
}

func main() {
	a := 42
	b := &a
	f(b)
	fmt.Println(*b)
}