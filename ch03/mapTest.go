package main

import "fmt"


func main() {
	a := map[string]int {
		"a": 3,
		"b": 5,
	}
	_, check := a["e"]
	if !check {
		fmt.Println("Error")
	}
}
