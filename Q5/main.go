package main

import (
	"fmt"
	"os"
)

func main() {
	// must run from root dir
	content, err := os.ReadFile("./Q5/Q5.sql")
	if err != nil {
		panic(err)
	}
    fmt.Print(string(content))
}