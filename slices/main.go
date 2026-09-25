package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 5, 7, 11}
	fmt.Println(list)

	half := list[:4]
	fmt.Println(half)
}
