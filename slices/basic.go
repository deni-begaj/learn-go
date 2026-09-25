package slices

import "fmt"

func appendElem(parent []int, elem int) []int {
	return append(parent, elem)
}

func appendSlice(parent []int, otherSlice []int) []int {
	for _, elem := range otherSlice {
		parent = appendElem(parent, elem)
	}
	return parent
}

func extractEvenNumbersSlice(parent []int) []int {
	newSlice := make([]int, 0, len(parent))
	for _, elem := range parent {
		if elem%2 == 0 {
			newSlice = append(newSlice, elem)
		}
	}
	return newSlice
}

func SlicesBasic() {
	fmt.Println("\n-------### Slices > Basic ###------")
	defer fmt.Println("------------------------------------")

	list := []int{1, 2, 3, 5, 7, 11}
	fmt.Println(list)

	half := list[:4]
	fmt.Println(half)

	halfPlusOne := appendElem(half, 6)
	fmt.Println(halfPlusOne)

	mergeSlice := appendSlice(half, []int{0, 3, 9})
	fmt.Println(mergeSlice)

	evenSlice := extractEvenNumbersSlice(mergeSlice)
	fmt.Println(evenSlice)
}
