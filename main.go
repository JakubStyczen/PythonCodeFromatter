package main

import "fmt"

func main() {
	example_string := "Kox =   2"
	res, err := FormatVariableDeclaration(example_string)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res, "Correct Format")
}
