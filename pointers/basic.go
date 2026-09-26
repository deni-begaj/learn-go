// Package pointers is this
package pointers

import "fmt"

// PointersBasic is a general method
func PointersBasic() {
	fmt.Println("\n------### Pointers > Basic ###------")
	defer fmt.Println("------------------------------------")

	var ptr *int

	num := 4

	ptr = &num

	fmt.Println("Num: ", num)
	fmt.Println("Ptr: ", ptr)
	fmt.Println("Ptr Ref Val:", *ptr)

	*ptr = 5

	fmt.Println("Num: ", num)
	fmt.Println("Prt New Ref Val:", *ptr)
}
