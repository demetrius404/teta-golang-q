package main

import "fmt"

type MyError struct{}

func (MyError) Error() string {
	return "MyError!"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	var err *MyError
	errorHandler(err)
}