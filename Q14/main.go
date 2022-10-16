package main

import (
	"fmt"
	"os"
)

func main() {
	// must run from root dir
	content, err := os.ReadFile("./Q14/Q14.sql")
	if err != nil {
		panic(err)
	}
    fmt.Print(string(content))
}