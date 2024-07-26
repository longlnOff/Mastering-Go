package main

import "fmt"

func main() {
	aMap := map[string]string {
		"12": "abc",
		"qa": "333",
	}
	for key, value := range aMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}