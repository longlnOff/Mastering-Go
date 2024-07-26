package main

import "fmt"

func main() {
	slice := make([]int, 3)
	arrayPtr := (*[3]int)(slice)
	fmt.Println(arrayPtr)
	fmt.Println([3]int(slice))
}