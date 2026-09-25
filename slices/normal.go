package slices

import "fmt"

func getNames(animals []Animal) []string {
	names := make([]string, 0, len(animals))
	for _, elem := range animals {
		names = append(names, elem.Name)
	}
	return names
}

func SlicesNormal() {
	fmt.Println("\n-------### Slices > Normal ###------")
	defer fmt.Println("------------------------------------")

	dog := Animal{"Canine", "Rex", 4, true}
	cat := Animal{"Feline", "Luna", 4, true}
	fish := Animal{"Fish", "Millie", 0, true}

	group := []Animal{dog, cat, fish}
	animalNames := getNames(group)

	fmt.Println(animalNames)
}
