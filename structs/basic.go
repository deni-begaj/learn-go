package structs

import "fmt"

type Person struct {
	Name string
	Age  int8
	City string
}

func StructsBasic() {
	fmt.Println("\n-------### Structs > Basic ###------")
	defer fmt.Println("------------------------------------")

	john := Person{"John Doe", 34, "Orlando"}
	fmt.Printf("%v is %v and is from %v\n", john.Name, john.Age, john.City)
}
